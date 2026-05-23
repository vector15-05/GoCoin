package tools

import (
	log "github.com/sirupsen/logrus"
)

type LoginDetails struct {
	Username string
	AuthToken string
}

type CoinDetails struct {
	Coins int64
	Username string
}

type DataBaseInterface interface {
	GetUserLoginDetails(username string) *LoginDetails
	GetUserCoinDetails(username string) *CoinDetails
	SetupDatabase() error
}

func NewDatabase() (DataBaseInterface, error) {
	var db mockDB
	if err := db.SetupDatabase(); err != nil {
		log.Error(err)
		return nil, err
	}
	return &db, nil
}

