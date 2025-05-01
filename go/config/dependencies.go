package config

import (
	"context"
	"stribog/internal/db"
)

type Dependencies struct {
	DB *db.Pool
}

func InitAppState(ctx context.Context) (*Dependencies, error) {
	dbPool, err := db.Init(ctx)
	if err != nil {
		return nil, err
	}

	return &Dependencies{
		DB: dbPool,
	}, nil
}
