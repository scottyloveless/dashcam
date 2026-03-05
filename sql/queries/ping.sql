-- name: WritePing :exec
INSERT INTO metrics (metric_name, value, device_id, requested_at, received_at)
VALUES (
	$1,
	$2,
	'ade17d9a-3081-4ae4-8ba5-f8253979bfaf',
	$3,
	$4
	);
