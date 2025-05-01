-- TRIGGER FUNCTION FOR updated_at
CREATE OR REPLACE FUNCTION update_modified_column()
RETURNS TRIGGER AS $$
BEGIN
   NEW.updated_at = NOW();
   RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- TRIGGERS
CREATE TRIGGER trg_update_forges_modtime
BEFORE UPDATE ON forges
FOR EACH ROW EXECUTE FUNCTION update_modified_column();

CREATE TRIGGER trg_update_sparks_modtime
BEFORE UPDATE ON sparks
FOR EACH ROW EXECUTE FUNCTION update_modified_column();

CREATE TRIGGER trg_update_forge_access_modtime
BEFORE UPDATE ON forge_access
FOR EACH ROW EXECUTE FUNCTION update_modified_column();

-- INDEXES
CREATE INDEX idx_forges_owner ON forges(owner_id);
CREATE INDEX idx_sparks_forge ON sparks(forge_id);
CREATE INDEX idx_forge_access_forge ON forge_access(forge_id);
CREATE INDEX idx_forge_access_user ON forge_access(user_id);
CREATE INDEX idx_api_keys_forge ON api_keys(forge_id);
CREATE INDEX idx_forge_access_active ON forge_access(forge_id, user_id) WHERE revoked_at IS NULL;
