package config

import (
	"testing"
)

func TestEnvWithBothFields(t *testing.T) {
	// Test case 1: Valid configuration with both fields populated
	env := Env{
		JWTKey:      "test-jwt-key",
		DatabaseDSN: "postgresql://user:password@localhost:5432/dbname",
	}

	if env.JWTKey != "test-jwt-key" {
		t.Errorf("Expected JWTKey to be 'test-jwt-key', got '%s'", env.JWTKey)
	}

	if env.DatabaseDSN != "postgresql://user:password@localhost:5432/dbname" {
		t.Errorf("Expected DatabaseDSN to be 'postgresql://user:password@localhost:5432/dbname', got '%s'", env.DatabaseDSN)
	}
}

func TestEmptyEnv(t *testing.T) {
	// Test case 2: Empty configuration validation
	env := Env{}

	if env.JWTKey != "" {
		t.Errorf("Expected JWTKey to be empty, got '%s'", env.JWTKey)
	}

	if env.DatabaseDSN != "" {
		t.Errorf("Expected DatabaseDSN to be empty, got '%s'", env.DatabaseDSN)
	}
}

func TestEnvWithOnlyJWTKey(t *testing.T) {
	// Test case 3: Configuration with only JWT key
	env := Env{
		JWTKey: "test-jwt-key",
	}

	if env.JWTKey != "test-jwt-key" {
		t.Errorf("Expected JWTKey to be 'test-jwt-key', got '%s'", env.JWTKey)
	}

	if env.DatabaseDSN != "" {
		t.Errorf("Expected DatabaseDSN to be empty, got '%s'", env.DatabaseDSN)
	}
}

func TestEnvWithOnlyDatabaseDSN(t *testing.T) {
	// Test case 4: Configuration with only Database DSN
	env := Env{
		DatabaseDSN: "postgresql://user:password@localhost:5432/dbname",
	}

	if env.JWTKey != "" {
		t.Errorf("Expected JWTKey to be empty, got '%s'", env.JWTKey)
	}

	if env.DatabaseDSN != "postgresql://user:password@localhost:5432/dbname" {
		t.Errorf("Expected DatabaseDSN to be 'postgresql://user:password@localhost:5432/dbname', got '%s'", env.DatabaseDSN)
	}
}

