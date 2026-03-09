package main

import (
	"context"
	"net/http"

	"github.com/scottyloveless/dashcam/internal/database"
)

type devicePingAndLocation struct {
	Device     database.Device
	Location   string
	PacketLoss float64
	RTTavg     float64
}

func (app *application) homepageHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	locations, err := app.queries.GetDistinctLocations(ctx)
	if err != nil {
		app.logger.Error(err.Error())
		return
	}

	err = app.templates["home"].Execute(w, locations)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
