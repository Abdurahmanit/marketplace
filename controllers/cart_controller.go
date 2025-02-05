package controllers

import (
	"marketplace/backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddToCart(c *gin.Context) {
	var item struct {
		GameID string `json:"game_id"`
	}
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Add game to cart
	userID := c.GetString("user_id")
	err := services.AddGameToCart(userID, item.GameID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Game added to cart"})
}

func RemoveFromCart(c *gin.Context) {
	gameID := c.Param("id")

	// Remove game from cart
	userID := c.GetString("user_id")
	err := services.RemoveGameFromCart(userID, gameID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Game removed from cart"})
}
