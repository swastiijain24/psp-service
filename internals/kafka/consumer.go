package kafka

import (
	"github.com/segmentio/kafka-go"
)

type Consumer struct {
	Reader *kafka.Reader
}

func NewConsumer(brokers []string , topic string , groupId string) *Consumer {

	Reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: brokers,
		Topic:   topic,
		GroupID: groupId,
		CommitInterval : 0, //disabling auto commit 
	})

	return &Consumer{
		Reader: Reader,
	}
}
