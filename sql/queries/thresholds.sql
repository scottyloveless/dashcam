-- name: GetActiveThreshold :one
SELECT * FROM thresholds
WHERE metric = $1
AND (device_id = $2 OR device_type = $3 OR (device_id IS NULL AND device_type IS NULL))
ORDER BY
  CASE
    WHEN device_id IS NOT NULL THEN 1
    WHEN device_type IS NOT NULL THEN 2
    ELSE 3
  END
LIMIT 1;
