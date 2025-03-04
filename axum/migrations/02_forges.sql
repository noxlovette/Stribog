-- Add migration script here
CREATE TABLE forges (
    id VARCHAR(21) PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT,
    owner_id TEXT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Create a new table for forge access permissions
CREATE TABLE forge_access (
    id VARCHAR(21) PRIMARY KEY,
    forge_id VARCHAR(21) NOT NULL REFERENCES forges(id) ON DELETE CASCADE,
    user_id VARCHAR(21) NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    access_role VARCHAR(50) NOT NULL DEFAULT 'viewer', -- e.g., 'viewer', 'editor', 'admin'
    added_by VARCHAR(21) NOT NULL REFERENCES users(id),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CONSTRAINT unique_forge_user UNIQUE (forge_id, user_id)
);

