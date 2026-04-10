package kafka

import (
	"github.com/segmentio/kafka-go"
)

type Consumer struct {
	reader *kafka.Reader
}

func NewConsumer(brokers []string , topic string ) *Consumer {

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: brokers,
		Topic:   topic,
		GroupID: "group-1",
	})

	return &Consumer{
		reader: reader,
	}
}

func (c* Consumer) ConsumeEvent(){

}
