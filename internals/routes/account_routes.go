package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/swastiijain24/psp/internals/handlers"
	apiAuth "github.com/swastiijain24/psp/internals/middlewares/api_key_auth"
)

func RegisterAccountRoutes(r *gin.Engine, apiKeyAuth *apiAuth.APIMiddleware , accountHandler *handlers.AccountHandler) {

	accountRoutes := r.Group("/account")
	accountRoutes.Use(apiKeyAuth.ApiAuthentication())
	{
		accountRoutes.GET("/discover", accountHandler.DiscoverAccounts)
		accountRoutes.POST("/mpin", accountHandler.SetMpin)
		accountRoutes.PUT("/mpin", accountHandler.ChangeMpin)
		accountRoutes.GET("/balance", accountHandler.GetBalance)
	}
}