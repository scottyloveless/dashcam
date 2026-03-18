-- name: GetDevices :many
SELECT *
FROM devices;

-- name: GetOneDeviceInfo :one
SELECT * FROM devices WHERE id = $1;

-- name: GetDistinctLocations :many
SELECT DISTINCT location FROM devices;

-- name: GetDevicesOneLocation :many
SELECT * FROM devices WHERE location = $1;
