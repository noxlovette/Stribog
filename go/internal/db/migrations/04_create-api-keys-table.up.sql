CREATE TABLE api_keys (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
    forge_id VARCHAR(21) NOT NULL REFERENCES forges (id) ON DELETE CASCADE,
    title TEXT NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW (),
    last_used_at TIMESTAMPTZ,
    key_hash TEXT NOT NULL
);
