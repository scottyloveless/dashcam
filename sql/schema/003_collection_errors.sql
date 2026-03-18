-- +goose Up
CREATE TYPE error_type
AS ENUM (
    'dns',
    'unreachable',
    'unauthorized',
    'collector',
    'saas_outage',
    'isp'
);

CREATE TABLE collection_errors (
    device_id uuid NOT NULL,
    requested_at timestamptz NOT NULL,
    metric_name text NOT NULL,
    type error_type NOT NULL,
    error_message text NOT NULL,
    retries integer NOT NULL,
    raw_response jsonb,
    CONSTRAINT fk_deviceid FOREIGN KEY (device_id) REFERENCES devices (id),
    PRIMARY KEY (requested_at, device_id, metric_name)
);

SELECT create_hypertable(
    'collection_errors',
    'requested_at',
    partitioning_column => 'device_id',
    number_partitions => 4,
    chunk_time_interval => interval '7 days'
);

-- +goose Down
DROP TABLE collection_errors;
DROP TYPE error_type;
