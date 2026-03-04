-- +goose Up
CREATE TABLE metrics (
	metric_name TEXT NOT NULL,
	value FLOAT8 NOT NULL,
	device_id UUID NOT NULL,
	collected_at TIMESTAMPTZ NOT NULL,
	CONSTRAINT fk_deviceid FOREIGN KEY (device_id) REFERENCES devices(id),
	PRIMARY KEY (collected_at, device_id, metric_name)
);

SELECT create_hypertable(
	'metrics',
	'collected_at',
	partitioning_column => 'device_id',
	chunk_time_interval => INTERVAL '7 days'
);

-- +goose Down
DROP TABLE metrics;
