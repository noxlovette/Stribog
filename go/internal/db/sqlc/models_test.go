package db

import (
	"database/sql/driver"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccessRoleConstants(t *testing.T) {
	// Test case 6: Verify all constant AccessRole values
	assert.Equal(t, AccessRole("viewer"), AccessRoleViewer)
	assert.Equal(t, AccessRole("editor"), AccessRoleEditor)
	assert.Equal(t, AccessRole("admin"), AccessRoleAdmin)
}

func TestAccessRoleScan(t *testing.T) {
	testCases := []struct {
		name        string
		input       interface{}
		expected    AccessRole
		expectError bool
	}{
		{
			name:        "Scan from string",
			input:       "editor",
			expected:    AccessRoleEditor,
			expectError: false,
		},
		{
			name:        "Scan from []byte",
			input:       []byte("admin"),
			expected:    AccessRoleAdmin,
			expectError: false,
		},
		{
			name:        "Scan from invalid type",
			input:       123,
			expected:    "",
			expectError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var role AccessRole
			err := role.Scan(tc.input)

			if tc.expectError {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), "unsupported scan type")
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected, role)
			}
		})
	}
}

func TestNullAccessRoleScan(t *testing.T) {
	testCases := []struct {
		name        string
		input       interface{}
		expected    AccessRole
		valid       bool
		expectError bool
	}{
		{
			name:        "Scan from valid string",
			input:       "viewer",
			expected:    AccessRoleViewer,
			valid:       true,
			expectError: false,
		},
		{
			name:        "Scan from null",
			input:       nil,
			expected:    "",
			valid:       false,
			expectError: false,
		},
		{
			name:        "Scan from []byte",
			input:       []byte("admin"),
			expected:    AccessRoleAdmin,
			valid:       true,
			expectError: false,
		},
		{
			name:        "Scan from invalid type",
			input:       123,
			expected:    "",
			valid:       true,
			expectError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var nullRole NullAccessRole
			err := nullRole.Scan(tc.input)

			if tc.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected, nullRole.AccessRole)
				assert.Equal(t, tc.valid, nullRole.Valid)
			}
		})
	}
}

func TestNullAccessRoleValue(t *testing.T) {
	testCases := []struct {
		name          string
		nullRole      NullAccessRole
		expectedValue driver.Value
		expectError   bool
	}{
		{
			name: "Value from valid role",
			nullRole: NullAccessRole{
				AccessRole: AccessRoleEditor,
				Valid:      true,
			},
			expectedValue: "editor",
			expectError:   false,
		},
		{
			name: "Value from null role",
			nullRole: NullAccessRole{
				AccessRole: AccessRoleViewer, // Should be ignored since Valid is false
				Valid:      false,
			},
			expectedValue: nil,
			expectError:   false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			value, err := tc.nullRole.Value()

			if tc.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tc.expectedValue, value)
			}
		})
	}
}

