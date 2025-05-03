package auth

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
)

func GenerateAPIKey() (string, string) {
	b := make([]byte, 32)
	rand.Read(b)
	apiKey := base64.URLEncoding.EncodeToString(b)
	hash := sha256.Sum256([]byte(apiKey))
	keyHash := base64.URLEncoding.EncodeToString(hash[:])
	return apiKey, keyHash
}
