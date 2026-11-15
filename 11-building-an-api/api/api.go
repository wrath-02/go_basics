// Package api holds the request/response types shared by the server
// and its clients, plus helpers for writing JSON errors.
package api

import (
	"encoding/json"
	"net/http"
)

// CoinBalanceParams are the query parameters for GET /account/coins.
type CoinBalanceParams struct {
	Username string
}

// CoinBalanceResponse is returned on success.
type CoinBalanceResponse struct {
	Code    int    `json:"code"`    // HTTP status code
	Balance int64  `json:"balance"` // account balance
	User    string `json:"username"`
}

// Error is returned whenever a request fails.
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// WriteJSON encodes v as JSON with the given status code.
func WriteJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

func writeError(w http.ResponseWriter, message string, code int) {
	WriteJSON(w, code, Error{Code: code, Message: message})
}

// RequestError: the client sent something invalid (400).
func RequestError(w http.ResponseWriter, err error) {
	writeError(w, err.Error(), http.StatusBadRequest)
}

// UnauthorizedError: missing or wrong credentials (401).
func UnauthorizedError(w http.ResponseWriter, err error) {
	writeError(w, err.Error(), http.StatusUnauthorized)
}

// NotFoundError: the resource doesn't exist (404).
func NotFoundError(w http.ResponseWriter, err error) {
	writeError(w, err.Error(), http.StatusNotFound)
}

// InternalError: something went wrong on our side (500).
// Never leak internal details to the client.
func InternalError(w http.ResponseWriter) {
	writeError(w, "an unexpected error occurred", http.StatusInternalServerError)
}
