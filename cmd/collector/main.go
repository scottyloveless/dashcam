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
}

type application struct {
	config    config
	logger    *slog.Logger
	dbpool    *pgxpool.Pool
	queries   *database.Queries
	protocols []Trigger
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
		config:    cfg,
		logger:    logger,
		dbpool:    pool,
		queries:   queries,
		protocols: []Trigger{},
	}

	for {
		fmt.Println("starting collector cycle...")
		app.protocols, err = app.triggerNetworkHelper()
		if err != nil {
			logger.Error(err.Error())
			os.Exit(1)
		}

		if len(app.protocols) <= 0 {
			time.Sleep(10 * time.Second)
			continue
		}

		for _, protocol := range app.protocols {
			if !protocol.Trigger.Enabled {
				continue
			}

			go func() {
				n := rand.IntN(10)

				time.Sleep(time.Duration(n) * time.Millisecond)
				if err = app.collectPing(protocol); err != nil {
					app.logger.Error(err.Error())
					return
				}
				fmt.Println("ping collected from " + protocol.IP.String())
			}()
		}
		time.Sleep(5 * time.Second)
	}
}
