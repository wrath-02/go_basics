package handlers

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/wrath-02/go_basics/11-building-an-api/api"
	"github.com/wrath-02/go_basics/11-building-an-api/internal/tools"
)

var (
	ErrMissingUsername = errors.New("query parameter 'username' is required")
	ErrUserNotFound    = errors.New("user not found")
)

// GetCoinBalance handles GET /account/coins?username=<name>.
// It returns a handler that closes over db (dependency injection).
func GetCoinBalance(db tools.DatabaseInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := api.CoinBalanceParams{Username: r.URL.Query().Get("username")}
		if params.Username == "" {
			api.RequestError(w, ErrMissingUsername)
			return
		}

		details, ok := db.GetUserCoins(params.Username)
		if !ok {
			slog.Warn("coins not found", "username", params.Username)
			api.NotFoundError(w, ErrUserNotFound)
			return
		}

		api.WriteJSON(w, http.StatusOK, api.CoinBalanceResponse{
			Code:    http.StatusOK,
			Balance: details.Coins,
			User:    details.Username,
		})
	}
}
