-- ACCESS ROLE ENUM
CREATE TYPE access_role AS ENUM ('viewer', 'editor', 'admin');

-- FORGE ACCESS TABLE
CREATE TABLE forge_access (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
    forge_id VARCHAR(21) NOT NULL REFERENCES forges (id) ON DELETE CASCADE,
    user_id UUID NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    access_role access_role NOT NULL DEFAULT 'viewer',
    added_by UUID NOT NULL REFERENCES users (id),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW (),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW (),
    revoked_at TIMESTAMPTZ,
    CONSTRAINT unique_forge_user UNIQUE (forge_id, user_id)
);
