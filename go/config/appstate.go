package config

import (
	"context"
	"stribog/internal/db"
	"stribog/internal/services"
)

type AppState struct {
	DB           *db.Pool
	ForgeService *services.ForgeService
}

func InitAppState(ctx context.Context) (*AppState, error) {
	dbPool, err := db.Init(ctx)
	if err != nil {
		return nil, err
	}

	return &AppState{
		DB:           dbPool,
		ForgeService: services.NewForgeService(dbPool),
	}, nil
}
