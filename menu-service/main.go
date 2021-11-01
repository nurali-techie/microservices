package main

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("starting menu-service")
	http.HandleFunc("/ping", pingHandler)

	log.Info("menu-service started")
	log.Fatal(http.ListenAndServe("0.0.0.0:2011", nil))
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("menu-service OK"))
}
