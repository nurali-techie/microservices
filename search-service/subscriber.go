package main

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	log "github.com/sirupsen/logrus"
)

var (
	KAFKA_BOOTSTRAP_SERVERS = "kafka:9092"
	KAFKA_TOPIC_MENUITEMS   = "menuitems"
)

func NewConsumer() *kafka.Consumer {
	config := &kafka.ConfigMap{
		"bootstrap.servers":  KAFKA_BOOTSTRAP_SERVERS,
		"group.id":           "demo-group",
		"enable.auto.commit": "false",
		"auto.offset.reset":  "earliest",
	}

	consumer, err := kafka.NewConsumer(config)
	if err != nil {
		panic(err)
	}
	return consumer
}

type MenuItemSubscribe struct {
	consumer *kafka.Consumer
	esClient *ElasticClient
}

func NewMenuItemsSubscriber(consumer *kafka.Consumer, esClient *ElasticClient) *MenuItemSubscribe {
	err := consumer.Subscribe(KAFKA_TOPIC_MENUITEMS, nil)
	if err != nil {
		panic(err)
	}

	return &MenuItemSubscribe{
		consumer: consumer,
		esClient: esClient,
	}
}

func (s *MenuItemSubscribe) Start() {
	go func() {
		for {
			msg, err := s.consumer.ReadMessage(-1)
			if err != nil {
				log.Errorf("Error: reading menuitems topics, %v", err)
				continue
			}
			log.Infof("Msg received, topic:%s, offset:%s, key:%s", *msg.TopicPartition.Topic, msg.TopicPartition.Offset, string(msg.Key))

			err = s.insertIntoES(msg.Value)
			if err != nil {
				log.Errorf("Error: insert to elasticserach failed, %v", err)
			}
		}
	}()
}

func (s *MenuItemSubscribe) insertIntoES(content []byte) error {
	var menuItem MenuItem
	if err := json.NewDecoder(bytes.NewBuffer(content)).Decode(&menuItem); err != nil {
		return err
	}

	log.Infof("menuitem: %v", menuItem)
	err := s.esClient.InsertMenuItem(&menuItem)
	if err != nil {
		return fmt.Errorf("insert to elasticserach failed for menu_item %q, %v", menuItem.ID, err)
	}
	return nil
}
