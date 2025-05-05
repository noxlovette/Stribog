-- FUNCTIONS
CREATE OR REPLACE FUNCTION set_updated_at_timestamp()
RETURNS TRIGGER AS $$
BEGIN
   NEW.updated_at = NOW();
   RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- TRIGGERS
DROP TRIGGER IF EXISTS trg_update_forges_modtime ON forges;
CREATE TRIGGER trg_update_forges_modtime
BEFORE UPDATE ON forges
FOR EACH ROW EXECUTE FUNCTION set_updated_at_timestamp();

DROP TRIGGER IF EXISTS trg_update_sparks_modtime ON sparks;
CREATE TRIGGER trg_update_sparks_modtime
BEFORE UPDATE ON sparks
FOR EACH ROW EXECUTE FUNCTION set_updated_at_timestamp();

DROP TRIGGER IF EXISTS trg_update_forge_access_modtime ON forge_access;
CREATE TRIGGER trg_update_forge_access_modtime
BEFORE UPDATE ON forge_access
FOR EACH ROW EXECUTE FUNCTION set_updated_at_timestamp();


-- INDEXES
CREATE INDEX idx_forges_owner ON forges(owner_id);
CREATE INDEX idx_sparks_forge ON sparks(forge_id);
CREATE INDEX idx_forge_access_forge ON forge_access(forge_id);
CREATE INDEX idx_forge_access_user ON forge_access(user_id);
CREATE INDEX idx_api_keys_forge ON api_keys(forge_id);
CREATE INDEX idx_forge_access_active ON forge_access(forge_id, user_id) WHERE revoked_at IS NULL;
CREATE INDEX idx_sparks_slug ON sparks(slug);
CREATE INDEX idx_spark_tags_tag ON spark_tags(tag);
CREATE UNIQUE INDEX idx_sparks_forge_slug ON sparks(forge_id, slug);
