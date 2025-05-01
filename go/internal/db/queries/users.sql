-- name: GetUserByID :one
SELECT * FROM users WHERE id = $1;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;

-- name: ListUsers :many
SELECT * FROM users ORDER BY created_at DESC;

-- name: CreateUser :exec
INSERT INTO users (id, email, password_hash, name, is_active)
VALUES ($1, $2, $3, $4, $5);
