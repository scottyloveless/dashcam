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
SELECT * FROM alerts
WHERE device_id = $1
AND alert_metric = $2
AND state IN ('open', 'acknowledged');

-- name: UpdateAlertLastOccurrence :exec
UPDATE alerts
SET last_occurrence = NOW(), severity = $1
WHERE id = $2;

-- name: GetAlerts :many
SELECT devices.nickname, alerts.alert_metric, alerts.severity, alerts.created_at, alerts.last_occurrence
FROM alerts
INNER JOIN devices ON alerts.device_id = devices.id
WHERE state IN ('open', 'acknowledged')
ORDER BY 
	CASE severity
		WHEN 'critical' THEN 1
		WHEN 'warning' THEN 2
	END,
	GREATEST(alerts.created_at, alerts.last_occurrence) DESC;
