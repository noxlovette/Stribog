-- name: GetForgeAndCheckReadAccess :one
SELECT f.id, f.title, f.description FROM forges f
WHERE f.id = $2 AND (
    f.owner_id = $1
    OR EXISTS (
        SELECT 1 FROM forge_access fa
        WHERE fa.forge_id = f.id AND fa.user_id = $1
    )
)
;

-- name: GetForgesAndCheckReadAccess :many
SELECT f.id, f.title, f.description FROM forges f
WHERE f.owner_id = $1
OR EXISTS (
    SELECT 1 FROM forge_access fa
    WHERE fa.forge_id = f.id AND fa.user_id = $1
)
ORDER BY f.updated_at DESC;


-- name: InsertForge :exec
INSERT INTO forges (id, title, description, owner_id)
VALUES ($1, $2, $3, $4)
ON CONFLICT DO NOTHING;

-- name: DeleteForge :exec
DELETE FROM forges WHERE id = $2 AND owner_id = $1;

-- name: UpdateForgeAndCheckWriteAccess :exec
UPDATE forges f
SET
    title = COALESCE(sqlc.narg('title'), f.title),
    description = COALESCE(sqlc.narg('description'), f.description)
WHERE f.id = $2 AND (
    f.owner_id = $1
    OR EXISTS (
        SELECT 1 FROM forge_access fa
        WHERE fa.forge_id = $2 AND fa.user_id = $1 AND fa.access_role = 'admin'
    )
);


-- name: UpsertForgeAccess :exec
INSERT INTO forge_access (forge_id, user_id, access_role, added_by)
VALUES ($1, $2, $3, $4)
ON CONFLICT (forge_id, user_id) DO UPDATE
SET access_role = $3, added_by = $4, updated_at = NOW();

-- name: DeleteForgeAccess :exec
DELETE FROM forge_access
WHERE user_id = $1 AND forge_id = $2 AND added_by = $3;

-- name: ListForgeAccess :many
SELECT fa.*, u.name as user_name, u.email as user_email
FROM forge_access fa
JOIN users u ON fa.user_id = u.id
WHERE fa.forge_id = $1
ORDER BY fa.updated_at DESC;
;

-- name: CheckWriteAccess :one
SELECT EXISTS (
    SELECT 1 FROM forges f
    WHERE (f.id = $2 AND f.owner_id = $1)
       OR EXISTS (
            SELECT 1 FROM forge_access
            WHERE forge_id = $2 AND user_id = $1 AND access_role = 'admin'
       )
) AS exists;

-- name: CheckReadAccess :one
SELECT EXISTS (
    SELECT 1 FROM forges f
    WHERE (f.id = $2 AND f.owner_id = $1)
       OR EXISTS (
            SELECT 1 FROM forge_access
            WHERE forge_id = $2 AND user_id = $1
       )
) AS exists;
