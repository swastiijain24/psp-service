package routes

import (
	"github.com/gin-gonic/gin"
	handler "github.com/swastiijain24/psp/internals/handlers"
)

func RegisterRoutes(r *gin.Engine, paymentHandler *handler.PaymentHandler) {

	npciRoutes := r.Group("/npci")
	{
		npciRoutes.POST("/payment", paymentHandler.ProcessPayment)
	}
}