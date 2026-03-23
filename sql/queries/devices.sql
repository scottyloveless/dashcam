-- name: GetDevices :many
SELECT
    id,
    created_at,
    nickname,
    hostname,
    ip_address,
    last_seen_at,
    enabled,
    polling_interval,
    serial,
    manufacturer,
    model,
    location,
    room,
    type,
    os,
    notes,
    tags
FROM devices;

-- name: GetOneDeviceInfo :one
SELECT
    id,
    created_at,
    nickname,
    hostname,
    ip_address,
    last_seen_at,
    enabled,
    polling_interval,
    serial,
    manufacturer,
    model,
    location,
    room,
    type,
    os,
    notes,
    tags
FROM devices
WHERE id = $1;

-- name: GetDistinctLocations :many
SELECT DISTINCT location FROM devices;

-- name: GetDevicesOneLocation :many
SELECT
    id,
    created_at,
    nickname,
    hostname,
    ip_address,
    last_seen_at,
    enabled,
    polling_interval,
    serial,
    manufacturer,
    model,
    location,
    room,
    type,
    os,
    notes,
    tags
FROM devices
WHERE location = $1;

-- name: CreateDevice :exec
INSERT INTO devices (
    id,
    nickname,
    hostname,
    ip_address,
    last_seen_at,
    polling_interval,
    serial,
    manufacturer,
    model,
    location,
    room,
    type,
    os,
    notes,
    tags
)
VALUES (
    uuid(),
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8,
    $9,
    $10,
    $11,
    $12,
    $13,
    $14
);
