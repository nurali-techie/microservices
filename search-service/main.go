package main

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("starting search-service")
	http.HandleFunc("/ping", pingHandler)

	log.Info("search-service started")
	log.Fatal(http.ListenAndServe("0.0.0.0:2021", nil))
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("search-service OK"))
}
