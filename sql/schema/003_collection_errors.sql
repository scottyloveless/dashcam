-- +goose Up
CREATE TYPE error_type AS ENUM ('dns', 'unreachable', 'unauthorized', 'collector', 'saas_outage', 'isp');

CREATE TABLE collection_errors (
	device_id UUID NOT NULL,
	requested_at TIMESTAMPTZ NOT NULL,
	metric_name TEXT NOT NULL,
	type error_type NOT NULL,
	error_message TEXT NOT NULL,
	retries INTEGER NOT NULL,
	CONSTRAINT fk_deviceid FOREIGN KEY (device_id) REFERENCES devices(id),
	PRIMARY KEY (requested_at, device_id, metric_name)
);

SELECT create_hypertable(
	'collection_errors',
	'requested_at',
	partitioning_column => 'device_id',
	number_partitions => 4,
	chunk_time_interval => INTERVAL '7 days'
);

-- +goose Down
DROP TABLE collection_errors;
DROP TYPE error_type;
