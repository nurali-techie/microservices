package main

import (
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	commons "github.com/nurali-techie/microservices/commons-go"
)

func main() {
	log.Info("starting menu-service")

	// database
	db := commons.SetupDatabase(getDBConfig(), getDatabaseSQL())
	defer db.Close()

	// repo
	restoRepo := NewRestoRepo(db)

	// handler
	r := mux.NewRouter()
	r.HandleFunc("/ping", pingHandler)

	restoHandler := NewRestoHandler(restoRepo)
	r.HandleFunc("/v1/restaurants", restoHandler.NewRestoHandler).Methods(http.MethodPost)
	r.HandleFunc("/v1/restaurants/{id}", restoHandler.GetRestoHandler).Methods(http.MethodGet)

	log.Info("menu-service started")
	log.Fatal(http.ListenAndServe("0.0.0.0:2011", r))
}
