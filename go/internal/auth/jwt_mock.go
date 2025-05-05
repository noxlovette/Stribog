package auth

import (
	"time"

	"github.com/google/uuid"
)

type MockTokenService struct {
	TokenToReturn string
	ErrToReturn   error
	ParsedUserID  uuid.UUID
}

func (m *MockTokenService) GenerateToken(userID string, ttl time.Duration) (string, error) {
	return m.TokenToReturn, m.ErrToReturn
}

func (m *MockTokenService) ParseToken(tokenStr string) (uuid.UUID, error) {
	return m.ParsedUserID, m.ErrToReturn
}
