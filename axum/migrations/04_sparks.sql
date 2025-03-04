-- Add migration script here
CREATE TABLE sparks (
    id VARCHAR(21) PRIMARY KEY,
    forge_id VARCHAR(21) NOT NULL REFERENCES forges(id) ON DELETE CASCADE,
    title TEXT NOT NULL,
    markdown TEXT NOT NULL,
    owner_id TEXT NOT NULL REFERENCES users(id),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
