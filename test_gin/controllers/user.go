package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"test_gin/database"
	"test_gin/models"
)

type UserRepo struct{}

func (repo *UserRepo) CreateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.CreateUser(ctx, database.Db, &user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}


func (repo *UserRepo) GetAllUsers(ctx *gin.Context) {
	users, err := models.GetAllUsers(ctx, database.Db)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

func (repo *UserRepo) GetUserByID(ctx *gin.Context) {
	id := ctx.Param("id")
	user, err := models.GetUserByID(ctx, database.Db, id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	ctx.JSON(http.StatusOK, user)
}


func (repo *UserRepo) UpdateUser(ctx *gin.Context) {
	var user models.User
	
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if  user.ID == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "user id missing"})
		return
	}

	if err := models.UpdateUser(ctx, database.Db, &user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (repo *UserRepo) DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := models.DeleteUser(ctx, database.Db, id); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}