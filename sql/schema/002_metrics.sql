-- +goose Up
CREATE TABLE metrics (
	type TEXT NOT NULL,
	value FLOAT4 NOT NULL,
	device_id UUID,
	CONSTRAINT fk_deviceid FOREIGN KEY (device_id) REFERENCES devices(id)
);

-- +goose Down
DROP TABLE metrics;
