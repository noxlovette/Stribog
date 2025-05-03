-- name: GetSparksByForgeIDPublic :many
SELECT
  s.id, s.title, s.markdown, s.slug,
  COALESCE(ARRAY_AGG(st.tag) FILTER (WHERE st.tag IS NOT NULL), '{}')::TEXT[] AS tags
FROM sparks s
JOIN forges f ON s.forge_id = f.id
LEFT JOIN spark_tags st ON st.spark_id = s.id
WHERE s.forge_id = $1
GROUP BY s.id, s.title, s.markdown, s.slug
ORDER BY s.updated_at DESC;
