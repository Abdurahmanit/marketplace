package routes

import (
	"github.com/Abdurahmanit/marketplace/backend/controllers"
	middleware "github.com/Abdurahmanit/marketplace/backend/middleare"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	// Маршруты для аутентификации
	auth := router.Group("/api/auth")
	{
		auth.POST("/register", controllers.Register)
		auth.POST("/login", controllers.Login)
	}

	// Маршруты для игр
	games := router.Group("/api/games")
	{
		games.GET("/", controllers.GetGames)
	}

	// Маршруты для корзины (защищены JWT middleware)
	cart := router.Group("/api/cart").Use(middleware.JWTAuthMiddleware())
	{
		cart.POST("/", controllers.AddToCart)
		cart.DELETE("/:id", controllers.RemoveFromCart)
	}

	// Маршруты для заказов (защищены JWT middleware)
	orders := router.Group("/api/orders").Use(middleware.JWTAuthMiddleware())
	{
		orders.POST("/", controllers.PlaceOrder)
		orders.GET("/", controllers.GetOrders)
	}
}
