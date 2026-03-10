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

type homePageData struct {
	Locations []string
}

func (app *application) homepageHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	locationStrSlice, err := app.queries.GetDistinctLocations(ctx)
	if err != nil {
		app.logger.Error(err.Error())
		return
	}

	hpData := homePageData{
		Locations: locationStrSlice,
	}

	err = app.templates["home"].Execute(w, hpData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
