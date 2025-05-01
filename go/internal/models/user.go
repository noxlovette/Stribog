package models

import (
	"github.com/google/uuid"
)

type WebUser struct {
	ID    uuid.UUID `json:"id"`
	Email string    `json:"email"`
	Name  string    `json:"name"`
}
