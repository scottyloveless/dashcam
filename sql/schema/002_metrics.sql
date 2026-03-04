-- +goose Up
CREATE TABLE metrics (
	name TEXT NOT NULL,
	value FLOAT8 NOT NULL,
	device_id UUID,
	timestamp TIMESTAMPTZ NOT NULL,
	answer_received TIMESTAMPTZ,
	CONSTRAINT fk_deviceid FOREIGN KEY (device_id) REFERENCES devices(id)
);

-- +goose Down
DROP TABLE metrics;
