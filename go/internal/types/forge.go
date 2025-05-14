package types

import (
	db "stribog/internal/db/sqlc"

	"github.com/google/uuid"
)

type WebForge struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
}

type WebAccess struct {
	ID         uuid.UUID     `json:"id,omitempty"`
	ForgeID    string        `json:"forgeID,omitempty"`
	UserID     uuid.UUID     `json:"userID,omitempty"`
	AccessRole db.AccessRole `json:"accessRole,omitempty"`
	UserName   string        `json:"userName,omitempty"`
	UserEmail  string        `json:"userEmail,omitempty"`
}
type ForgeUpdateRequest struct {
	Title       *string `json:"title,omitempty"`
	Description *string `json:"description,omitempty"`
}
type ForgeCreateRequest struct {
	Title       string  `json:"title,omitempty"`
	Description *string `json:"description,omitempty"`
}
type AccessCreateRequest struct {
	UserID     uuid.UUID     `json:"userID,omitempty"`
	AccessRole db.AccessRole `json:"accessRole,omitempty"`
}

type AccessDeleteRequest struct {
	UserID uuid.UUID `json:"userID,omitempty"`
}
