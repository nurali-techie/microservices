package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("search-service OK"))
}

func SearchHandler(esClient *ElasticClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
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
