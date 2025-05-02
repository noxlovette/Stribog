package middleware

import (
	"context"
	"net/http"
	"stribog/internal/auth"

	"github.com/gin-gonic/gin"
)

type ctxKey string

const UserIDKey ctxKey = "userID"

func GetUserID(ctx context.Context) (string, bool) {
	userID, ok := ctx.Value(UserIDKey).(string)
	return userID, ok
}

func AuthMiddleware(tokenSvc auth.TokenService) gin.HandlerFunc {
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

		// inject into context
		ctx := context.WithValue(c.Request.Context(), UserIDKey, userID)
		c.Request = c.Request.WithContext(ctx)

		// proceed to next middleware/handler
		c.Next()
	}
}
