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
	ForgeID    string        `json:"forge_id,omitempty"`
	UserID     uuid.UUID     `json:"user_id,omitempty"`
	AccessRole db.AccessRole `json:"access_role,omitempty"`
	UserName   string        `json:"user_name,omitempty"`
	UserEmail  string        `json:"user_email,omitempty"`
}
type ForgeUpdateRequest struct {
	Title       *string `json:"title,omitempty"`
	Description *string `json:"description,omitempty"`
}
type ForgeCreateRequest struct {
	Title       string  `json:"title,omitempty"`
	Description *string `json:"description,omitempty"`
}
