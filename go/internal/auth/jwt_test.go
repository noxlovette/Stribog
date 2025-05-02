package auth

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

func TestGenerateAndParseToken(t *testing.T) {
	auth := NewJWTAuth("mysecretkey")
	userID := "user123"
	ttl := time.Minute * 5

	token, err := auth.GenerateToken(userID, ttl)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	parsedUserID, err := auth.ParseToken(token)
	assert.NoError(t, err)
	assert.Equal(t, userID, parsedUserID)
}

func TestParseToken_InvalidSignature(t *testing.T) {
	auth1 := NewJWTAuth("key1")
	auth2 := NewJWTAuth("key2")
	userID := "user456"
	ttl := time.Minute * 5

	token, err := auth1.GenerateToken(userID, ttl)
	assert.NoError(t, err)

	parsedUserID, err := auth2.ParseToken(token)
	assert.Error(t, err)
	assert.Empty(t, parsedUserID)
}

func TestParseToken_Expired(t *testing.T) {
	auth := NewJWTAuth("mysecretkey")
	userID := "user789"
	ttl := time.Second * -1 // Already expired

	token, err := auth.GenerateToken(userID, ttl)
	assert.NoError(t, err)

	time.Sleep(time.Millisecond * 10) // Just to ensure token is read as expired

	parsedUserID, err := auth.ParseToken(token)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), jwt.ErrTokenExpired.Error())
	assert.Empty(t, parsedUserID)
}

func TestParseToken_InvalidFormat(t *testing.T) {
	auth := NewJWTAuth("mysecretkey")

	_, err := auth.ParseToken("not.a.jwt.token")
	assert.Error(t, err)
}
