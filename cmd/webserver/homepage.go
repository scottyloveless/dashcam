package main

import (
	"context"
	"html/template"
	"net/http"

	"github.com/scottyloveless/dashcam/internal/database"
)

type deviceAndPing struct {
	Device database.Device	
	Ping PingPair
}

type PingPair struct {
	RttAvg float64
	PacketLoss float64
}

func (app *application) homepageHandler(w http.ResponseWriter, r *http.Request) {
	devices, err := app.queries.GetDevices(r.Context())
	if err != nil {
		app.logger.Error(err.Error())
		return
	}

	app.queries.GetPingByDeviceID(context.Background(), device.)

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
