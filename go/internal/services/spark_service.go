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
)

type SparkService struct {
	querier db.Querier
}

func NewSparkService(q db.Querier) *SparkService {
	return &SparkService{
		querier: q,
	}
}

var (
	ErrInvalidSparkID    = fmt.Errorf("invalid spark ID")
	ErrCreateSparkFailed = fmt.Errorf("failed to create spark")
	ErrListSparksFailed  = fmt.Errorf("failed to list sparks")
)

func (s *SparkService) CreateSpark(ctx context.Context, forgeID string) (*string, error) {
	userID, ok := middleware.GetUserID(ctx)
	if !ok {
		return nil, fmt.Errorf("%w: user ID missing or not a UUID", appError.ErrInvalidUserId)
	}

	ok, err := s.querier.CheckWriteAccess(ctx, db.CheckWriteAccessParams{
		OwnerID: userID,
		ID:      forgeID,
	})
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrCreateSparkFailed, err)
	}
	if !ok {
		return nil, fmt.Errorf("%w: user does not have write access to forge", appError.ErrForbidden)
	}

	sparkID := nanoid.New()

	err = s.querier.InsertSpark(ctx, db.InsertSparkParams{
		ID:       sparkID,
		OwnerID:  userID,
		ForgeID:  forgeID,
		Title:    "New Spark",
		Markdown: "Markdown",
	})

	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrCreateForgeFailed, err)
	}

	return &sparkID, nil
}

func (s *SparkService) ListSparksByForgeID(ctx context.Context, forgeID string) ([]*types.WebSpark, error) {
	userID, ok := middleware.GetUserID(ctx)
	if !ok {
		return nil, fmt.Errorf("%w: user ID missing or not a UUID", appError.ErrInvalidUserId)
	}

	forges, err := s.querier.GetSparksByForgeIDAndCheckReadAccess(ctx, db.GetSparksByForgeIDAndCheckReadAccessParams{
		OwnerID: userID,
		ForgeID: forgeID,
	})
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrListForgesFailed, err)
	}

	webForges := make([]*types.WebSpark, len(forges))
	for i, forge := range forges {
		webForges[i] = &types.WebSpark{
			ID:       forge.ID,
			Title:    forge.Title,
			Markdown: forge.Markdown,
		}
	}

	return webForges, nil
}

func (s *SparkService) GetSpark(ctx context.Context, sparkID string) (*types.WebSpark, error) {
	userID, ok := middleware.GetUserID(ctx)
	if !ok {
		return nil, fmt.Errorf("%w: user ID missing or not a UUID", appError.ErrInvalidUserId)
	}

	spark, err := s.querier.GetSparkAndCheckReadAccess(ctx, db.GetSparkAndCheckReadAccessParams{
		OwnerID: userID,
		ID:      sparkID,
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("%w: forge not found", appError.ErrNotFound)
		}
		return nil, fmt.Errorf("failed to get forge: %w", err)
	}

	return &types.WebSpark{
		ID:       spark.ID,
		Title:    spark.Title,
		Markdown: spark.Markdown,
	}, nil
}

func (s *SparkService) DeleteSpark(ctx context.Context, sparkID string) error {
	userID, ok := middleware.GetUserID(ctx)
	if !ok {
		return fmt.Errorf("%w: user ID missing or not a UUID", appError.ErrInvalidUserId)
	}

	err := s.querier.DeleteSparkAndCheckAdminAccess(ctx, db.DeleteSparkAndCheckAdminAccessParams{
		OwnerID: userID,
		ID:      sparkID,
	})
	if err != nil {
		return fmt.Errorf("failed to delete forge: %w", err)
	}

	return nil
}

func (s *SparkService) UpdateSpark(ctx context.Context, update types.SparkUpdateRequest, forgeID string) error {
	userID, ok := middleware.GetUserID(ctx)
	if !ok {
		return fmt.Errorf("%w: user ID missing or not a UUID", appError.ErrInvalidUserId)
	}

	err := s.querier.UpdateSparkAndCheckWriteAccess(ctx, db.UpdateSparkAndCheckWriteAccessParams{
		OwnerID:  userID,
		ID:       forgeID,
		Title:    update.Title,
		Markdown: update.Markdown,
	})
	if err != nil {
		return fmt.Errorf("failed to update spark: %w", err)
	}

	return nil
}
