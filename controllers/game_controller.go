package controllers

import (
	"net/http"

	"github.com/Abdurahmanit/marketplace/backend/services"

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
