package main

import (
	"marketplace/backend/config"
	"marketplace/backend/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize MongoDB connection
	config.ConnectDB()

	// Create a Gin router
	router := gin.Default()

	// Set up routes
	routes.SetupRoutes(router)

	// Start the server
	router.Run(":8080")
}
