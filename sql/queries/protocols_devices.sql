-- name: GetProtocolsDevices :many
SELECT * FROM devices_protocols;

-- name: GetIPfromDeviceID :one
SELECT ip_address FROM devices WHERE id = $1;

-- name: GetIPandTypefromDeviceID :one
SELECT ip_address, type FROM devices WHERE id = $1;
