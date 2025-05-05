package services

import (
	"context"
	"fmt"
	"stribog/internal/auth"
	db "stribog/internal/db/sqlc"
	appError "stribog/internal/errors"
	"stribog/internal/types"
)

type APIKeyService struct {
	querier db.Querier
}

func NewAPIKeyService(q db.Querier) *APIKeyService {
	return &APIKeyService{
		querier: q,
	}
}

var (
	ErrListAPIKey   = fmt.Errorf("error listing api key")
	ErrAddAPIKey    = fmt.Errorf("error adding api key")
	ErrRemoveAPIKey = fmt.Errorf("error removing api key")
)

func (s *APIKeyService) ListAPIKeys(ctx context.Context, forgeID string) ([]*types.WebAPIKey, error) {
	userID, ok := auth.GetUserID(ctx)
	if !ok {
		return nil, fmt.Errorf("%w: user ID missing or not a UUID", appError.ErrInvalidUserId)
	}

	hasAccess, err := s.querier.CheckAdminAccess(ctx, db.CheckAdminAccessParams{
		OwnerID: userID,
		ID:      forgeID,
	})
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrCheckUserAccess, err)
	}

	if !hasAccess {
		return nil, fmt.Errorf("%w: user does not have access to the forge", ErrAccessDenied)
	}

	apiKeys, err := s.querier.GetAPIKeysByForgeID(ctx, forgeID)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrListAPIKey, err)
	}

	webKeys := make([]*types.WebAPIKey, len(apiKeys))
	for i, key := range apiKeys {
		webKeys[i] = &types.WebAPIKey{
			ID:       key.ID,
			Title:    key.Title,
			IsActive: key.IsActive,
		}

	}
	return webKeys, nil
}

func (s *APIKeyService) CreateAPIKey(ctx context.Context, forgeID string, create types.CreateAPIKey) (*string, error) {
	userID, ok := auth.GetUserID(ctx)
	if !ok {
		return nil, fmt.Errorf("%w: user ID missing or not a UUID", appError.ErrInvalidUserId)
	}
	hasAccess, err := s.querier.CheckAdminAccess(ctx, db.CheckAdminAccessParams{
		OwnerID: userID,
		ID:      forgeID,
	})
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrCheckUserAccess, err)
	}

	if !hasAccess {
		return nil, fmt.Errorf("%w: user does not have access to the forge", ErrAccessDenied)
	}

	apiKey, keyHash := auth.GenerateAPIKey()

	err = s.querier.InsertAPIKey(ctx, db.InsertAPIKeyParams{
		ForgeID: forgeID,
		Title:   create.Title,
		KeyHash: keyHash,
	})
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrAddAPIKey, err)
	}

	return &apiKey, nil
}

func (s *APIKeyService) DeleteAPIKey(ctx context.Context, forgeID string, delete types.APIKeyID) error {
	userID, ok := auth.GetUserID(ctx)
	if !ok {
		return fmt.Errorf("%w: user ID missing or not a UUID", appError.ErrInvalidUserId)
	}

	hasAccess, err := s.querier.CheckAdminAccess(ctx, db.CheckAdminAccessParams{
		OwnerID: userID,
		ID:      forgeID,
	})
	if err != nil {
		return fmt.Errorf("%w: %v", ErrCheckUserAccess, err)
	}

	if !hasAccess {
		return fmt.Errorf("%w: user does not have access to the forge", ErrAccessDenied)
	}

	err = s.querier.DeleteAPIKey(ctx, delete.KeyID)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrRemoveAPIKey, err)
	}

	return nil
}
func (s *APIKeyService) ValidateKey(ctx context.Context, keyHash string) (string, error) {
	id, err := s.querier.GetAPIKeyIDByHash(ctx, keyHash)
	if err != nil {
		return "", fmt.Errorf("%w: user does not have access to the forge", ErrAccessDenied)
	}

	return id.String(), nil

}
