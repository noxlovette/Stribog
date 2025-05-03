package services

import (
	"context"
	"errors"
	"fmt"
	"stribog/internal/auth"
	db "stribog/internal/db/sqlc"
	appError "stribog/internal/errors"
	"stribog/internal/middleware"
	types "stribog/internal/types"
	"strings"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrPasswordTooShort     = errors.New("password too short")
	ErrEmailTaken           = errors.New("email already in use")
	ErrAuthenticationFailed = errors.New("authentication failed")
	ErrEmailNotFound        = errors.New("email not found")
	ErrInvalidRefreshToken  = errors.New("invalid refresh token")
	ErrHashError            = errors.New("error when generating hash")
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

type TokenPair struct {
	AccessToken  string
	RefreshToken string
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
		return uuid.Nil, ErrHashError
	}

	id, err := s.querier.CreateUser(ctx, db.CreateUserParams{
		Name:         &req.Name,
		Email:        req.Email,
		PasswordHash: string(hash),
	})
	if err != nil {
		return uuid.Nil, appError.ErrOperationFailed
	}

	return id, nil
}

func (s *UserService) Login(ctx context.Context, req types.LoginRequest) (*TokenPair, error) {
	user, err := s.querier.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, ErrEmailNotFound
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		return nil, ErrAuthenticationFailed
	}

	accessToken, err := s.tokens.GenerateToken(user.ID.String(), time.Hour)
	if err != nil {
		return nil, err
	}

	refreshToken, err := s.tokens.GenerateToken("refresh:"+user.ID.String(), time.Hour*24*30)
	if err != nil {
		return nil, err
	}

	return &TokenPair{AccessToken: accessToken, RefreshToken: refreshToken}, nil
}

func (s *UserService) Refresh(ctx context.Context, refreshToken string) (string, error) {
	userID, err := s.tokens.ParseToken(refreshToken)
	if err != nil {
		return "", ErrInvalidRefreshToken
	}

	if !strings.HasPrefix(userID, "refresh:") {
		return "", ErrInvalidRefreshToken
	}

	realUserID := strings.TrimPrefix(userID, "refresh:")

	accessToken, err := s.tokens.GenerateToken(realUserID, time.Hour)
	if err != nil {
		return "", appError.ErrOperationFailed
	}

	return accessToken, nil
}

func (s *UserService) GetUser(ctx context.Context) (*types.WebUser, error) {
	userIDStr, ok := ctx.Value(middleware.UserIDKey).(string)
	if !ok {
		return nil, fmt.Errorf("%w: user ID missing or not a string", appError.ErrInvalidUserId)
	}

	parsedUserID, err := uuid.Parse(userIDStr)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", appError.ErrInvalidUserId, err)
	}

	user, err := s.querier.GetUserByID(ctx, parsedUserID)
	if err != nil {
		return nil, appError.ErrInvalidUserId
	}

	return &types.WebUser{
		Email: user.Email,
		Name:  user.Name,
	}, nil
}

func (s *UserService) DeleteUser(ctx context.Context) error {
	userIDStr, ok := ctx.Value(middleware.UserIDKey).(string)
	if !ok {
		return fmt.Errorf("%w: user ID missing or not a string", appError.ErrInvalidUserId)
	}

	parsedUserID, err := uuid.Parse(userIDStr)
	if err != nil {
		return fmt.Errorf("%w: %v", appError.ErrInvalidUserId, err)
	}
	s.querier.DeleteUser(ctx, parsedUserID)

	return nil
}

func (s *UserService) UpdateUser(ctx context.Context, update types.UserUpdateRequest) error {
	userIDStr, ok := ctx.Value(middleware.UserIDKey).(string)
	if !ok {
		return fmt.Errorf("%w: user ID missing or not a string", appError.ErrInvalidUserId)
	}

	parsedUserID, err := uuid.Parse(userIDStr)
	if err != nil {
		return fmt.Errorf("%w: %v", appError.ErrInvalidUserId, err)
	}

	var passwordHash *string
	if update.Password != nil && *update.Password != "" {
		hash, err := bcrypt.GenerateFromPassword([]byte(*update.Password), bcrypt.DefaultCost)
		if err != nil {
			return fmt.Errorf("%w: %v", ErrHashError, err)
		}
		hashed := string(hash)
		passwordHash = &hashed
	}

	arg := db.UpdateUserParams{
		ID:           parsedUserID,
		Name:         update.Name,
		Email:        update.Email,
		PasswordHash: passwordHash,
	}

	err = s.querier.UpdateUser(ctx, arg)
	if err != nil {
		return appError.ErrOperationFailed
	}

	return nil
}
