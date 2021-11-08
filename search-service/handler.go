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

func SearchHandler(esClient *ElasticClient, cache *Cache, menuService *MenuService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		log.Infof("search called with name %q", name)

		menuItems, err := esClient.SearchMenuItem(name)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error: searching in elasticsearch failed, %v", err), http.StatusInternalServerError)
			return
		}

		searchResponse := make([]*SearchResult, 0)
		for _, menuItem := range menuItems {
			resto, err := cache.GetRestaurant(menuItem.RestoID)
			if err != nil {
				log.Errorf("cache get failed for restaurant, id %q, %v", menuItem.RestoID, err)
			}
			if resto == nil {
				resto, err = menuService.GetRestaurant(menuItem.RestoID)
				if err != nil {
					log.Errorf("grpc call failed for restaurant, id %q, %v", menuItem.RestoID, err)
					continue
				}
				err = cache.SetRestaurant(resto)
				if err != nil {
					log.Errorf("cache set failed for restaurnat, id %q, %v", menuItem.RestoID, err)
				}
			}
			searchResult := &SearchResult{
				MenuItem:  *menuItem,
				RestoName: resto.Name,
			}
			searchResponse = append(searchResponse, searchResult)
		}

		if err := json.NewEncoder(w).Encode(searchResponse); err != nil {
			http.Error(w, fmt.Sprintf("Error: search result to json failed, %v", err), http.StatusInternalServerError)
			return
		}
	}
}
