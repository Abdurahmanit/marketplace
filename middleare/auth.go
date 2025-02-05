package middleware

import (
	"net/http"

	"github.com/Abdurahmanit/marketplace/backend/utils"
	"github.com/gin-gonic/gin"
)

// JWTAuthMiddleware проверяет JWT токен и добавляет user_id в контекст
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Получаем токен из заголовка Authorization
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token is required"})
			c.Abort()
			return
		}

		// Проверяем и декодируем токен
		claims, err := utils.ParseJWT(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Добавляем user_id в контекст
		userID := claims["user_id"].(string)
		c.Set("user_id", userID)

		// Передаем управление следующему обработчику
		c.Next()
	}
}
