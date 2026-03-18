-- name: GetProtocolsDevices :many
SELECT
    device_id,
    protocol_id,
    enabled,
    assigned_at,
    port,
    polling_rate,
    encryption,
    vault_reverence
FROM devices_protocols;

-- name: GetIPfromDeviceID :one
SELECT ip_address
FROM devices
WHERE id = $1;

-- name: GetIPandTypefromDeviceID :one
SELECT
    ip_address,
    type
FROM devices
WHERE id = $1;
