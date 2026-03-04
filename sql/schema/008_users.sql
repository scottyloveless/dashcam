-- +goose Up
CREATE TYPE roles_enum AS ENUM ('administrator', 'editor', 'read-only');

CREATE TABLE users (
    id UUID PRIMARY KEY,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    role roles_enum NOT NULL DEFAULT 'read-only',
    job_title TEXT,
    email VARCHAR(100) NOT NULL UNIQUE,
    phone_number VARCHAR(15) NOT NULL, --for SMS alerts in future
    password_hash TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    last_login TIMESTAMPTZ,
    on_call BOOLEAN NOT NULL DEFAULT FALSE
);

-- +goose Down
DROP TABLE users;
DROP TYPE roles_enum;