-- name: GetDevices :many
SELECT * FROM devices;

-- name: GetOneDeviceInfo :one
SELECT * FROM devices WHERE id = $1;
