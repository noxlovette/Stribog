package types

import "github.com/google/uuid"

type WebAPIKey struct {
	ID       uuid.UUID `json:"id"`
	Title    string    `json:"title"`
	IsActive bool      `json:"is_active"`
}
type CreateAPIKey struct {
	Title string `json:"title"`
}

type APIKeyID struct {
	KeyID uuid.UUID `json:"key_id"`
}
