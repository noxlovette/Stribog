package middleware

import (
	"crypto/sha256"
	"encoding/base64"
	"stribog/internal/services"
	"strings"

	"github.com/gin-gonic/gin"
)

func APIKeyAuthMiddleware(svc *services.APIKeyService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(401, gin.H{"error": "missing API key"})
			return
		}

		key := strings.TrimPrefix(authHeader, "Bearer ")
		hash := sha256.Sum256([]byte(key))
		keyHash := base64.URLEncoding.EncodeToString(hash[:])

		keyID, err := svc.ValidateKey(c.Request.Context(), keyHash)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": "invalid API key"})
			return
		}

		c.Set("api_key_id", keyID)
		c.Next()
	}
}
