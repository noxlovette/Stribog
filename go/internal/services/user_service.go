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

var (
	ErrPasswordTooShort     = errors.New("password too short")
	ErrEmailTaken           = errors.New("email already in use")
	ErrAuthenticationFailed = errors.New("authentication failed")
	ErrEmailNotFound        = errors.New("email not found")
)

type UserService struct {
	queries *db.Queries
}

func NewUserService(queries *db.Queries) *UserService {
	return &UserService{queries}
}

func (s *UserService) RegisterUser(ctx context.Context, req types.SignupRequest) (uuid.UUID, error) {
	if len(req.Password) < 8 {
		return uuid.Nil, ErrPasswordTooShort
	}

	exists, err := s.queries.CheckEmailExists(ctx, req.Email)
	if err != nil {
		return uuid.Nil, fmt.Errorf("checking email existence: %w", err)
	}
	if exists {
		return uuid.Nil, ErrEmailTaken
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return uuid.Nil, fmt.Errorf("hashing password: %w", err)
	}

	id, err := s.queries.CreateUser(ctx, db.CreateUserParams{
		Name:         types.ToPgText(req.Name),
		Email:        req.Email,
		PasswordHash: string(hash),
	})
	if err != nil {
		return uuid.Nil, fmt.Errorf("creating user: %w", err)
	}

	return id, nil
}

func (s *UserService) Login(ctx context.Context, req types.LoginRequest) (types.WebUser, error) {
	user, err := s.queries.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return types.WebUser{}, fmt.Errorf("getting user by email: %w", err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		return types.WebUser{}, ErrAuthenticationFailed
	}

	return types.WebUser{
		Email: user.Email,
		Name:  types.PgTextToString(user.Name),
	}, nil
}
