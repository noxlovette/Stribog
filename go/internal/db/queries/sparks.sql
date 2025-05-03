-- name: GetSparkAndCheckReadAccess :one
SELECT s.id, s.title, s.markdown FROM sparks s
JOIN forges f ON s.forge_id = f.id
WHERE s.id = $2 AND (
    s.owner_id = $1
    OR f.owner_id = $1
    OR EXISTS (
        SELECT 1 FROM forge_access fa
        WHERE fa.forge_id = s.forge_id AND fa.user_id = $1
    )
)
;

-- name: GetSparksAndCheckReadAccess :many
SELECT s.id, s.title, s.markdown FROM sparks s
JOIN forges f ON s.forge_id = f.id
WHERE
    s.owner_id = $1
    OR f.owner_id = $1
    OR EXISTS (
        SELECT 1 FROM forge_access fa
        WHERE fa.forge_id = s.forge_id AND fa.user_id = $1
        )
ORDER BY f.updated_at DESC;

-- name: InsertSpark :exec
INSERT INTO sparks (id, title, markdown, forge_id, owner_id)
VALUES ($1, $2, $3, $4, $5)
ON CONFLICT DO NOTHING;

-- name: UpdateSparkAndCheckWriteAccess :exec
UPDATE sparks s
SET
    title = COALESCE(sqlc.narg('title'), s.title),
    markdown = COALESCE(sqlc.narg('markdown'), s.markdown)
WHERE s.id = $2 AND (
    s.owner_id = $1
    OR EXISTS (
        SELECT 1 FROM forges f
        WHERE f.id = s.forge_id AND f.owner_id = $1
    )
    OR EXISTS (
        SELECT 1 FROM forge_access fa
        WHERE fa.forge_id = s.forge_id
        AND fa.user_id = $1
        AND fa.access_role IN ('admin', 'editor')
    )
)
;
-- name: DeleteSparkAndCheckAdminAccess :exec
DELETE FROM sparks s
WHERE s.id = $2 AND (
    s.owner_id = $1
    OR EXISTS (
        SELECT 1 FROM forges f
        WHERE f.id = s.forge_id AND f.owner_id = $1
    )
    OR EXISTS (
        SELECT 1 FROM forge_access fa
        WHERE fa.forge_id = s.forge_id
        AND fa.user_id = $1
        AND fa.access_role = 'admin'
    )
);

-- name: GetSparksByForgeIDAndCheckReadAccess :many
SELECT s.id, s.title, s.markdown FROM sparks s
JOIN forges f ON s.forge_id = f.id
WHERE
    s.forge_id = $2 AND (
    s.owner_id = $1
    OR f.owner_id = $1
    OR EXISTS (
        SELECT 1 FROM forge_access fa
        WHERE fa.forge_id = s.forge_id AND fa.user_id = $1
        )
        )
ORDER BY f.updated_at DESC;
