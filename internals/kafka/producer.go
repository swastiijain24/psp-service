package kafka

import (
	"context"
	"time"

	"github.com/segmentio/kafka-go"
)

type Producer struct{
	writer *kafka.Writer
}

func NewProducer( address string, topic string) *Producer {
	writer :=  &kafka.Writer{
		Addr:     kafka.TCP(address),
		Topic:    topic,
		Balancer: &kafka.Hash{},

		BatchSize:    100,
		BatchTimeout: 10 * time.Millisecond,
		RequiredAcks: kafka.RequireAll,
	}

	return &Producer{
		writer: writer,
	}
}

func (p *Producer) ProduceEvent(ctx context.Context, key string, value []byte) error {
	return p.writer.WriteMessages(context.Background(), kafka.Message{
		Key: []byte(key),
		Value: value,
	})
}
