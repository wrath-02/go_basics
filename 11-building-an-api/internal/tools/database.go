// Package tools contains the data layer. Handlers depend on the
// DatabaseInterface, not on a concrete database, so we can swap the
// mock for a real database (or a test fake) without touching handlers.
package tools

type LoginDetails struct {
	AuthToken string
	Username  string
}

type CoinDetails struct {
	Coins    int64
	Username string
}

type DatabaseInterface interface {
	GetUserLoginDetails(username string) (*LoginDetails, bool)
	GetUserCoins(username string) (*CoinDetails, bool)
}

// NewDatabase returns the database used by the server.
func NewDatabase() (DatabaseInterface, error) {
	return newMockDB(), nil
}
