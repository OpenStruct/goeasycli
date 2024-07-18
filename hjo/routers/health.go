package routers

import (
	"github.com/gin-gonic/gin"
	"hjo/controllers"
)

func HealthRoute(router *gin.Engine) {
	status := new(controllers.HealthRepo)
	router.GET("/health", status.Health)
}
