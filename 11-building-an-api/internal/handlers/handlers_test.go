package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/wrath-02/go_basics/11-building-an-api/api"
	"github.com/wrath-02/go_basics/11-building-an-api/internal/tools"
)

// Run with: go test -v ./11-building-an-api/...

func TestGetCoinBalance(t *testing.T) {
	db, _ := tools.NewDatabase()
	router := NewRouter(db)

	tests := []struct {
		name       string
		url        string
		token      string
		wantStatus int
		wantCoins  int64
	}{
		{"valid user and token", "/account/coins?username=alex", "123ABC", http.StatusOK, 100},
		{"wrong token", "/account/coins?username=alex", "WRONG", http.StatusUnauthorized, 0},
		{"missing token", "/account/coins?username=alex", "", http.StatusUnauthorized, 0},
		{"unknown user", "/account/coins?username=nobody", "123ABC", http.StatusUnauthorized, 0},
		{"missing username", "/account/coins", "123ABC", http.StatusUnauthorized, 0},
		{"wrong method", "/account/coins?username=alex", "123ABC", http.StatusMethodNotAllowed, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			method := http.MethodGet
			if tt.name == "wrong method" {
				method = http.MethodPost
			}
			req := httptest.NewRequest(method, tt.url, nil)
			if tt.token != "" {
				req.Header.Set("Authorization", tt.token)
			}
			rec := httptest.NewRecorder()

			router.ServeHTTP(rec, req)

			if rec.Code != tt.wantStatus {
				t.Fatalf("status = %d, want %d (body: %s)", rec.Code, tt.wantStatus, rec.Body)
			}
			if tt.wantStatus == http.StatusOK {
				var resp api.CoinBalanceResponse
				if err := json.NewDecoder(rec.Body).Decode(&resp); err != nil {
					t.Fatalf("decode: %v", err)
				}
				if resp.Balance != tt.wantCoins {
					t.Errorf("balance = %d, want %d", resp.Balance, tt.wantCoins)
				}
			}
		})
	}
}

func TestHealth(t *testing.T) {
	db, _ := tools.NewDatabase()
	rec := httptest.NewRecorder()
	NewRouter(db).ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/health", nil))
	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200", rec.Code)
	}
}

// Handlers can also be tested directly, without the router or middleware.
func TestGetCoinBalanceMissingUsername(t *testing.T) {
	db, _ := tools.NewDatabase()
	rec := httptest.NewRecorder()
	GetCoinBalance(db).ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/account/coins", nil))
	if rec.Code != http.StatusBadRequest {
		t.Fatalf("status = %d, want 400", rec.Code)
	}
}
