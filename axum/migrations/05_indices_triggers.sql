-- Add migration script here
CREATE OR REPLACE FUNCTION update_modified_column()

-- TRIGGER CREATE
RETURNS TRIGGER AS $$
BEGIN
   NEW.updated_at = NOW();
   RETURN NEW;
END;
$$ LANGUAGE 'plpgsql';

-- USERS UPDATE
CREATE TRIGGER update_users_modtime
BEFORE UPDATE ON users
FOR EACH ROW EXECUTE FUNCTION update_modified_column();

-- FORGES UPDATE
CREATE TRIGGER update_forges_modtime
BEFORE UPDATE ON forges
FOR EACH ROW EXECUTE FUNCTION update_modified_column();

-- SPARKS UPDATE
CREATE TRIGGER update_sparks_modtime
BEFORE UPDATE ON sparks
FOR EACH ROW EXECUTE FUNCTION update_modified_column();

-- Create indexes for performance
CREATE INDEX idx_forges_owner ON forges(owner_id);
CREATE INDEX idx_api_keys_forge ON api_keys(forge_id);
CREATE INDEX idx_sparks_forge ON sparks(forge_id);