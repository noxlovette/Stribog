package services

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	db "stribog/internal/db/sqlc"
	appError "stribog/internal/errors"

	"stribog/internal/types"
)

type PublicService struct {
	querier db.Querier
}

func NewPublicService(q db.Querier) *PublicService {
	return &PublicService{
		querier: q,
	}
}

func (s *PublicService) ListSparksPublic(ctx context.Context, forgeID string) ([]*types.WebSpark, error) {
	sparks, err := s.querier.GetSparksByForgeIDPublic(ctx, forgeID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("%w: forge not found", appError.ErrNotFound)
		}
		return nil, fmt.Errorf("failed to get sparks: %w", err)
	}
	webSparks := make([]*types.WebSpark, len(sparks))
	for i, spark := range sparks {
		webSparks[i] = &types.WebSpark{
			ID:       spark.ID,
			Title:    spark.Title,
			Markdown: spark.Markdown,
			Tags:     spark.Tags,
			Slug:     spark.Slug,
		}

	}

	return webSparks, nil
}
