-- +goose Up
CREATE TYPE direction_enum AS ENUM ('above', 'below', 'both');

CREATE TABLE thresholds (
    id uuid PRIMARY KEY,
    device_id uuid,
    device_type text,
    metric text NOT NULL,
    warning_value float8 NOT NULL,
    critical_value float8 NOT NULL,
    direction direction_enum NOT NULL,
    created_at timestamptz NOT NULL DEFAULT NOW(),
    modified_at timestamptz NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_deviceid FOREIGN
    KEY (device_id)
    REFERENCES devices (id)
    ON DELETE CASCADE
);

CREATE UNIQUE INDEX metric_deviceid_idx ON thresholds (metric, device_id)
WHERE device_id IS NOT NULL;
CREATE UNIQUE INDEX metric_devicetype_idx ON thresholds (metric, device_type)
WHERE device_type IS NOT NULL;
CREATE UNIQUE INDEX metric_idx ON thresholds (metric)
WHERE device_id IS NULL AND device_type IS NULL;

-- +goose Down
DROP TABLE thresholds;
DROP TYPE direction_enum;
