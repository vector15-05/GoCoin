package tools

import (
	"time"
)

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"alex":{
		AuthToken: "1234567890",
		Username: "alex",
	},
	"bob":{
		AuthToken: "0987654321",
		Username: "bob",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"alex":{
		Coins: 100,
		Username: "alex",
	},
	"bob":{
		Coins: 50,
		Username: "bob",
	},
}

func (db *mockDB) GetUserLoginDetails(username string) *LoginDetails{	
	time.Sleep(time.Second * 1)

	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]

	if !ok {
		return nil
	}

	return &clientData
}

func (db *mockDB) GetUserCoinDetails(username string) *CoinDetails{	
	time.Sleep(time.Second * 1)

	var clientData = CoinDetails{}
	clientData, ok := mockCoinDetails[username]

	if !ok {
		return nil
	}

	return &clientData
}

func (db *mockDB) SetupDatabase() error{
	return nil
}