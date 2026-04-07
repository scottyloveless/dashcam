package main

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"math/rand/v2"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/scottyloveless/dashcam/internal/database"
)

// const version = "1.0.0"

type config struct {
	env string
	db  struct {
		dsn          string
		maxOpenConns int
		maxIdleConns int
		maxIdleTime  time.Duration
	}
}

type application struct {
	config    config
	logger    *slog.Logger
	pool      *pgxpool.Pool
	queries   *database.Queries
	protocols []Trigger
}

func main() {
	var cfg config

	flag.StringVar(&cfg.env, "env", "development", "development|staging|production")

	flag.IntVar(&cfg.db.maxOpenConns, "db-max-open-conns", 25, "Postgres max open connections")
	flag.IntVar(&cfg.db.maxIdleConns, "db-max-idle-conns", 25, "Postgres max idle connections")
	flag.DurationVar(&cfg.db.maxIdleTime, "db-max-idle-time", 15*time.Minute, "Postgres max idle time")

	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	err := godotenv.Load(".env")
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	cfg.db.dsn = os.Getenv("DATABASE_URL")

	pool, err := openDB(cfg)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	queries := database.New(pool)

	app := application{
		config:    cfg,
		logger:    logger,
		pool:      pool,
		queries:   queries,
		protocols: []Trigger{},
	}

	fmt.Println("starting collector cycle...")
	app.protocols, err = app.triggerNetworkHelper()
	if err != nil {
		logger.Error(err.Error())
		time.Sleep(10 * time.Second)
		return
	}

	if len(app.protocols) <= 0 {
		time.Sleep(10 * time.Second)
		return
	}

	for _, protocol := range app.protocols {
		if !protocol.Trigger.Enabled {
			continue
		}

		go func(p Trigger) {
			pollingRate := p.Trigger.PollingRate.Microseconds

			ticker := time.NewTicker(time.Duration(pollingRate) * time.Microsecond)
			defer ticker.Stop()

			for range ticker.C {
				time.Sleep(time.Duration(rand.IntN(20)) * time.Millisecond)

				if err = app.collectPing(p); err != nil {
					app.logger.Error(err.Error())
					continue
				}
			}
		}(protocol)
	}

	watchdogTicker := time.NewTicker(5 * time.Second)
	defer watchdogTicker.Stop()

	// watchDogCycleCount := 1
	ctx := context.Background()

	go func() {
		for range watchdogTicker.C {
			// app.logger.Info("watchdog cycle: " + strconv.Itoa(watchDogCycleCount))
			alerts, err := app.queries.GetAlerts(ctx)
			if err != nil {
				app.logger.Error(err.Error())
				continue
			}
			for _, alert := range alerts {
				app.logger.Info(alert.Nickname)
				app.watchdog(alert, ctx)
			}

			// watchDogCycleCount++
		}
	}()

	select {}
}

func openDB(cfg config) (*pgxpool.Pool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool, err := pgxpool.New(ctx, cfg.db.dsn)
	if err != nil {
		return nil, err
	}

	err = pool.Ping(ctx)
	if err != nil {
		pool.Close()
		return nil, err
	}

	return pool, nil
}
