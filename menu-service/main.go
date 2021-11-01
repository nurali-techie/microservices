package main

import (
	"net/http"

	log "github.com/sirupsen/logrus"

	commons "github.com/nurali-techie/microservices/commons-go"
)

func main() {
	log.Info("starting menu-service")

	db := commons.SetupDatabase(getDBConfig(), getDatabaseSQL())
	defer db.Close()

	http.HandleFunc("/ping", pingHandler)

	log.Info("menu-service started")
	log.Fatal(http.ListenAndServe("0.0.0.0:2011", nil))
}
