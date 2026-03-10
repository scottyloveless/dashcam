-- +goose Up
CREATE TYPE state_enum AS ENUM ('open', 'acknowledged', 'resolved', 'cleared');
CREATE TYPE severity_enum AS ENUM ('warning', 'critical');

CREATE TABLE alerts (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    last_occurrence TIMESTAMPTZ,
    ack_at TIMESTAMPTZ,
    resolved_at TIMESTAMPTZ,
    cleared_at TIMESTAMPTZ,
    device_id UUID NOT NULL,
    ack_by UUID,
    resolved_by UUID,
    alert_metric TEXT NOT NULL,
    threshold_id UUID NOT NULL,
    severity severity_enum NOT NULL,
    state state_enum NOT NULL DEFAULT 'open',
    notes TEXT,
    CONSTRAINT fk_thresholdid FOREIGN KEY (threshold_id) REFERENCES thresholds(id) ON DELETE CASCADE,
    CONSTRAINT fk_deviceid FOREIGN KEY (device_id) REFERENCES devices(id) ON DELETE CASCADE
);
-- +goose Down
DROP TABLE alerts;
DROP TYPE state_enum;
DROP TYPE severity_enum;
