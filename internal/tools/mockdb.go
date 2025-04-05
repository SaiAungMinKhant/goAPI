package tools

import "time"

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"alex": {
		AuthToken: "12345",
		Username:  "alex",
	},
	"bob": {
		AuthToken: "b0bAuthT0ken",
		Username:  "bob",
	},
	"charlie": {
		AuthToken: "CharlieAuth",
		Username:  "charlie",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"alex": {
		Coins:    100,
		Username: "alex",
	},
	"bob": {
		Coins:    50,
		Username: "bob",
	},
	"charlie": {
		Coins:    200,
		Username: "charlie",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	// Simulate DB call
	time.Sleep(time.Second * 1 )

	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
			return nil
	}
	
	return &clientData
}

func (d *mockDB) GetUserCoins(username string) *CoinDetails {
	// Simulate DB call
	time.Sleep(time.Second * 1 )

	var clientData = CoinDetails{}
	clientData, ok := mockCoinDetails[username]
	if !ok {
			return nil
	}
	return &clientData
}

func (d *mockDB) SetupDatabase() error {
	return nil
}