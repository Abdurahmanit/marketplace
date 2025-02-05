package controllers

import (
	"marketplace/backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetGames(c *gin.Context) {
	games, err := services.GetAllGames()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, games)
}
