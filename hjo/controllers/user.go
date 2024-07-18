package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"hjo/database"
	"hjo/models"
)

type UserRepo struct{}

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user
// @Tags user
// @Accept  json
// @Produce  json
// @Param user body models.GoEasyCLITestUser true "User object"
// @Success 200 {object} models.GoEasyCLITestUser
// @Router /user/add [post]
func (repo *UserRepo) CreateUser(ctx *gin.Context) {
	var user models.GoEasyCLITestUser
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

// GetAllUsers godoc
// @Summary Get all users
// @Description Get all users
// @Tags user
// @Produce  json
// @Success 200 {object} []models.GoEasyCLITestUser
// @Router /user/list [get]
func (repo *UserRepo) GetAllUsers(ctx *gin.Context) {
	users, err := models.GetAllUsers(ctx, database.Db)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

// GetUserByID godoc
// @Summary Get user by ID
// @Description Get user by ID
// @Tags user
// @Produce  json
// @Param id path string true "User ID"
// @Success 200 {object} models.GoEasyCLITestUser
// @Router /user/{id} [get]
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

// UpdateUser godoc
// @Summary Update user
// @Description Update user
// @Tags user
// @Accept  json
// @Produce  json
// @Param user body models.GoEasyCLITestUser true "User object"
// @Success 200 {object} models.GoEasyCLITestUser
// @Router /user/update [patch]
func (repo *UserRepo) UpdateUser(ctx *gin.Context) {
	var user models.GoEasyCLITestUser

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if user.ID == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "user id missing"})
		return
	}

	if err := models.UpdateUser(ctx, database.Db, &user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

// DeleteUser godoc
// @Summary Delete user
// @Description Delete user
// @Tags user
// @Param id path string true "User ID"
// @Success 200 {string} string	"User deleted successfully"
// @Router /user/delete/{id} [delete]
func (repo *UserRepo) DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := models.DeleteUser(ctx, database.Db, id); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
