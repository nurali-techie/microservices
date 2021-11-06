package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("search-service OK"))
}

func SearchHandler(esClient *ElasticClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")

		name := r.URL.Query().Get("name")
		log.Infof("search called with name %q", name)

		menuItems, err := esClient.SearchMenuItem(name)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error: searching in elasticsearch failed, %v", err), http.StatusInternalServerError)
			return
		}

		if err := json.NewEncoder(w).Encode(menuItems); err != nil {
			http.Error(w, fmt.Sprintf("Error: menu items to search result json failed, %v", err), http.StatusInternalServerError)
			return
		}
	}
}
