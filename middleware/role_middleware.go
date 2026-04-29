package middleware

import (
	"dot/auth"
	"dot/models"
	"strings"

	"github.com/gin-gonic/gin"
)

func RoleMiddleware(roles ...models.UserRole) gin.HandlerFunc {
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

		satisfy := false
		for _, role := range roles {
			if userClaims.Role == role {
				satisfy = true
			}
		}
		if !satisfy {
			c.JSON(401, gin.H{"error": "invalid role"})
			c.Abort()
			return
		}

		c.Set("userID", userClaims.ID)
		c.Set("role", userClaims.Role)

		c.Next()
	}
}
