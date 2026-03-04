-- +goose Up
CREATE TABLE devices_protocols (
	device_id UUID NOT NULL,
	protocol_id UUID NOT NULL,
	enabled BOOLEAN NOT NULL DEFAULT TRUE,
	assigned_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	port INTEGER NOT NULL,
	polling_rate INTERVAL NOT NULL,
	encryption BOOLEAN,
	vault_reference TEXT,
	PRIMARY KEY (device_id, protocol_id),
	FOREIGN KEY (device_id) REFERENCES devices(id) ON DELETE CASCADE,
	FOREIGN KEY (protocol_id) REFERENCES protocols(id) ON DELETE CASCADE
);
-- +goose Down
DROP TABLE devices_protocols;
