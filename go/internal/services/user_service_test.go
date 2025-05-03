package services

import (
	"context"
	"testing"

	"stribog/internal/auth"
	sqlc "stribog/internal/db/sqlc"
	"stribog/internal/db/sqlc/mock"
	appError "stribog/internal/errors"
	"stribog/internal/types"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"golang.org/x/crypto/bcrypt"
)

func TestUserService_RegisterUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockQuerier := mock.NewMockQuerier(ctrl)

	mockTokenSvc := &auth.MockTokenService{
		TokenToReturn: "dummy-token",
		ErrToReturn:   nil,
		ParsedUserID:  uuid.MustParse("11111111-1111-1111-1111-111111111111"),
	}

	service := NewUserService(mockQuerier, mockTokenSvc)
	ctx := context.Background()

	tests := []struct {
		name       string
		input      types.SignupRequest
		setupMocks func()
		expectErr  error
	}{
		{
			name: "successful registration",
			input: types.SignupRequest{
				Name:     "Test User",
				Email:    "test@example.com",
				Password: "securepassword",
			},
			setupMocks: func() {
				mockQuerier.EXPECT().CheckEmailExists(ctx, "test@example.com").Return(false, nil)
				mockQuerier.EXPECT().CreateUser(ctx, gomock.Any()).Return(uuid.New(), nil)
			},
			expectErr: nil,
		},
		{
			name: "email already taken",
			input: types.SignupRequest{
				Name:     "Test User",
				Email:    "taken@example.com",
				Password: "anotherpass",
			},
			setupMocks: func() {
				mockQuerier.EXPECT().CheckEmailExists(ctx, "taken@example.com").Return(true, nil)
			},
			expectErr: ErrEmailTaken,
		},
		{
			name: "password too short",
			input: types.SignupRequest{
				Name:     "Short Pass",
				Email:    "shortpass@example.com",
				Password: "short",
			},
			setupMocks: func() {},
			expectErr:  ErrPasswordTooShort,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.setupMocks()
			_, err := service.RegisterUser(ctx, tc.input)
			if tc.expectErr != nil {
				require.ErrorIs(t, err, tc.expectErr)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestUserService_Login(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	name := "User Name"

	mockQuerier := mock.NewMockQuerier(ctrl)
	mockTokenSvc := &auth.MockTokenService{
		TokenToReturn: "dummy-token",
		ErrToReturn:   nil,
		ParsedUserID:  uuid.MustParse("11111111-1111-1111-1111-111111111111"),
	}

	service := NewUserService(mockQuerier, mockTokenSvc)
	ctx := context.Background()

	hash, _ := bcrypt.GenerateFromPassword([]byte("correctpassword"), bcrypt.DefaultCost)

	tests := []struct {
		name       string
		input      types.LoginRequest
		setupMocks func()
		expectErr  error
	}{
		{
			name: "successful login",
			input: types.LoginRequest{
				Email:    "user@example.com",
				Password: "correctpassword",
			},
			setupMocks: func() {
				mockQuerier.EXPECT().GetUserByEmail(ctx, "user@example.com").Return(sqlc.GetUserByEmailRow{
					Email:        "user@example.com",
					PasswordHash: string(hash),
					Name:         &name,
				}, nil)
			},
			expectErr: nil,
		},
		{
			name: "wrong password",
			input: types.LoginRequest{
				Email:    "user@example.com",
				Password: "wrongpassword",
			},
			setupMocks: func() {
				mockQuerier.EXPECT().GetUserByEmail(ctx, "user@example.com").Return(sqlc.GetUserByEmailRow{
					Email:        "user@example.com",
					PasswordHash: string(hash),
					Name:         &name,
				}, nil)
			},
			expectErr: ErrAuthenticationFailed,
		},
		{
			name: "email not found",
			input: types.LoginRequest{
				Email:    "nonexistent@example.com",
				Password: "anyPassword",
			},
			setupMocks: func() {
				mockQuerier.EXPECT().GetUserByEmail(ctx, "nonexistent@example.com").Return(sqlc.GetUserByEmailRow{}, ErrEmailNotFound)
			},
			expectErr: ErrEmailNotFound,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.setupMocks()
			_, err := service.Login(ctx, tc.input)
			if tc.expectErr != nil {
				require.ErrorIs(t, err, tc.expectErr)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
func TestUserService_GetUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockQuerier := mock.NewMockQuerier(ctrl)
	mockTokenSvc := &auth.MockTokenService{ParsedUserID: uuid.New()}
	service := NewUserService(mockQuerier, mockTokenSvc)
	name := "Test"

	tests := []struct {
		name       string
		ctx        context.Context
		setupMocks func(uuid.UUID)
		expectErr  error
	}{
		{
			name: "successful get",
			ctx:  context.WithValue(context.Background(), auth.UserIDKey, mockTokenSvc.ParsedUserID),
			setupMocks: func(id uuid.UUID) {
				mockQuerier.EXPECT().
					GetUserByID(gomock.Any(), id).
					Return(sqlc.GetUserByIDRow{Name: &name, Email: "test@example.com"}, nil)
			},
			expectErr: nil,
		},
		{
			name:       "invalid uuid",
			ctx:        context.WithValue(context.Background(), auth.UserIDKey, "not-a-uuid"),
			setupMocks: func(_ uuid.UUID) {}, // Not called
			expectErr:  appError.ErrInvalidUserId,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.setupMocks(mockTokenSvc.ParsedUserID)
			_, err := service.GetUser(tc.ctx)
			if tc.expectErr != nil {
				require.ErrorIs(t, err, tc.expectErr)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestUserService_DeleteUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockQuerier := mock.NewMockQuerier(ctrl)
	mockTokenSvc := &auth.MockTokenService{ParsedUserID: uuid.New()}
	service := NewUserService(mockQuerier, mockTokenSvc)

	tests := []struct {
		name       string
		ctx        context.Context
		setupMocks func(uuid.UUID)
		expectErr  error
	}{
		{
			name: "successful delete",
			ctx:  context.WithValue(context.Background(), auth.UserIDKey, mockTokenSvc.ParsedUserID),
			setupMocks: func(id uuid.UUID) {
				mockQuerier.EXPECT().
					DeleteUser(gomock.Any(), id).
					Return(nil)
			},
			expectErr: nil,
		},
		{
			name:       "bad uuid",
			ctx:        context.WithValue(context.Background(), auth.UserIDKey, "not-a-uuid"),
			setupMocks: func(_ uuid.UUID) {},
			expectErr:  appError.ErrInvalidUserId,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			tc.setupMocks(mockTokenSvc.ParsedUserID)
			err := service.DeleteUser(tc.ctx)
			if tc.expectErr != nil {
				require.ErrorIs(t, err, tc.expectErr)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestUserService_UpdateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockQuerier := mock.NewMockQuerier(ctrl)
	mockTokenSvc := &auth.MockTokenService{ParsedUserID: uuid.New()}
	service := NewUserService(mockQuerier, mockTokenSvc)

	name := "Updated"
	email := "new@example.com"

	validPassword := "validpass"
	tests := []struct {
		name       string
		req        types.UserUpdateRequest
		setupMocks func(uuid.UUID, *string)
		expectErr  error
	}{
		{
			name: "update with password",
			req: types.UserUpdateRequest{
				Name:     &name,
				Email:    &email,
				Password: &validPassword,
			},
			setupMocks: func(id uuid.UUID, hash *string) {
				mockQuerier.EXPECT().
					UpdateUser(gomock.Any(), gomock.Any()).
					Return(nil)

			},
			expectErr: nil,
		},
		{
			name: "bcrypt error",
			req: types.UserUpdateRequest{
				Password: func() *string {
					s := string(make([]byte, 1000)) // Too long for bcrypt
					return &s
				}(),
			},
			setupMocks: func(uuid.UUID, *string) {},
			expectErr:  ErrHashError,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.WithValue(context.Background(), auth.UserIDKey, mockTokenSvc.ParsedUserID)

			var hashPtr *string
			if tc.expectErr != ErrHashError && tc.req.Password != nil && *tc.req.Password != "" {
				h, err := bcrypt.GenerateFromPassword([]byte(*tc.req.Password), bcrypt.DefaultCost)
				require.NoError(t, err)
				s := string(h)
				hashPtr = &s
			}

			tc.setupMocks(mockTokenSvc.ParsedUserID, hashPtr)

			err := service.UpdateUser(ctx, tc.req)
			if tc.expectErr != nil {
				require.ErrorIs(t, err, tc.expectErr)
			} else {
				require.NoError(t, err)
			}
		})
	}

}
