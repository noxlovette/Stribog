package middleware

import (
	"context"
	"net/http"
	"stribog/internal/auth"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddlewareCookie(tokenSvc auth.TokenService) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr, err := c.Cookie("access_token")
		if err != nil || tokenStr == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing access token"})
			return
		}

		userID, err := tokenSvc.ParseToken(tokenStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}

		ctx := context.WithValue(c.Request.Context(), auth.UserIDKey, userID)
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}
func AuthMiddlewareHeader(tokenSvc auth.TokenService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing or invalid authorization header"})
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		userID, err := tokenSvc.ParseToken(tokenStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}

		ctx := context.WithValue(c.Request.Context(), auth.UserIDKey, userID)
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}
