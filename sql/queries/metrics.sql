-- name: GetPacketLossByDeviceID :one
SELECT *
FROM metrics
WHERE device_id = $1 AND name = 'packet_loss'
ORDER BY requested_at DESC
LIMIT 1;

-- name: GetRttAvgByDeviceID :one
SELECT *
FROM metrics
WHERE device_id = $1 AND name = 'rtt_avg'
ORDER BY requested_at DESC
LIMIT 1;
