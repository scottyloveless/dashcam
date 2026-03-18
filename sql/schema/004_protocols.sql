-- +goose Up
CREATE TABLE protocols (
    id UUID NOT NULL PRIMARY KEY,
    name TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- +goose Down
DROP TABLE protocols;
