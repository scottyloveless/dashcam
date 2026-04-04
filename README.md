# dashcam - IT Infrastructure Monitoring in Go

## How it works

### Collectors

- Modular
- Devices can have any number of collectors defined
- Custom collectors can be easily made
- Writes them to a TimeScaleDB metrics table in Postgres.
- Each device is collected in it's own goroutine.

### Thresholds

- Can be defined per device, device category, or global defaults.
- Defined when above, below, or both above and below your defined number.

### Watchdog

- Runs every 5 seconds and iterates over the alerts table
- checks last five metrics to see if the alert has been resolved.
- If the last five metrics look good, the alert will auto clear.

## Dependencies

1. go 1.26.1
2. PosgreSQL 17
3. TimeScaleDB (PostgreSQL plugin)
4. goose - for database migrations

## Installation

### 1. Clone this repo

```bash
git clone https://github.com/scottyloveless/dashcam
cd dashcam
```

### 2. Add PostgreSQL DSN to .env

Create a new database in your PosgreSQL instance and add the DSN to your .env file as `DATABASE_URL`.

Example:

```
DATABASE_URL="postgres://user:password@localhost:5432/dbname
```

If you run into an error about ssl, add `?sslmode=disable` to the end of your DSN, like this:

```
DATABASE_URL="postgres://user:password@localhost:5432/dbname?sslmode=disable"
```

### 3. Run collector service to begin gathering metrics

```bash
go run ./cmd/collector/
```

### 4. Run webserver service to view dashboard in separate terminal session

```bash
   go run ./cmd/webserver/
```

### 5. View dashboard in browser

`https://localhost:4000` (or your custom port defined in cmd/webserver/main.go)
