package main

import (
	"context"
	"flag"
	"fmt"
	"html/template"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/scottyloveless/dashcam/internal/database"

	"github.com/jackc/pgx/v5"
)

// const version = "1.0.0"
type config struct {
	port int
	env  string
}

//
// type application struct {
// 	config config
// 	logger *slog.Logger
// }

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "web server port")
	flag.StringVar(&cfg.env, "env", "development", "development|staging|production")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	// app := application{
	// 	config: cfg,
	// 	logger: logger,
	// }

	err := godotenv.Load(".env")
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	dburl := os.Getenv("DATABASE_URL")

	ctx := context.Background()
	conn, err := pgx.Connect(ctx, dburl)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	defer func() {
		cerr := conn.Close(ctx)
		if cerr != nil && err == nil {
			err = cerr
		}
	}()

	db := database.New(conn)

	devices, err := db.GetDevices(context.Background())
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	tpl, err := template.ParseFiles("test_template.html")
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err = tpl.Execute(w, devices)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	fmt.Println("Starting webserver on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
	fmt.Println("Stopping webserver...")

	// mux := http.NewServeMux()
	// mux.HandleFunc("/v1/healthcheck", app.healthcheckHandler)
	//
	// srv := &http.Server{
	// 	Addr:         fmt.Sprintf(":%d", cfg.port),
	// 	Handler:      app.routes(),
	// 	IdleTimeout:  time.Minute,
	// 	ReadTimeout:  5 * time.Second,
	// 	WriteTimeout: 10 * time.Second,
	// 	ErrorLog:     slog.NewLogLogger(logger.Handler(), slog.LevelError),
	// }
	//
	// logger.Info("starting server", "addr", srv.Addr, "env", cfg.env)
	//
	// err = srv.ListenAndServe()
	// logger.Error(err.Error())
	os.Exit(1)
}
