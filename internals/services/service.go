package services

import (
	"context"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

type Service struct {
	producer *kafka.Writer
	consumer *kafka.Reader
}

func NewService(producer *kafka.Writer, consumer *kafka.Reader) *Service {
	return &Service{
		producer: producer,
		consumer: consumer,
	}
}

func (s *Service) ProduceMsg() {
	msg := kafka.Message{
		Key:   []byte("key1"),
		Value: []byte("txn request sent from psp!"),
	}

	err := s.producer.WriteMessages(context.Background(), msg)
	if err != nil {
		log.Fatal("failed to write message:", err)
	}

	// log.Println("Message sent successfully")

}

func (s *Service) ConsumeMsg() {
	ctx := context.Background()
	for {
		msg, err := s.consumer.FetchMessage(ctx)
		if err != nil {
			log.Fatal("error reading message:", err)
		}

		fmt.Printf("Message: %s\n", string(msg.Value))
		err = s.consumer.CommitMessages(ctx, msg)
		if err != nil {
			log.Fatal("commit failed:", err)
		}
	}
}
