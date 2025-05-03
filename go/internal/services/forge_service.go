package services

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	db "stribog/internal/db/sqlc"
	appError "stribog/internal/errors"

	"stribog/internal/middleware"
	"stribog/internal/types"

	"github.com/aidarkhanov/nanoid"
	"github.com/google/uuid"
)

type ForgeService struct {
	querier db.Querier
}

func NewForgeService(q db.Querier) *ForgeService {
	return &ForgeService{
		querier: q,
	}
}

var (
	ErrInvalidForgeID    = fmt.Errorf("invalid forge ID")
	ErrCreateForgeFailed = fmt.Errorf("failed to create forge")
	ErrListForgesFailed  = fmt.Errorf("failed to list forges")
)

func (s *ForgeService) CreateForge(ctx context.Context, create types.ForgeCreateRequest) (*string, error) {
	userIDStr, ok := ctx.Value(middleware.UserIDKey).(string)
	if !ok {
		return nil, fmt.Errorf("%w: user ID missing or not a string", appError.ErrInvalidUserId)
	}

	parsedUserID, err := uuid.Parse(userIDStr)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", appError.ErrInvalidUserId, err)
	}

	forgeID := nanoid.New()

	err = s.querier.InsertForge(ctx, db.InsertForgeParams{
		ID:          forgeID,
		OwnerID:     parsedUserID,
		Title:       create.Title,
		Description: create.Description,
	})

	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrCreateForgeFailed, err)
	}

	return &forgeID, nil
}

func (s *ForgeService) ListForges(ctx context.Context) ([]*types.WebForge, error) {
	userIDStr, ok := ctx.Value(middleware.UserIDKey).(string)
	if !ok {
		return nil, fmt.Errorf("%w: user ID missing or not a string", appError.ErrInvalidUserId)
	}

	parsedUserID, err := uuid.Parse(userIDStr)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", appError.ErrInvalidUserId, err)
	}

	forges, err := s.querier.GetForgesAndCheckReadAccess(ctx, parsedUserID)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrListForgesFailed, err)
	}

	webForges := make([]*types.WebForge, len(forges))
	for i, forge := range forges {
		webForges[i] = &types.WebForge{
			ID:          forge.ID,
			Title:       forge.Title,
			Description: *forge.Description,
		}
	}

	return webForges, nil
}

func (s *ForgeService) GetForge(ctx context.Context, forgeID string) (*types.WebForge, error) {
	userIDStr, ok := ctx.Value(middleware.UserIDKey).(string)
	if !ok {
		return nil, fmt.Errorf("%w: user ID missing or not a string", appError.ErrInvalidUserId)
	}

	parsedUserID, err := uuid.Parse(userIDStr)
	if err != nil {
		return nil, fmt.Errorf("failed to parse user ID: %w", err)
	}

	forge, err := s.querier.GetForgeAndCheckReadAccess(ctx, db.GetForgeAndCheckReadAccessParams{
		OwnerID: parsedUserID,
		ID:      forgeID,
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("%w: forge not found", appError.ErrNotFound)
		}
		return nil, fmt.Errorf("failed to get forge: %w", err)
	}

	return &types.WebForge{
		ID:          forge.ID,
		Title:       forge.Title,
		Description: *forge.Description,
	}, nil
}

func (s *ForgeService) DeleteForge(ctx context.Context, forgeID string) error {
	userIDStr, ok := ctx.Value(middleware.UserIDKey).(string)
	if !ok {
		return fmt.Errorf("%w: user ID missing or not a string", appError.ErrInvalidUserId)
	}

	parsedUserID, err := uuid.Parse(userIDStr)
	if err != nil {
		return fmt.Errorf("%w: %v", appError.ErrInvalidUserId, err)
	}

	err = s.querier.DeleteForge(ctx, db.DeleteForgeParams{
		OwnerID: parsedUserID,
		ID:      forgeID,
	})
	if err != nil {
		return fmt.Errorf("failed to delete forge: %w", err)
	}

	return nil
}

func (s *ForgeService) UpdateForge(ctx context.Context, update types.ForgeUpdateRequest, forgeID string) error {
	userIDStr, ok := ctx.Value(middleware.UserIDKey).(string)
	if !ok {
		return fmt.Errorf("%w: user ID missing or not a string", appError.ErrInvalidUserId)
	}

	parsedUserID, err := uuid.Parse(userIDStr)
	if err != nil {
		return fmt.Errorf("%w: %v", appError.ErrInvalidUserId, err)
	}

	err = s.querier.UpdateForgeAndCheckWriteAccess(ctx, db.UpdateForgeAndCheckWriteAccessParams{
		OwnerID:     parsedUserID,
		ID:          forgeID,
		Title:       update.Title,
		Description: update.Description,
	})
	if err != nil {
		return fmt.Errorf("failed to update forge: %w", err)
	}

	return nil
}
