-- name: WriteAlert :exec
INSERT INTO alerts (
    id,
    last_occurrence,
    device_id,
    alert_metric,
    threshold_id,
    severity
) VALUES (
    $1,
    NOW(),
    $2,
    $3,
    $4,
    $5
);

-- name: CheckAlert :one
SELECT
    id,
    created_at,
    last_occurrence,
    ack_at,
    resolved_at,
    cleared_at,
    device_id,
    ack_by,
    resolved_by,
    alert_metric,
    threshold_id,
    severity,
    state,
    notes
FROM alerts
WHERE
    device_id = $1
    AND alert_metric = $2
    AND state IN ('open', 'acknowledged');

-- name: UpdateAlertLastOccurrence :exec
UPDATE alerts
SET
    last_occurrence = NOW(),
    severity = $1
WHERE id = $2;

-- name: GetAlerts :many
SELECT
    devices.id AS device_id,
    devices.nickname,
    alerts.alert_metric,
    alerts.severity,
    alerts.created_at,
    alerts.last_occurrence,
    alerts.id
FROM alerts
INNER JOIN devices
    ON alerts.device_id = devices.id
WHERE alerts.state IN ('open', 'acknowledged')
ORDER BY
    CASE alerts.severity
        WHEN 'critical' THEN 1
        WHEN 'warning' THEN 2
    END,
    GREATEST(alerts.created_at, alerts.last_occurrence) DESC;

-- name: ClearAlert :exec
UPDATE alerts
SET
    state = 'cleared',
    cleared_at = NOW()
WHERE id = $1;
