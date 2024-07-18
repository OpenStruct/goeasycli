package routers

import (
	_ "hjo/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Project API Documentation
// @version 1.0

// @host localhost:8080
// @BasePath /v1
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
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL("http://localhost:8080/swagger/doc.json"), ginSwagger.DefaultModelsExpandDepth(-1)))

	return router
}
