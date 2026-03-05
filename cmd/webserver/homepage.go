package main

import (
	"html/template"
	"net/http"
)

func (app *application) homepageHandler(w http.ResponseWriter, r *http.Request) {
	devices, err := app.queries.GetDevices(r.Context())
	if err != nil {
		app.logger.Error(err.Error())
		return
	}

	tpl, err := template.ParseFiles("cmd/server/test_template.html")
	if err != nil {
		app.logger.Error(err.Error())
		return
	}
	err = tpl.Execute(w, devices)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
