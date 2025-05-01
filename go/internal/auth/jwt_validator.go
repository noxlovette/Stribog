package auth

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

type JWTValidator struct {
	secret []byte
}

func NewJWTValidator(secret string) *JWTValidator {
	return &JWTValidator{secret: []byte(secret)}
}

func (j *JWTValidator) ParseToken(tokenStr string) (*jwt.RegisteredClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.secret, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*jwt.RegisteredClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token claims")
}
