package controllers

import (
	"github.com/Abdurahmanit/marketplace/backend/services"

	"net/http"

	"github.com/gin-gonic/gin"
)

func PlaceOrder(c *gin.Context) {
	userID := c.GetString("user_id")
	err := services.PlaceOrder(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order placed successfully"})
}

func GetOrders(c *gin.Context) {
	userID := c.GetString("user_id")
	orders, err := services.GetUserOrders(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, orders)
}
