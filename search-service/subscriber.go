package main

import (
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
}

func NewMenuItemsSubscriber(consumer *kafka.Consumer) *MenuItemSubscribe {
	err := consumer.Subscribe(KAFKA_TOPIC_MENUITEMS, nil)
	if err != nil {
		panic(err)
	}

	return &MenuItemSubscribe{
		consumer: consumer,
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
			log.Infof("Msg received, topic:%s, offset:%s, key:%s, value:%s", *msg.TopicPartition.Topic, msg.TopicPartition.Offset, string(msg.Key), string(msg.Value))
		}
	}()
}
