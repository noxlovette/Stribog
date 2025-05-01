CREATE TABLE sparks (
    id VARCHAR(21) PRIMARY KEY, -- NanoID from app
    forge_id VARCHAR(21) NOT NULL REFERENCES forges (id) ON DELETE CASCADE,
    title TEXT NOT NULL,
    markdown TEXT NOT NULL,
    owner_id UUID NOT NULL REFERENCES users (id) ON DELETE SET NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW (),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW ()
);
