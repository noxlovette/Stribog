package services

import (
	"context"
	"fmt"
	db "stribog/internal/db/sqlc"
	appError "stribog/internal/errors"
	"stribog/internal/middleware"
	"stribog/internal/types"

	"github.com/google/uuid"
)

type AccessService struct {
	querier db.Querier
}

func NewAccessService(q db.Querier) *AccessService {
	return &AccessService{
		querier: q,
	}
}

var (
	ErrAccessDenied      = fmt.Errorf("access denied")
	ErrCheckUserAccess   = fmt.Errorf("error checking user access")
	ErrListForgeAccess   = fmt.Errorf("error listing forge access")
	ErrAddForgeAccess    = fmt.Errorf("error adding forge access")
	ErrRemoveForgeAccess = fmt.Errorf("error removing forge access")
)

func (s *AccessService) ListForgeAccess(ctx context.Context, userID string, forgeID string) ([]*types.WebAccess, error) {
	userIDStr, ok := ctx.Value(middleware.UserIDKey).(string)
	if !ok {
		return nil, fmt.Errorf("%w: user ID missing or not a string", appError.ErrInvalidUserId)
	}

	parsedUserID, err := uuid.Parse(userIDStr)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", appError.ErrInvalidUserId, err)
	}

	hasAccess, err := s.querier.CheckReadAccess(ctx, db.CheckReadAccessParams{
		OwnerID: parsedUserID,
		ID:      forgeID,
	})
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrCheckUserAccess, err)
	}

	if !hasAccess {
		return nil, fmt.Errorf("%w: user does not have access to the forge", ErrAccessDenied)
	}

	accessList, err := s.querier.ListForgeAccess(ctx, forgeID)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrListForgeAccess, err)
	}

	webAccessList := make([]*types.WebAccess, len(accessList))
	for i, al := range accessList {
		webAccessList[i] = &types.WebAccess{
			ID:         al.ID,
			ForgeID:    al.ForgeID,
			UserID:     al.UserID,
			AccessRole: al.AccessRole,
			UserName:   *al.UserName,
			UserEmail:  al.UserEmail,
		}
	}

	return webAccessList, nil
}

func (s *AccessService) AddForgeAccess(ctx context.Context, userID string, forgeID string, accessRole db.AccessRole) error {
	userIDStr, ok := ctx.Value(middleware.UserIDKey).(string)
	if !ok {
		return fmt.Errorf("%w: user ID missing or not a string", appError.ErrInvalidUserId)
	}

	adderUserID, err := uuid.Parse(userIDStr)
	if err != nil {
		return fmt.Errorf("%w: %v", appError.ErrInvalidUserId, err)
	}

	addedUserID, err := uuid.Parse(userID)
	if err != nil {
		return fmt.Errorf("%w: %v", appError.ErrInvalidUserId, err)
	}

	hasAccess, err := s.querier.CheckWriteAccess(ctx, db.CheckWriteAccessParams{
		OwnerID: adderUserID,
		ID:      forgeID,
	})
	if err != nil {
		return fmt.Errorf("%w: %v", ErrCheckUserAccess, err)
	}

	if !hasAccess {
		return fmt.Errorf("%w: user does not have access to the forge", ErrAccessDenied)
	}

	err = s.querier.UpsertForgeAccess(ctx, db.UpsertForgeAccessParams{
		ForgeID:    forgeID,
		UserID:     addedUserID,
		AccessRole: accessRole,
		AddedBy:    adderUserID,
	})
	if err != nil {
		return fmt.Errorf("%w: %v", ErrAddForgeAccess, err)
	}

	return nil
}

func (s *AccessService) RemoveForgeAccess(ctx context.Context, userID string, forgeID string) error {
	userIDStr, ok := ctx.Value(middleware.UserIDKey).(string)
	if !ok {
		return fmt.Errorf("%w: user ID missing or not a string", appError.ErrInvalidUserId)
	}

	userIDUUID, err := uuid.Parse(userIDStr)
	if err != nil {
		return fmt.Errorf("%w: %v", appError.ErrInvalidUserId, err)
	}

	hasAccess, err := s.querier.CheckWriteAccess(ctx, db.CheckWriteAccessParams{
		OwnerID: userIDUUID,
		ID:      forgeID,
	})
	if err != nil {
		return fmt.Errorf("%w: %v", ErrCheckUserAccess, err)
	}

	if !hasAccess {
		return fmt.Errorf("%w: user does not have access to the forge", ErrAccessDenied)
	}

	err = s.querier.DeleteForgeAccess(ctx, db.DeleteForgeAccessParams{
		ForgeID: forgeID,
		UserID:  userIDUUID,
	})
	if err != nil {
		return fmt.Errorf("%w: %v", ErrRemoveForgeAccess, err)
	}

	return nil
}
