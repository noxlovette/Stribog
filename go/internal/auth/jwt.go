package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type TokenService interface {
	GenerateToken(userID string, ttl time.Duration) (string, time.Time, error)
	ParseToken(tokenStr string) (uuid.UUID, error)
}

type JWTAuth struct {
	secret []byte
}

func NewJWTAuth(secret string) *JWTAuth {
	return &JWTAuth{secret: []byte(secret)}
}

func (j *JWTAuth) GenerateToken(userID string, ttl time.Duration) (string, time.Time, error) {
	expiration := time.Now().Add(ttl)
	claims := jwt.RegisteredClaims{
		Subject:   userID,
		ExpiresAt: jwt.NewNumericDate(expiration),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(j.secret)
	if err != nil {
		return "", time.Time{}, err
	}
	return signedToken, expiration, nil
}

func (j *JWTAuth) ParseToken(tokenStr string) (uuid.UUID, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.secret, nil
	})
	if err != nil {
		return uuid.UUID{}, err
	}

	claims, ok := token.Claims.(*jwt.RegisteredClaims)
	if !ok || !token.Valid {
		return uuid.UUID{}, errors.New("invalid token")
	}

	parsedUUID, err := uuid.Parse(claims.Subject)
	if err != nil {
		return uuid.UUID{}, errors.New("subject is not a valid UUID")
	}

	return parsedUUID, nil
}
