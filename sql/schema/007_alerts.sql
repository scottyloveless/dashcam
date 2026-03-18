-- +goose Up
CREATE TYPE state_enum AS ENUM ('open', 'acknowledged', 'resolved', 'cleared');
CREATE TYPE severity_enum AS ENUM ('warning', 'critical');

CREATE TABLE alerts (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at timestamptz NOT NULL DEFAULT now(),
    last_occurrence timestamptz,
    ack_at timestamptz,
    resolved_at timestamptz,
    cleared_at timestamptz,
    device_id uuid NOT NULL,
    ack_by uuid,
    resolved_by uuid,
    alert_metric text NOT NULL,
    threshold_id uuid NOT NULL,
    severity severity_enum NOT NULL,
    state state_enum NOT NULL DEFAULT 'open',
    notes text,

    CONSTRAINT fk_thresholdid
    FOREIGN KEY (threshold_id)
    REFERENCES thresholds (id)
    ON DELETE CASCADE,

    CONSTRAINT fk_deviceid
    FOREIGN KEY (device_id)
    REFERENCES devices (id)
    ON DELETE CASCADE
);
-- +goose Down
DROP TABLE alerts;
DROP TYPE state_enum;
DROP TYPE severity_enum;
