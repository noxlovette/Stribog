-- name: GetSparkAndCheckReadAccess :one
SELECT
  s.id, s.title, s.markdown, s.slug, s.updated_at,
  COALESCE(ARRAY_AGG(st.tag) FILTER (WHERE st.tag IS NOT NULL), '{}')::TEXT[] AS tags
FROM sparks s
JOIN forges f ON s.forge_id = f.id
LEFT JOIN spark_tags st ON st.spark_id = s.id
WHERE s.id = $2 AND (
  s.owner_id = $1
  OR f.owner_id = $1
  OR EXISTS (
    SELECT 1 FROM forge_access fa
    WHERE fa.forge_id = s.forge_id AND fa.user_id = $1
  )
)
GROUP BY s.id, s.title, s.markdown, s.slug;


-- name: GetSparksByForgeIDAndCheckReadAccess :many
SELECT
  s.id, s.title, s.markdown, s.slug, s.updated_at,
  COALESCE(ARRAY_AGG(st.tag) FILTER (WHERE st.tag IS NOT NULL), '{}')::TEXT[] AS tags
FROM sparks s
JOIN forges f ON s.forge_id = f.id
LEFT JOIN spark_tags st ON st.spark_id = s.id
WHERE s.forge_id = $2 AND (
  s.owner_id = $1
  OR f.owner_id = $1
  OR EXISTS (
    SELECT 1 FROM forge_access fa
    WHERE fa.forge_id = s.forge_id AND fa.user_id = $1
  )
)
GROUP BY s.id, s.title, s.markdown, s.slug
ORDER BY s.updated_at DESC;


-- name: InsertSpark :exec
INSERT INTO sparks (id, forge_id, owner_id, slug)
VALUES ($1, $2, $3, $4);

-- name: UpdateSparkAndCheckWriteAccess :exec
UPDATE sparks s
SET
    title = COALESCE(sqlc.narg('title'), s.title),
    markdown = COALESCE(sqlc.narg('markdown'), s.markdown),
    slug = COALESCE(sqlc.narg('slug'), s.slug)
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
);


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

-- name: GetTagsForSpark :many
SELECT tag FROM spark_tags WHERE spark_id = $1;

-- name: DeleteSparkTags :exec
DELETE FROM spark_tags WHERE spark_id = $1;

-- name: InsertSparkTag :exec
INSERT INTO spark_tags (spark_id, tag) VALUES ($1, $2);
