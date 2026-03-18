-- +goose Up
CREATE TABLE metrics (
    metric_name TEXT NOT NULL,
    value FLOAT8 NOT NULL,
    device_id UUID NOT NULL,
    requested_at TIMESTAMPTZ NOT NULL,
    received_at TIMESTAMPTZ NOT NULL,
    CONSTRAINT fk_deviceid FOREIGN KEY (device_id) REFERENCES devices (id),
    PRIMARY KEY (requested_at, device_id, metric_name)
);

SELECT create_hypertable(
    'metrics',
    'requested_at',
    partitioning_column => 'device_id',
    number_partitions => 4,
    chunk_time_interval => INTERVAL '7 days'
);

-- +goose Down
DROP TABLE metrics;
