CREATE TABLE sparks (
    id VARCHAR(21) PRIMARY KEY, -- NanoID from app
    forge_id VARCHAR(21) NOT NULL REFERENCES forges (id) ON DELETE CASCADE,
    title TEXT NOT NULL DEFAULT 'New Spark',
    markdown TEXT NOT NULL DEFAULT 'markdown',
    owner_id UUID NOT NULL REFERENCES users (id) ON DELETE RESTRICT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW (),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW (),
    slug TEXT NOT NULL DEFAULT '',
    tags TEXT[] DEFAULT ARRAY[]::TEXT[]
);
CREATE TABLE spark_tags (
    spark_id VARCHAR(21) NOT NULL REFERENCES sparks(id) ON DELETE CASCADE,
    tag TEXT NOT NULL,
    PRIMARY KEY (spark_id, tag)
);
