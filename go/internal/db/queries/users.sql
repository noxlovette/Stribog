-- name: GetUserByID :one
SELECT email, name FROM users WHERE id = $1;

-- name: GetUserByEmail :one
SELECT id, email, name, password_hash FROM users WHERE email = $1;

-- name: ListUsers :many
SELECT email, name FROM users ORDER BY created_at DESC
LIMIT $1 OFFSET $2;

-- name: CheckEmailExists :one
SELECT EXISTS (SELECT 1 FROM users WHERE email = $1) AS exists;

-- name: CreateUser :one
INSERT INTO users (email, password_hash, name)
VALUES ($1, $2, $3)
RETURNING id;
