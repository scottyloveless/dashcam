-- +goose Up
CREATE TABLE devices (
	id UUID PRIMARY KEY, --UUIDv7 to get time created as well
	created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	nickname TEXT NOT NULL UNIQUE,
	hostname TEXT UNIQUE,
	ip_address INET NOT NULL,
	last_seen_at TIMESTAMPTZ,
	enabled BOOLEAN NOT NULL DEFAULT TRUE,
	polling_interval INTERVAL,
	serial TEXT,
	manufacturer TEXT,
	model TEXT,
	location TEXT NOT NULL,
	room TEXT,
	type TEXT NOT NULL,
	os TEXT,
	notes TEXT,
	tags JSONB
);

-- +goose Down
DROP TABLE devices;
