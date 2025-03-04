-- Add migration script here
CREATE TABLE users (
    id VARCHAR(21) PRIMARY KEY,
    name VARCHAR NOT NULL CHECK (LENGTH(name) >= 3),
    username VARCHAR NOT NULL,
    email VARCHAR NOT NULL,
    pass VARCHAR NOT NULL,
    role VARCHAR NOT NULL,
    verified BOOLEAN NOT NULL DEFAULT false,
    joined TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CONSTRAINT user_username_key UNIQUE (username),
    CONSTRAINT user_email_key UNIQUE (email),
    CONSTRAINT email_format CHECK (email ~* '^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$')
);