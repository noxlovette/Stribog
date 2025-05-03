package services

import (
	"context"
	"fmt"
	db "stribog/internal/db/sqlc"
	appError "stribog/internal/errors"
	"stribog/internal/middleware"
	"stribog/internal/types"
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

func (s *AccessService) ListForgeAccess(ctx context.Context, forgeID string) ([]*types.WebAccess, error) {
	userID, ok := middleware.GetUserID(ctx)
	if !ok {
		return nil, fmt.Errorf("%w: user ID missing or not a UUID", appError.ErrInvalidUserId)
	}

	hasAccess, err := s.querier.CheckReadAccess(ctx, db.CheckReadAccessParams{
		OwnerID: userID,
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

func (s *AccessService) CreateForgeAccess(ctx context.Context, forgeID string, create types.AccessCreateRequest) error {
	adderUserID, ok := middleware.GetUserID(ctx)
	if !ok {
		return fmt.Errorf("%w: user ID missing or not a UUID", appError.ErrInvalidUserId)
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
		UserID:     create.UserID,
		AccessRole: create.AccessRole,
		AddedBy:    adderUserID,
	})
	if err != nil {
		return fmt.Errorf("%w: %v", ErrAddForgeAccess, err)
	}

	return nil
}

func (s *AccessService) DeleteForgeAccess(ctx context.Context, forgeID string, delete types.AccessDeleteRequest) error {
	userID, ok := middleware.GetUserID(ctx)
	if !ok {
		return fmt.Errorf("%w: user ID missing or not a UUID", appError.ErrInvalidUserId)
	}

	hasAccess, err := s.querier.CheckWriteAccess(ctx, db.CheckWriteAccessParams{
		OwnerID: userID,
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
		AddedBy: userID,
		UserID:  delete.UserID,
	})
	if err != nil {
		return fmt.Errorf("%w: %v", ErrRemoveForgeAccess, err)
	}

	return nil
}
