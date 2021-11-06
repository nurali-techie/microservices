package main

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("starting customer-webui")

	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/search", searchHandler)

	log.Info("customer-webui started")
	log.Fatal(http.ListenAndServe("0.0.0.0:2000", nil))
}
