// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: api_keys.sql

package db

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const deleteAPIKey = `-- name: DeleteAPIKey :exec
DELETE FROM api_keys WHERE id = $1
`

func (q *Queries) DeleteAPIKey(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.Exec(ctx, deleteAPIKey, id)
	return err
}

const getAPIKeyIDByHash = `-- name: GetAPIKeyIDByHash :one
SELECT id FROM api_keys WHERE key_hash = $1 AND is_active = TRUE
`

func (q *Queries) GetAPIKeyIDByHash(ctx context.Context, keyHash string) (uuid.UUID, error) {
	row := q.db.QueryRow(ctx, getAPIKeyIDByHash, keyHash)
	var id uuid.UUID
	err := row.Scan(&id)
	return id, err
}

const getAPIKeysByForgeID = `-- name: GetAPIKeysByForgeID :many
SELECT id, title, is_active, created_at, last_used_at FROM api_keys WHERE forge_id = $1 AND is_active = TRUE
`

type GetAPIKeysByForgeIDRow struct {
	ID         uuid.UUID
	Title      string
	IsActive   bool
	CreatedAt  time.Time
	LastUsedAt pgtype.Timestamptz
}

func (q *Queries) GetAPIKeysByForgeID(ctx context.Context, forgeID string) ([]GetAPIKeysByForgeIDRow, error) {
	rows, err := q.db.Query(ctx, getAPIKeysByForgeID, forgeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetAPIKeysByForgeIDRow
	for rows.Next() {
		var i GetAPIKeysByForgeIDRow
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.IsActive,
			&i.CreatedAt,
			&i.LastUsedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const insertAPIKey = `-- name: InsertAPIKey :exec
INSERT INTO api_keys (forge_id, title, key_hash)
VALUES ($1, $2, $3)
`

type InsertAPIKeyParams struct {
	ForgeID string
	Title   string
	KeyHash string
}

func (q *Queries) InsertAPIKey(ctx context.Context, arg InsertAPIKeyParams) error {
	_, err := q.db.Exec(ctx, insertAPIKey, arg.ForgeID, arg.Title, arg.KeyHash)
	return err
}

const toggleAPIKeyStatus = `-- name: ToggleAPIKeyStatus :exec
UPDATE api_keys SET is_active = NOT is_active WHERE id = $1
`

func (q *Queries) ToggleAPIKeyStatus(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.Exec(ctx, toggleAPIKeyStatus, id)
	return err
}
