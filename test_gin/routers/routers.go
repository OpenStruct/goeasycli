package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	//config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	router.Use(cors.New(config))
	HealthRoute(router)
	v1 := router.Group("v1")
	{

		userGroup := v1.Group("user")
		{
			UserRoutes(userGroup)
		}
	}
	return router
}
