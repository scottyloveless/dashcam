-- +goose Up
ALTER TABLE alerts
ADD COLUMN display_message TEXT;

UPDATE alerts a
SET
    display_message = COALESCE(d.nickname, d.hostname, 'unknown-device')
    || ' — ' || a.alert_metric
    || ' ' || a.severity::TEXT
FROM devices d
WHERE a.device_id = d.id;

ALTER TABLE alerts ALTER COLUMN display_message SET NOT NULL;

-- +goose Down
ALTER TABLE alerts DROP COLUMN display_message;
