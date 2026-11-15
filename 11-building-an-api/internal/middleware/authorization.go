// Package middleware contains functions that wrap http.Handlers to run
// code before/after the real handler (auth, logging, ...).
package middleware

import (
	"errors"
	"net/http"

	"github.com/wrath-02/go_basics/11-building-an-api/api"
	"github.com/wrath-02/go_basics/11-building-an-api/internal/tools"
)

var ErrUnauthorized = errors.New("invalid username or token")

// Authorization checks that the "Authorization" header matches the token
// stored for the "username" query parameter.
//
// A middleware is just: func(next http.Handler) http.Handler
func Authorization(db tools.DatabaseInterface) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			username := r.URL.Query().Get("username")
			token := r.Header.Get("Authorization")

			if username == "" || token == "" {
				api.UnauthorizedError(w, ErrUnauthorized)
				return
			}

			login, ok := db.GetUserLoginDetails(username)
			if !ok || login.AuthToken != token {
				api.UnauthorizedError(w, ErrUnauthorized)
				return
			}

			next.ServeHTTP(w, r) // authorized: continue to the real handler
		})
	}
}
