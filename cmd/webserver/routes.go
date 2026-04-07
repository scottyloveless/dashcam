package main

import (
	"net/http"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	mux.HandleFunc("GET /v1/healthcheck", app.healthcheckHandler)
	mux.HandleFunc("GET /login", app.loginGetHandler)
	mux.HandleFunc("POST /login", app.loginPostHandler)
	mux.Handle("GET /", app.requireAuth(http.HandlerFunc(app.homepageHandler)))
	mux.Handle("GET /device_partial", app.requireAuth(http.HandlerFunc(app.devicePartialHandler)))
	mux.Handle("GET /alerts_partial", app.requireAuth(http.HandlerFunc(app.alertsPartialHandler)))
	mux.Handle("GET /device/{id}", app.requireAuth(http.HandlerFunc(app.devicePageHandler)))
	mux.Handle("POST /logout", app.requireAuth(http.HandlerFunc(app.logoutHandler)))

	return mux
}
