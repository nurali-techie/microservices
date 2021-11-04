package main

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("starting search-service")

	// elasticsearch
	log.Info("creating elasticsearch indexes")
	esClient := NewElasticClient()
	esClient.CreateIndexes()
	log.Info("elasticsearch indexes created")

	// kafka
	log.Info("creating kafka consumers")
	consumer := NewConsumer()
	defer consumer.Close()
	menuItemSub := NewMenuItemsSubscriber(consumer, esClient)
	menuItemSub.Start()
	log.Info("kafka consumers created")

	// handler
	log.Info("registering handlers")
	http.HandleFunc("/ping", pingHandler)
	log.Info("handlers registered")

	log.Info("search-service started")
	log.Fatal(http.ListenAndServe("0.0.0.0:2021", nil))
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("search-service OK"))
}
