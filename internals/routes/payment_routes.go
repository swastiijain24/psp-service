package routes

import (
	"github.com/gin-gonic/gin"
	handler "github.com/swastiijain24/psp/internals/handlers"
	apiAuth "github.com/swastiijain24/psp/internals/middlewares/api_key_auth"
)

func RegisterNpciRoutes(r *gin.Engine, apiKeyAuth *apiAuth.APIMiddleware , paymentHandler *handler.PaymentHandler) {

	npciRoutes := r.Group("/npci")
	npciRoutes.Use(apiKeyAuth.ApiAuthentication())

	{
		npciRoutes.POST("/payment", paymentHandler.ProcessPayment)
		npciRoutes.GET("status/:transactionId", paymentHandler.GetTxnStatus)
	}
}