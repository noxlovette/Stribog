package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTGenerator struct {
	secret []byte
}

func NewJWTGenerator(secret string) *JWTGenerator {
	return &JWTGenerator{secret: []byte(secret)}
}

func (j *JWTGenerator) GenerateToken(userID string, ttl time.Duration) (string, error) {
	claims := jwt.RegisteredClaims{
		Subject:   userID,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(ttl)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString(j.secret)
	if err != nil {
		return "", err
	}

	return signed, nil
}
