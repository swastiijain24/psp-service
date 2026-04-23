package routes

import (
	"github.com/gin-gonic/gin"
	handler "github.com/swastiijain24/psp/internals/handlers"
	apiAuth "github.com/swastiijain24/psp/internals/middlewares/api_key_auth"
)

func RegisterVpaRoutes(r *gin.Engine,  apiKeyAuth *apiAuth.APIMiddleware ,vpaHandler *handler.VpaHandler) {

	vpaRoutes := r.Group("/vpa")
	vpaRoutes.Use(apiKeyAuth.ApiAuthentication())

	{
		vpaRoutes.POST("/register")
	}
}