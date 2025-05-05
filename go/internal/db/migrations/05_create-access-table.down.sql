DROP TABLE IF EXISTS forge_access;

DROP TYPE IF EXISTS access_role;

ALTER TABLE spark_tags
DROP CONSTRAINT spark_tags_spark_id_fkey;
