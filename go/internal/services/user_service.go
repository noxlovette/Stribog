package services

import (
	"context"
	"errors"
	"fmt"
	"stribog/internal/auth"
	db "stribog/internal/db/sqlc"
	appError "stribog/internal/errors"
	types "stribog/internal/types"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"
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
	logger  *zap.Logger
}

func NewUserService(q db.Querier, tokens auth.TokenService, logger *zap.Logger) *UserService {
	return &UserService{
		querier: q,
		tokens:  tokens,
		logger:  logger,
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

func (s *UserService) Login(ctx context.Context, req types.LoginRequest) (*types.TokenPair, error) {
	user, err := s.querier.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, ErrEmailNotFound
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		return nil, ErrAuthenticationFailed
	}

	accessToken, accessExpiresAt, err := s.tokens.GenerateToken(user.ID.String(), time.Hour)
	if err != nil {
		return nil, err
	}
	s.logger.Info("Generated access token", zap.String("userID", user.ID.String()), zap.Time("expiresAt", accessExpiresAt))

	refreshToken, refreshExpiresAt, err := s.tokens.GenerateToken(user.ID.String(), 24*time.Hour*30)
	if err != nil {
		return nil, err
	}
	s.logger.Info("Generated refresh token", zap.String("userID", user.ID.String()), zap.Time("expiresAt", refreshExpiresAt))

	var tokenPair types.TokenPair
	tokenPair.AccessToken.Token = accessToken
	tokenPair.AccessToken.ExpiresAt = accessExpiresAt
	tokenPair.RefreshToken.Token = refreshToken
	tokenPair.RefreshToken.ExpiresAt = refreshExpiresAt

	return &tokenPair, nil
}

func (s *UserService) Refresh(ctx context.Context, refreshToken string) (*types.TokenPair, error) {
	s.logger.Info("Starting token refresh process")

	userID, err := s.tokens.ParseToken(refreshToken)
	if err != nil {
		s.logger.Warn("Failed to parse refresh token", zap.Error(err))
		return nil, ErrInvalidRefreshToken
	}

	accessToken, accessExpiresAt, err := s.tokens.GenerateToken(userID.String(), time.Hour)
	if err != nil {
		s.logger.Error("Failed to generate access token", zap.Error(err))
		return nil, appError.ErrOperationFailed
	}

	var tokenPair types.TokenPair
	tokenPair.AccessToken.Token = accessToken
	tokenPair.AccessToken.ExpiresAt = accessExpiresAt

	return &tokenPair, nil
}

func (s *UserService) GetUser(ctx context.Context) (*types.WebUser, error) {
	userID, ok := auth.GetUserID(ctx)
	if !ok {
		return nil, fmt.Errorf("%w: user ID missing or not a UUID", appError.ErrInvalidUserId)
	}

	user, err := s.querier.GetUserByID(ctx, userID)
	if err != nil {
		return nil, appError.ErrInvalidUserId
	}

	return &types.WebUser{
		Email: user.Email,
		Name:  user.Name,
	}, nil
}

func (s *UserService) DeleteUser(ctx context.Context) error {
	userID, ok := auth.GetUserID(ctx)
	if !ok {
		return fmt.Errorf("%w: user ID missing or not a UUID", appError.ErrInvalidUserId)
	}

	s.querier.DeleteUser(ctx, userID)

	return nil
}

func (s *UserService) UpdateUser(ctx context.Context, update types.UserUpdateRequest) error {
	userID, ok := auth.GetUserID(ctx)
	if !ok {
		return fmt.Errorf("%w: user ID missing or not a UUID", appError.ErrInvalidUserId)
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
		ID:           userID,
		Name:         update.Name,
		Email:        update.Email,
		PasswordHash: passwordHash,
	}

	err := s.querier.UpdateUser(ctx, arg)
	if err != nil {
		return appError.ErrOperationFailed
	}

	return nil
}
