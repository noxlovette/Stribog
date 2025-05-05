-- name: GetUserByID :one
SELECT email, name FROM users WHERE id = $1;

-- name: GetUserByEmail :one
SELECT id, email, name, password_hash FROM users WHERE email = $1;

-- name: CheckEmailExists :one
SELECT EXISTS (SELECT 1 FROM users WHERE email = $1) AS exists;

-- name: CreateUser :one
INSERT INTO users (email, password_hash, name)
VALUES ($1, $2, $3)
RETURNING id;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;

-- name: UpdateUser :exec
UPDATE users
SET
    name = COALESCE(sqlc.narg('name'), name),
    email = COALESCE(sqlc.narg('email'), email),
    password_hash = COALESCE(sqlc.narg('password_hash'), password_hash)
WHERE id = sqlc.arg('id');
