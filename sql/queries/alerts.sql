-- name: WriteAlert :exec
INSERT INTO alerts (
    last_occurrence,
    device_id,
    alert_metric,
    threshold_id,
    severity,
    display_message
) VALUES (
    NOW(),
    $1,
    $2,
    $3,
    $4::severity_enum,
    (
        SELECT
            COALESCE(nickname, hostname, 'unknown-device')
            || ' — ' || $2::text || ' ' || ($4::severity_enum)::text
        FROM devices
        WHERE id = $1
    )
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
    display_message,
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
    alerts.id,
    alerts.display_message
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

-- name: ListOpenExternalAlerts :many
SELECT
    id,
    source,
    source_ref,
    external_device_name,
    alert_metric,
    severity,
    state,
    display_message,
    created_at,
    last_occurrence,
    ack_at
FROM alerts
WHERE
    source != 'internal'
    AND state IN ('open', 'acknowledged')
ORDER BY
    CASE severity WHEN 'critical' THEN 0 ELSE 1 END,
    CASE state WHEN 'open' THEN 0 ELSE 1 END,
    last_occurrence DESC NULLS LAST
LIMIT 50;
