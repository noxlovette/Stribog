package db

import (
	"context"
	"fmt"
	"time"

	"stribog/internal/tools"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Pool struct {
	*pgxpool.Pool
}

func Init(ctx context.Context) (*Pool, error) {
	dbpool, err := pgxpool.New(ctx, tools.GetEnvVar("DATABASE_URL"))
	if err != nil {
		return nil, fmt.Errorf("unable to create connection pool: %w", err)
	}

	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	return &Pool{Pool: dbpool}, nil
}
