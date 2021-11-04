package main

import (
	"context"
	"fmt"

	"github.com/olivere/elastic/v7"
)

var (
	ES_HOST_URL          = "http://search_elasticsearch:9200"
	MENU_ITEM_INDEX_NAME = "menu_items"
)

const MENU_ITEMS_MAPPING = `
{
  "settings": {
    "number_of_shards": 1,
    "number_of_replicas": 0
  },
  "mappings": {
    "properties": {
      "name": {
        "type": "text"
      },
      "category": {
        "type": "keyword"
      },
      "cuisine": {
        "type": "keyword"
      },
      "price": {
        "type": "float"
      },
      "restaurant_id": {
        "type": "keyword"
      }
    }
  }
}`

type ElasticClient struct {
	client *elastic.Client
}

func NewElasticClient() *ElasticClient {
	client, err := elastic.NewClient(
		elastic.SetURL(ES_HOST_URL),
	)
	if err != nil {
		panic(err)
	}
	return &ElasticClient{
		client: client,
	}
}

func (e *ElasticClient) CreateIndexes() {
	ctx := context.Background()

	if exists, err := e.client.IndexExists(MENU_ITEM_INDEX_NAME).Do(ctx); err != nil {
		panic(err)
	} else if exists {
		return
	}

	res, err := e.client.CreateIndex(MENU_ITEM_INDEX_NAME).BodyString(MENU_ITEMS_MAPPING).Do(ctx)
	if err != nil {
		panic(err)
	}

	if !res.Acknowledged {
		panic(fmt.Sprintf("%q index not acknowledged", MENU_ITEM_INDEX_NAME))
	}
}

func (e *ElasticClient) InsertMenuItem(menuItem *MenuItem) error {
	indexService := e.client.Index().Index(MENU_ITEM_INDEX_NAME)
	_, err := indexService.Id(menuItem.ID).BodyJson(menuItem).Do(context.Background())
	return err
}
