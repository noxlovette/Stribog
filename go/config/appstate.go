package config

import (
	"context"
	"fmt"
	"stribog/internal/db"
)

type AppState struct {
	DB *db.Pool
}

func InitAppState(ctx context.Context) (*AppState, error) {
	dbPool, err := db.Init(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize database: %w", err)
	}

	return &AppState{
		DB: dbPool,
	}, nil
}
