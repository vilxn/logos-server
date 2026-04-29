package middleware

import (
	"dot/auth"
	"strings"

	"github.com/gin-gonic/gin"
)

// ВСЕГДА ИСПОЛЬЗОВАТЬ c.Abort()

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(401, gin.H{"error": "missing token"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		userClaims, err := auth.ValidateToken(tokenString)
		if err != nil {
			c.JSON(401, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}

		// сохраняем в context
		c.Set("userID", userClaims.ID)
		c.Set("role", userClaims.Role)

		c.Next()
	}
}
