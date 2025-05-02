package services

import (
	"context"
	"errors"
	"fmt"
	"stribog/internal/auth"
	db "stribog/internal/db/sqlc"
	types "stribog/internal/types"
	"time"

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
	querier db.Querier
	tokens  auth.TokenService
}

func NewUserService(q db.Querier, tokens auth.TokenService) *UserService {
	return &UserService{
		querier: q,
		tokens:  tokens,
	}
}

func (s *UserService) RegisterUser(ctx context.Context, req types.SignupRequest) (uuid.UUID, error) {
	if len(req.Password) < 8 {
		return uuid.Nil, ErrPasswordTooShort
	}

	exists, err := s.querier.CheckEmailExists(ctx, req.Email)
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

	id, err := s.querier.CreateUser(ctx, db.CreateUserParams{
		Name:         types.ToPgText(req.Name),
		Email:        req.Email,
		PasswordHash: string(hash),
	})
	if err != nil {
		return uuid.Nil, fmt.Errorf("creating user: %w", err)
	}

	return id, nil
}

func (s *UserService) Login(ctx context.Context, req types.LoginRequest) (string, error) {
	user, err := s.querier.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return "", ErrEmailNotFound
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		return "", ErrAuthenticationFailed
	}

	token, err := s.tokens.GenerateToken(user.ID.String(), time.Hour*24)
	if err != nil {
		return "", err
	}

	return token, nil
}
