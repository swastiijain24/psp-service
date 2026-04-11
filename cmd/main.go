package main

import (
	"sync"

	"github.com/gin-gonic/gin"
	handler "github.com/swastiijain24/psp/internals/handlers"
	"github.com/swastiijain24/psp/internals/kafka"
	"github.com/swastiijain24/psp/internals/routes"
	"github.com/swastiijain24/psp/internals/services"
)

var ResponseMap sync.Map

func main() {
	
	r:= gin.New()

	vpaService := services.NewVpaService()

	paymentReqProducer := kafka.NewProducer("localhost:9092", "payment.request.v1")
	paymentService := services.NewPaymentService(vpaService, paymentReqProducer)
	paymentHandler := handler.NewHandler(paymentService)

	routes.RegisterRoutes(r, paymentHandler)
}
