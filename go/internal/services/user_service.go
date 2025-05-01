package services

import (
	"context"
	"errors"
	"fmt"
	db "stribog/internal/db/sqlc"
	types "stribog/internal/types"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	queries *db.Queries
}

func NewUserService(queries *db.Queries) *UserService {
	return &UserService{queries}
}

func (s *UserService) RegisterUser(ctx context.Context, req types.SignupRequest) (uuid.UUID, error) {
	if len(req.Password) < 8 {
		return uuid.Nil, errors.New("password too short")
	}

	exists, err := s.queries.CheckEmailExists(ctx, req.Email)
	if err != nil {
		return uuid.Nil, fmt.Errorf("could not check email: %w", err)
	}
	if exists {
		return uuid.Nil, fmt.Errorf("email taken")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return uuid.Nil, fmt.Errorf("password hashing failed: %w", err)
	}

	id, err := s.queries.CreateUser(ctx, db.CreateUserParams{
		Name: types.ToPgText(req.Name), Email: req.Email, PasswordHash: string(hash),
	})
	if err != nil {
		return uuid.Nil, fmt.Errorf("failed to create user: %w", err)
	}

	return id, nil
}
