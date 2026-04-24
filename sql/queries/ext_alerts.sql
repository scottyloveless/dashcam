-- name: UpsertExternalAlert :exec
INSERT INTO alerts (
    created_at,
    last_occurrence,
    source,
    source_ref,
    external_device_name,
    alert_metric,
    severity,
    display_message,
    external_raw_json
)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8,
    $9
)
ON CONFLICT (source, source_ref) WHERE source_ref IS NOT NULL
DO UPDATE SET
    severity = excluded.severity,
    last_occurrence = excluded.last_occurrence,
    external_device_name = excluded.external_device_name,
    alert_metric = excluded.alert_metric,
    external_raw_json = excluded.external_raw_json,
    display_message = excluded.display_message;
