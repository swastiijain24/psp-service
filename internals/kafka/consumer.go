package kafka

import (
	"github.com/segmentio/kafka-go"
)

type Consumer struct {
	Reader *kafka.Reader
}

func NewConsumer(brokers []string , topic string ) *Consumer {

	Reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: brokers,
		Topic:   topic,
		GroupID: "group-1",
	})

	return &Consumer{
		Reader: Reader,
	}
}
