-- +goose Up
CREATE TABLE sessions (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    refreshed_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    expires_at TIMESTAMPTZ NOT NULL,
    token_hash TEXT NOT NULL UNIQUE,
    CONSTRAINT fk_userid
    FOREIGN KEY (user_id)
    REFERENCES users (id)
    ON DELETE RESTRICT
);

-- +goose Down
DROP TABLE sessions;

