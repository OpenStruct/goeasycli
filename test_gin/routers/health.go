package routers

import (
	"github.com/gin-gonic/gin"
	"test_gin/controllers"
)

func HealthRoute(router *gin.Engine) {
	status := new(controllers.HealthRepo)
	router.GET("/health", status.Health)
}
