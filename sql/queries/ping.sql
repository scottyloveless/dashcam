-- name: WritePing :exec
INSERT INTO metrics (metric_name, value, device_id, requested_at, received_at)
VALUES (
	$1,
	$2,
	$3,
	$4,
	$5
	);
