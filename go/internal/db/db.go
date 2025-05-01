package db

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Pool struct {
	*pgxpool.Pool
}

func InitDB(ctx context.Context, dbDSN string) (*Pool, error) {
	dbpool, err := pgxpool.New(ctx, dbDSN)
	if err != nil {
		return nil, fmt.Errorf("unable to create connection pool: %w", err)
	}

	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	return &Pool{Pool: dbpool}, nil
}
