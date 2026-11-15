// Topic 11: Building an API with only the standard library.
//
// Run:  go run ./11-building-an-api/cmd/api
// Call: curl -H "Authorization: 123ABC" "localhost:8000/account/coins?username=alex"
package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/wrath-02/go_basics/11-building-an-api/internal/handlers"
	"github.com/wrath-02/go_basics/11-building-an-api/internal/tools"
)

const addr = ":8000"

func main() {
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, nil)))

	db, err := tools.NewDatabase()
	if err != nil {
		slog.Error("database setup failed", "err", err)
		os.Exit(1)
	}

	srv := &http.Server{
		Addr:         addr,
		Handler:      handlers.NewRouter(db),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// Start the server in a goroutine so main can wait for Ctrl+C.
	go func() {
		slog.Info("starting GO API service", "addr", addr)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error("server error", "err", err)
			os.Exit(1)
		}
	}()

	// Graceful shutdown: wait for SIGINT/SIGTERM, then let in-flight requests finish.
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	<-ctx.Done()

	slog.Info("shutting down")
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(shutdownCtx); err != nil {
		slog.Error("shutdown error", "err", err)
	}
}
