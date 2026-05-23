package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/vector15-05/GoCoin/internal/handlers"
	log "github.com/sirupsen/logrus"
)

func main(){

	log.SetReportCaller(true)
	r := chi.NewRouter()
	handlers.Handlers(r)

	fmt.Println("Starting Go API service....")

	goCoin := `
   ______   ______       ______   ______   __  __   __    
  /\  ___\ /\  __ \     /\  ___\ /\  __ \ /\ \/\ "-.\ \   
  \ \ \____\ \ \/\ \    \ \ \____\ \ \/\ \\ \ \ \ \-.  \  
   \ \_____\\ \_____\    \ \_____\\ \_____\\ \_\ \_\\"\_\ 
    \/_____/ \/_____/     \/_____/ \/_____/ \/_/\/_/ \/_/ 
`

	fmt.Println(goCoin)

	err := http.ListenAndServe("localhost:8000",r)
	if err != nil {
		log.Error(err)
	}

}