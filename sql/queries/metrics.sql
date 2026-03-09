-- name: GetPacketLossByDeviceID :one
SELECT *
FROM metrics
WHERE device_id = $1 AND metric_name = 'packet_loss'
ORDER BY requested_at DESC
LIMIT 1;

-- name: GetRttAvgByDeviceID :one
SELECT *
FROM metrics
WHERE device_id = $1 AND metric_name = 'rtt_avg'
ORDER BY requested_at DESC
LIMIT 1;

-- name: GetAllMetricsForOneDeviceByID :many
SELECT *
FROM metrics
WHERE device_id = $1
AND requested_at > NOW() - INTERVAL '12 hours'
ORDER BY requested_at DESC;
