-- DROP INDEXES
DROP INDEX IF EXISTS idx_sparks_forge_slug;

DROP INDEX IF EXISTS idx_spark_tags_tag;

DROP INDEX IF EXISTS idx_sparks_slug;

DROP INDEX IF EXISTS idx_forge_access_active;

DROP INDEX IF EXISTS idx_api_keys_forge;

DROP INDEX IF EXISTS idx_forge_access_user;

DROP INDEX IF EXISTS idx_forge_access_forge;

DROP INDEX IF EXISTS idx_sparks_forge;

DROP INDEX IF EXISTS idx_forges_owner;

-- DROP TRIGGERS
DROP TRIGGER IF EXISTS trg_update_forge_access_modtime ON forge_access;

DROP TRIGGER IF EXISTS trg_update_sparks_modtime ON sparks;

DROP TRIGGER IF EXISTS trg_update_forges_modtime ON forges;

-- DROP TRIGGER FUNCTION
DROP FUNCTION IF EXISTS set_updated_at_timestamp;
