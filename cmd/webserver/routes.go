package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.ServeFiles("/static/*filepath", http.Dir("static"))

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodGet, "/", app.homepageHandler)
	router.HandlerFunc(http.MethodGet, "/device_partial", app.devicePartialHandler)
	router.HandlerFunc(http.MethodGet, "/device/:id", app.devicePageHandler)

	return router
}
