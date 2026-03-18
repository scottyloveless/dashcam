-- +goose Up
CREATE TYPE roles_enum AS ENUM ('administrator', 'editor', 'read-only');

CREATE TABLE users (
    id uuid PRIMARY KEY,
    first_name text NOT NULL,
    last_name text NOT NULL,
    role roles_enum NOT NULL DEFAULT 'read-only',
    job_title text,
    email varchar(100) NOT NULL UNIQUE,
    phone_number varchar(15) NOT NULL, --for SMS alerts in future
    password_hash text NOT NULL,
    created_at timestamptz NOT NULL DEFAULT NOW(),
    last_login timestamptz,
    on_call boolean NOT NULL DEFAULT FALSE
);

-- +goose Down
DROP TABLE users;
DROP TYPE roles_enum;
