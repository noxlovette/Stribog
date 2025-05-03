package auth

import (
	"context"

	"github.com/google/uuid"
)

type ctxKey string

const UserIDKey ctxKey = "userID"

func GetUserID(ctx context.Context) (uuid.UUID, bool) {
	userID, ok := ctx.Value(UserIDKey).(uuid.UUID)
	return userID, ok
}
