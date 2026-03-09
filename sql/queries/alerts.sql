-- name: WriteAlert :exec
INSERT INTO alerts (
	last_occurrence,
	device_id,
	alert_metric,
	threshold_id,
	severity
	) VALUES (
	NOW(),
	$1,
	$2,
	$3,
	$4
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
