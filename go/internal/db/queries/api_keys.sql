-- name: GetAPIKeyIDByHash :one
SELECT id FROM api_keys WHERE key_hash = $1 AND is_active = TRUE;

-- name: InsertAPIKey :exec
INSERT INTO api_keys (forge_id, title, key_hash)
VALUES ($1, $2, $3);

-- name: GetAPIKeysByForgeID :many
SELECT id, title, is_active, created_at, last_used_at FROM api_keys WHERE forge_id = $1 AND is_active = TRUE;

-- name: DeleteAPIKey :exec
DELETE FROM api_keys WHERE id = $1;
-- name: ToggleAPIKeyStatus :exec
UPDATE api_keys SET is_active = NOT is_active WHERE id = $1;
