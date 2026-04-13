package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/swastiijain24/psp/internals/handlers"
	"github.com/swastiijain24/psp/internals/kafka"
	"github.com/swastiijain24/psp/internals/repository"
	"github.com/swastiijain24/psp/internals/services"
	"github.com/swastiijain24/psp/internals/workers"
	"github.com/swastiijain24/psp/internals/routes"
)

func main() {
	redisStore := repository.NewRedisStore("localhost:6379")
	kafkaProducer := kafka.NewProducer("localhost:9092", "payment.request.v1")
	vpaSvc := services.NewVpaService()

	paymentSvc := services.NewPaymentService(vpaSvc, kafkaProducer, redisStore)
	paymentHandler := handlers.NewPaymentHandler(paymentSvc)

	consumer := kafka.NewConsumer([]string {"localhost:9092"}, "payment.response.v1")
	worker := workers.NewResponseWorker(consumer, paymentSvc)
	
	go worker.StartConsumingResponse(context.Background())
	r := gin.Default()
	routes.RegisterRoutes(r, paymentHandler)

	log.Println("PSP API starting on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
	

	select {}
}