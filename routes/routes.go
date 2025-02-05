package routes

import (
	"marketplace/backend/controllers"
	"marketplace/backend/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	// Authentication routes
	auth := router.Group("/api/auth")
	{
		auth.POST("/register", controllers.Register)
		auth.POST("/login", controllers.Login)
	}

	// Game routes
	games := router.Group("/api/games")
	{
		games.GET("/", controllers.GetGames)
	}

	// Cart routes (protected by JWT middleware)
	cart := router.Group("/api/cart").Use(middleware.JWTAuth())
	{
		cart.POST("/", controllers.AddToCart)
		cart.DELETE("/:id", controllers.RemoveFromCart)
	}

	// Order routes (protected by JWT middleware)
	orders := router.Group("/api/orders").Use(middleware.JWTAuth())
	{
		orders.POST("/", controllers.PlaceOrder)
		orders.GET("/", controllers.GetOrders)
	}
}
