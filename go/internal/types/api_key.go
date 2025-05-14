package types

import (
	"time"

	"github.com/google/uuid"
)

type WebAPIKey struct {
	ID       uuid.UUID `json:"id"`
	Title    string    `json:"title"`
	IsActive bool      `json:"isActive"`
	CreatedAt time.Time `json:"createdAt"`
	LastUsed time.Time `json:"lastUsed"`
}
type CreateAPIKey struct {
	Title string `json:"title"`
}

type APIKeyID struct {
	KeyID uuid.UUID `json:"keyID"`
}
