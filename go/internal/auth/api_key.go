package auth

import (
	"crypto/rand"
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"strings"

	"github.com/gin-gonic/gin"
)

func GenerateAPIKey() (string, string) {
	b := make([]byte, 32)
	rand.Read(b)
	apiKey := base64.URLEncoding.EncodeToString(b)
	hash := sha256.Sum256([]byte(apiKey))
	keyHash := base64.URLEncoding.EncodeToString(hash[:])
	return apiKey, keyHash
}

func APIKeyAuthMiddleware(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(401, gin.H{"error": "missing API key"})
			return
		}

		key := strings.TrimPrefix(authHeader, "Bearer ")
		hash := sha256.Sum256([]byte(key))
		keyHash := base64.URLEncoding.EncodeToString(hash[:])

		var keyID string
		err := db.QueryRow(`SELECT id FROM api_keys WHERE key_hash = $1 AND is_active = TRUE`, keyHash).Scan(&keyID)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": "invalid API key"})
			return
		}

		// attach context info if needed
		c.Set("api_key_id", keyID)
		c.Next()
	}
}
