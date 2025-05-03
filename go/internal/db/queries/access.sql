
-- name: CheckAdminAccess :one
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

-- name: CheckWriteAccess :one
SELECT EXISTS (
    SELECT 1 FROM forges f
    WHERE (f.id = $2 AND f.owner_id = $1)
       OR EXISTS (
            SELECT 1 FROM forge_access
            WHERE forge_id = $2 AND user_id = $1 AND access_role IN ('admin', 'editor')
       )
) AS exists;
