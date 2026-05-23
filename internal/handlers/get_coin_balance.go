package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/vector15-05/GoCoin/internal/tools"
	"github.com/vector15-05/GoCoin/api"
	log "github.com/sirupsen/logrus"
	"github.com/gorilla/schema"
)

func GetCoinBalance(w http.ResponseWriter, r *http.Request){

	var params = api.CoinBalanceParams{}
	decoder := schema.NewDecoder()

	if err := decoder.Decode(&params, r.URL.Query()); err != nil {
		log.Error(err)
		api.InternalErrorHandler(w, err)
		return
	}

	database, err := tools.NewDatabase()
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w, err)
		return
	}

	tokenDetails := database.GetUserCoinDetails(params.Username)

	if tokenDetails == nil {
		log.Error("No Coin Details Found for user: " + params.Username)
		api.RequestErrorHandler(w, errors.New("no coin details found"))
		return
	}

	response := api.CoinBalanceResponse{
		Code:    http.StatusOK,
		Balance: tokenDetails.Coins,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Error(err)
		api.InternalErrorHandler(w, err)
		return
	}
}