-- Add migration script here
CREATE TABLE api_keys (
    id VARCHAR(21) PRIMARY KEY,
    forge_id TEXT NOT NULL REFERENCES forges(id) ON DELETE CASCADE,
    title TEXT NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    last_used_at TIMESTAMPTZ
);