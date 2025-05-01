package services

import (
	"context"
	"testing"

	sqlc "stribog/internal/db/sqlc"
	"stribog/internal/db/sqlc/mock"
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
	service := NewUserService(mockQuerier)
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

	mockQuerier := mock.NewMockQuerier(ctrl)
	service := NewUserService(mockQuerier)
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
					Name:         types.ToPgText("User Name"),
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
					Name:         types.ToPgText("User Name"),
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
