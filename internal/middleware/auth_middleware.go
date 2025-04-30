package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rest-project/internal/auth"
	"strings"
)

func AuthRequired(requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing or malformed"})
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		// Валидируем токен
		_, claims, err := auth.ValidateJWT(tokenStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token", "details": err.Error()})
			return
		}

		// Проверка существования и типа user_id в claims
		userIDFloat, ok := claims["user_id"].(float64)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid user ID in token payload"})
			return
		}

		roleClaim, ok := claims["role"].(string)
		if !ok || roleClaim != requiredRole {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Insufficient permissions", "required_role": requiredRole, "user_role": roleClaim})
			return
		}

		c.Set("userID", uint(userIDFloat))
		c.Set("role", roleClaim)

		c.Next()
	}
}
