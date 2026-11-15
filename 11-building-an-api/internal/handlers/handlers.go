// Package handlers wires URL routes to handler functions.
package handlers

import (
	"net/http"

	"github.com/wrath-02/go_basics/11-building-an-api/api"
	"github.com/wrath-02/go_basics/11-building-an-api/internal/middleware"
	"github.com/wrath-02/go_basics/11-building-an-api/internal/tools"
)

// NewRouter builds the HTTP router. Since Go 1.22 the standard ServeMux
// supports methods and wildcards in patterns ("GET /items/{id}"), so no
// third-party router is needed for most APIs.
func NewRouter(db tools.DatabaseInterface) http.Handler {
	mux := http.NewServeMux()

	// Public route
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		api.WriteJSON(w, http.StatusOK, map[string]string{"status": "ok"})
	})

	// Protected route: Authorization runs first, then GetCoinBalance.
	auth := middleware.Authorization(db)
	mux.Handle("GET /account/coins", auth(GetCoinBalance(db)))

	// Logger wraps everything.
	return middleware.Logger(mux)
}
