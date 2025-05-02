package middleware

import (
	"net/http"
	"stribog/internal/auth"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(tokenSvc auth.TokenService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing bearer token"})
			return
		}
		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		userID, err := tokenSvc.ParseToken(tokenStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}

		c.Set("userID", userID) // put it in context
		c.Next()
	}
}
