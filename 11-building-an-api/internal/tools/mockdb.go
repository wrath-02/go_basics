package tools

import "time"

// mockDB is an in-memory stand-in for a real database.
type mockDB struct {
	logins map[string]LoginDetails
	coins  map[string]CoinDetails
}

func newMockDB() *mockDB {
	return &mockDB{
		logins: map[string]LoginDetails{
			"alex":  {AuthToken: "123ABC", Username: "alex"},
			"jason": {AuthToken: "456DEF", Username: "jason"},
			"marie": {AuthToken: "789GHI", Username: "marie"},
		},
		coins: map[string]CoinDetails{
			"alex":  {Coins: 100, Username: "alex"},
			"jason": {Coins: 200, Username: "jason"},
			"marie": {Coins: 300, Username: "marie"},
		},
	}
}

func (d *mockDB) GetUserLoginDetails(username string) (*LoginDetails, bool) {
	time.Sleep(10 * time.Millisecond) // simulate a DB round trip
	details, ok := d.logins[username]
	return &details, ok
}

func (d *mockDB) GetUserCoins(username string) (*CoinDetails, bool) {
	time.Sleep(10 * time.Millisecond)
	details, ok := d.coins[username]
	return &details, ok
}
