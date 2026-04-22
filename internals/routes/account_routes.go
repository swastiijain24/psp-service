package routes

import (
	"github.com/gin-gonic/gin"
	handler "github.com/swastiijain24/psp/internals/handlers"
)

func RegisterAccountRoutes(r *gin.Engine, accountHandler *handler.AccountHandler) {

	accountRoutes := r.Group("/account")
	{
		accountRoutes.POST("/discover")
	}
}