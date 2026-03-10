package main

import (
	"net/http"
)

func (app *application) alertsPartialHandler(w http.ResponseWriter, r *http.Request) {
	alerts, err := app.queries.GetAlerts(r.Context())
	if err != nil {
		app.logger.Error(err.Error())
		return
	}

	err = app.templates["alertsPartial"].Execute(w, alerts)
	if err != nil {
		app.logger.Error(err.Error())
		return
	}
}
