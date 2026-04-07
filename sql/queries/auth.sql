-- name: GetUserByEmail :one
SELECT
  id,
  first_name,
  last_name,
  role,
  job_title,
  email,
  phone_number,
  password_hash,
  created_at,
  last_login,
  on_call
FROM users
WHERE email = $1
LIMIT 1;

-- name: CreateSession :exec
INSERT INTO sessions (
  id,
  user_id,
  created_at,
  refreshed_at,
  expires_at,
  token_hash
)
VALUES (
  gen_random_uuid(),
  $1,
  now(),
  now(),
  now() + INTERVAL '1 hour',
  $2
);

-- name: GetSessionByTokenHash :one
SELECT
  id,
  user_id,
  created_at,
  refreshed_at,
  expires_at
FROM sessions
WHERE token_hash = $1
LIMIT 1;

-- name: DeleteSession :exec
DELETE FROM sessions
WHERE token_hash = $1;

-- name: UpdateLastLogin :exec
UPDATE users
SET last_login = NOW()
WHERE id = $1;

-- name: CreateUser :exec
INSERT INTO users (
    id,
    first_name,
    last_name,
    role,
    job_title,
    email,
    phone_number,
    password_hash
)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8
);
