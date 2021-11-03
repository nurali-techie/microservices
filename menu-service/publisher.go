package main

import (
	"bytes"
	"encoding/json"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var (
	KAFKA_BOOTSTRAP_SERVERS = "kafka:9092"
	KAFKA_TOPIC_MENUITEMS   = "menuitems"
)

func NewKafkaProducer() *kafka.Producer {
	config := &kafka.ConfigMap{
		"bootstrap.servers": KAFKA_BOOTSTRAP_SERVERS,
	}
	producer, err := kafka.NewProducer(config)
	if err != nil {
		panic(err)
	}
	return producer
}

type Publisher struct {
	producer *kafka.Producer
}

func NewPublisher(producer *kafka.Producer) *Publisher {
	return &Publisher{
		producer: producer,
	}
}

func (p *Publisher) PublishMenuItem(menuItem *MenuItem) error {
	var value bytes.Buffer
	if err := json.NewEncoder(&value).Encode(menuItem); err != nil {
		return err
	}
	key := []byte(menuItem.ID)

	msg := &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &KAFKA_TOPIC_MENUITEMS, Partition: kafka.PartitionAny},
		Key:            key,
		Value:          value.Bytes(),
	}

	return p.producer.Produce(msg, nil)
}
