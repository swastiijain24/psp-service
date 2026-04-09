package main

import (
	

	"github.com/swastiijain24/npci-shared/constants"
	"github.com/swastiijain24/psp/internals/kafka"
	"github.com/swastiijain24/psp/internals/services"
)

func main() {
	address := "localhost:9092"
	brokers := []string{address}

	producer := kafka.NewProducer(address, constants.TopicPaymentRequest)
	defer producer.Close()

	consumer := kafka.NewConsumer(brokers, constants.TopicPaymentResponse)
	defer consumer.Close()

	service := services.NewService(producer, consumer)

	go service.ConsumeMsg()
	go service.ProduceMsg()

	
	select {}

}
