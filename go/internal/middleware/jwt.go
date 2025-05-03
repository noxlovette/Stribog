package middleware

import (
	"context"
	"net/http"
	"stribog/internal/auth"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ctxKey string

const UserIDKey ctxKey = "userID"

func GetUserID(ctx context.Context) (uuid.UUID, bool) {
	userID, ok := ctx.Value(UserIDKey).(uuid.UUID)
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

		ctx := context.WithValue(c.Request.Context(), UserIDKey, userID)
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}
