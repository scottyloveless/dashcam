package main

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/scottyloveless/dashcam/internal/database"

	"github.com/jackc/pgx/v5/pgxpool"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
}

type application struct {
	config       config
	logger       *slog.Logger
	dbpool       *pgxpool.Pool
	queries      *database.Queries
	activeDevice database.Device
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "web server port")
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

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		ErrorLog:     slog.NewLogLogger(logger.Handler(), slog.LevelError),
	}

	logger.Info("starting server", "addr", srv.Addr, "env", cfg.env)

	err = srv.ListenAndServe()
	logger.Error(err.Error())
	os.Exit(1)
}
