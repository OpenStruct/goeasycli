package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"hjo/config"
)

type HealthRepo struct{}

// Health godoc
// @Summary Check the health of the application
// @Description Check the health of the application
// @Tags health
// @Router /health [get]
func (repo *HealthRepo) Health(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status":      "I am breathing!",
		"environment": config.CFG.V.GetString("APP_ENV"),
	})
}
