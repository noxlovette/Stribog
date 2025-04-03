package db

import (
	"context"
	"fmt"

	"stribog/internal/tools"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Pool struct {
	*pgxpool.Pool
}

func Init(ctx context.Context) (*Pool, error) {
	dbpool, err := pgxpool.New(ctx, tools.GetEnvVar("DATABASE_URL"))
	if err != nil {
		return nil, fmt.Errorf("unable to create connection pool: %v", err)
	}

	if err := dbpool.Ping(ctx); err != nil {
		dbpool.Close()
		return nil, fmt.Errorf("unable to ping database: %v", err)
	}

	return &Pool{Pool: dbpool}, nil
}
