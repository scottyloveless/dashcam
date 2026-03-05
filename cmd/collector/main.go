package main

import (
	"context"
	"flag"
	"log/slog"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/scottyloveless/dashcam/internal/database"
)

// const version = "1.0.0"

type config struct {
	env string
}

type application struct {
	config  config
	logger  *slog.Logger
	dbpool  *pgxpool.Pool
	queries *database.Queries
}

func main() {
	var cfg config

	flag.StringVar(&cfg.env, "env", "development", "development|staging|production")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	err := godotenv.Load(".env")
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	dburl := os.Getenv("DATABASE_URL")

	ctx := context.Background()
	pool, err := pgxpool.New(ctx, dburl)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	defer pool.Close()

	err = pool.Ping(ctx)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	logger.Info("database connection successful")

	queries := database.New(pool)

	app := application{
		config:  cfg,
		logger:  logger,
		dbpool:  pool,
		queries: queries,
	}
	for {
		requestedTime := time.Now()
		stats := app.sendPing()
		receivedTime := time.Now()

		statsPayload := database.WritePingParams{
			MetricName: "packet_loss",
			Value:      stats.PacketLoss,
			RequestedAt: pgtype.Timestamptz{
				Time:             requestedTime,
				InfinityModifier: 0,
				Valid:            true,
			},
			ReceivedAt: pgtype.Timestamptz{
				Time:             receivedTime,
				InfinityModifier: 0,
				Valid:            true,
			},
		}

		statsPayload2 := database.WritePingParams{
			MetricName: "rtt_avg",
			Value:      float64(stats.AvgRtt),
			RequestedAt: pgtype.Timestamptz{
				Time:             requestedTime,
				InfinityModifier: 0,
				Valid:            true,
			},
			ReceivedAt: pgtype.Timestamptz{
				Time:             receivedTime,
				InfinityModifier: 0,
				Valid:            true,
			},
		}

		err = app.queries.WritePing(context.Background(), statsPayload)
		if err != nil {
			app.logger.Error(err.Error())
			os.Exit(1)
		}

		err = app.queries.WritePing(context.Background(), statsPayload2)
		if err != nil {
			app.logger.Error(err.Error())
			os.Exit(1)
		}
		app.logger.Info("successfully wrote to database")
		app.logger.Info("waiting 10 seconds")
		time.Sleep(10 * time.Second)
	}
}
