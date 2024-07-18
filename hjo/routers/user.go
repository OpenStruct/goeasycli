package routers

import (
	"github.com/gin-gonic/gin"
	"hjo/controllers"
	"hjo/middlewares"
)

func UserRoutes(userGroup *gin.RouterGroup) {
	user := new(controllers.UserRepo)
	userGroup.POST("/add", user.CreateUser)
	userGroup.GET("/list", user.GetAllUsers)
	userGroup.GET("/:id", user.GetUserByID)
	userGroup.DELETE("/delete/:id", middlewares.TestAuthMiddleware(), user.DeleteUser)
	userGroup.PATCH("/update", user.UpdateUser)
}
