package auth

import "time"

type MockTokenService struct {
	TokenToReturn string
	ErrToReturn   error
	ParsedUserID  string
}

func (m *MockTokenService) GenerateToken(userID string, ttl time.Duration) (string, error) {
	return m.TokenToReturn, m.ErrToReturn
}

func (m *MockTokenService) ParseToken(tokenStr string) (string, error) {
	return m.ParsedUserID, m.ErrToReturn
}
