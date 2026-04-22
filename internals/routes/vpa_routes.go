package routes

import (
	"github.com/gin-gonic/gin"
	handler "github.com/swastiijain24/psp/internals/handlers"
)

func RegisterVpaRoutes(r *gin.Engine, vpaHandler *handler.VpaHandler) {

	vpaRoutes := r.Group("/vpa")
	{
		vpaRoutes.POST("/register")
	}
}