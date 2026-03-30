package main

import (
	"context"
	"flag"
	"fmt"
	"html/template"
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
	db   struct {
		dsn          string
		maxOpenConns int
		maxIdleConns int
		maxIdleTime  time.Duration
	}
}

type application struct {
	config    config
	logger    *slog.Logger
	dbpool    *pgxpool.Pool
	queries   *database.Queries
	templates map[string]*template.Template
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "web server port")
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

	logger.Info("database connection successful")

	pool, err := openDB(cfg)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	queries := database.New(pool)

	templates := make(map[string]*template.Template)
	hometemplate, err := template.ParseFiles("cmd/webserver/templates/home_template.html")
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	devicePartialTemplate, err := template.ParseFiles("cmd/webserver/partials/device_partial.html")
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	deviceInfoTemplate, err := template.ParseFiles("cmd/webserver/templates/deviceInfo_template.html")
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	alertPartialTemplate, err := template.ParseFiles("cmd/webserver/partials/alerts_partial.html")
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	templates["home"] = hometemplate
	templates["devicePartial"] = devicePartialTemplate
	templates["alertsPartial"] = alertPartialTemplate
	templates["deviceInfo"] = deviceInfoTemplate

	app := &application{
		config:    cfg,
		logger:    logger,
		dbpool:    pool,
		queries:   queries,
		templates: templates,
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

func openDB(cfg config) (*pgxpool.Pool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	config, err := pgxpool.ParseConfig(cfg.db.dsn)
	if err != nil {
		return nil, err
	}
	config.MaxConns = int32(cfg.db.maxOpenConns)
	config.MaxConnIdleTime = cfg.db.maxIdleTime

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, err
	}
	defer pool.Close()

	err = pool.Ping(ctx)
	if err != nil {
		defer pool.Close()
		return nil, err
	}

	return pool, nil
}
