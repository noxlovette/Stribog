package config

import (
	"context"
	"stribog/internal/auth"
	"stribog/internal/db"
	"stribog/internal/utils"

	"go.uber.org/zap"
)

type App struct {
	DB           *db.Pool
	Logger       *zap.Logger
	JWTValidator *auth.JWTValidator
	JWTGenerator *auth.JWTGenerator
	Env          *Env
}

func InitAppState(ctx context.Context) (*App, error) {

	env := &Env{
		JWTKey:      utils.GetEnvVar("JWT_KEY"),
		DatabaseDSN: utils.GetEnvVar("DATABASE_DSN"),
	}

	logger, err := zap.NewDevelopment()
	if err != nil {
		return nil, err
	}

	dbPool, err := db.InitDB(ctx, env.DatabaseDSN)
	if err != nil {
		return nil, err
	}

	jwtValidator := auth.NewJWTValidator(env.JWTKey)
	jwtGenerator := auth.NewJWTGenerator(env.JWTKey)

	return &App{
		DB:           dbPool,
		Logger:       logger,
		JWTValidator: jwtValidator,
		JWTGenerator: jwtGenerator,
		Env:          env,
	}, nil
}
