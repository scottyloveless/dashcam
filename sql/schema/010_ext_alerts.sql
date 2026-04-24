-- +goose Up
ALTER TABLE alerts
ADD COLUMN source TEXT NOT NULL DEFAULT 'internal',
ADD COLUMN source_ref TEXT,
ADD COLUMN external_device_name TEXT,
ADD COLUMN external_raw_json JSONB;

ALTER TABLE alerts
ALTER COLUMN device_id DROP NOT NULL,
ALTER COLUMN threshold_id DROP NOT NULL;

ALTER TABLE alerts
ADD CONSTRAINT chk_source
CHECK (source IN ('internal', 'ninja')),
ADD CONSTRAINT chk_device_reference CHECK (
    (source = 'internal' AND device_id IS NOT NULL AND external_device_name IS NULL)
    OR
    (source != 'internal' AND device_id IS NULL AND external_device_name IS NOT NULL)
),
ADD CONSTRAINT chk_threshold_reference CHECK (
    (source = 'internal' AND threshold_id IS NOT NULL)
    OR
    (source != 'internal' AND threshold_id IS NULL)
);

CREATE UNIQUE INDEX alerts_source_ref_unique
ON alerts (source, source_ref)
WHERE source_ref IS NOT NULL;

-- +goose Down
DROP INDEX IF EXISTS alerts_source_ref_unique;

ALTER TABLE alerts
DROP CONSTRAINT IF EXISTS chk_threshold_reference,
DROP CONSTRAINT IF EXISTS chk_device_reference,
DROP CONSTRAINT IF EXISTS chk_source;

ALTER TABLE alerts
ALTER COLUMN threshold_id SET NOT NULL,
ALTER COLUMN device_id SET NOT NULL;

ALTER TABLE alerts
DROP COLUMN external_device_name,
DROP COLUMN source_ref,
DROP COLUMN source,
DROP COLUMN external_raw_json;
