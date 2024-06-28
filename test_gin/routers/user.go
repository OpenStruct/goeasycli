package routers

import (
	"github.com/gin-gonic/gin"
	"test_gin/controllers"
)

func UserRoutes(userGroup *gin.RouterGroup) {
	user := new(controllers.UserRepo)
	userGroup.POST("/add", user.CreateUser)
	userGroup.GET("/list", user.GetAllUsers)
	userGroup.GET("/:id", user.GetUserByID)
	userGroup.DELETE("/delete/:id", user.DeleteUser)
	userGroup.PATCH("/update", user.UpdateUser)
}
