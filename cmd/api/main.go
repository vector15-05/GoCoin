package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/avukadin/goapi/internal/handlers"
	log "github.com/sirupsen/logrus"
)

func main(){

	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handlers(r)

	fmt.Println("Starting Go API service....")

	goCoin := `
   ______   ______       ______   ______   __  __   __    
  /\  ___\ /\  __ \     /\  ___\ /\  __ \ /\ \/\ "-.\ \   
  \ \ \____\ \ \/\ \    \ \ \____\ \ \/\ \\ \ \ \ \-.  \  
   \ \_____\\ \_____\    \ \_____\\ \_____\\ \_\ \_\\"\_\ 
    \/_____/ \/_____/     \/_____/ \/_____/ \/_/\/_/ \/_/ 
`

	err := http.ListenAndServe("localhost:8000",r)
	if err != nil {
		log.Error(err)
	}

}