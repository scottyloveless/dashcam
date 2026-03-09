package main

import (
	"context"
	"fmt"
	"net/http"
)

func (app *application) devicePartialHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	query := r.URL.Query().Get("location")

	devices, err := app.queries.GetDevicesOneLocation(r.Context(), query)
	if err != nil {
		app.logger.Error(err.Error())
		return
	}

	var dp []devicePingAndLocation

	for _, device := range devices {
		pl, err := app.queries.GetPacketLossByDeviceID(ctx, device.ID)
		if err != nil {
			fmt.Println(err.Error())
		}
		rtt, err := app.queries.GetRttAvgByDeviceID(ctx, device.ID)
		if err != nil {
			fmt.Println(err.Error())
		}
		dap := devicePingAndLocation{
			Device:     device,
			PacketLoss: pl.Value,
			RTTavg:     rtt.Value,
			Location:   device.Location,
		}

		dp = append(dp, dap)
	}

	err = app.templates["devicePartial"].Execute(w, dp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
