package main

import (
	"context"
	"fmt"
	"html/template"
	"net/http"
)

func (app *application) devicePartialHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	devices, err := app.queries.GetDevices(r.Context())
	if err != nil {
		app.logger.Error(err.Error())
		return
	}

	var dp []deviceAndPing

	for _, device := range devices {
		pl, err := app.queries.GetPacketLossByDeviceID(ctx, device.ID)
		if err != nil {
			fmt.Println(err.Error())
		}
		rtt, err := app.queries.GetRttAvgByDeviceID(ctx, device.ID)
		if err != nil {
			fmt.Println(err.Error())
		}
		dap := deviceAndPing{
			Device:     device,
			PacketLoss: pl.Value,
			RTTavg:     rtt.Value,
		}

		dp = append(dp, dap)
	}

	tpl, err := template.ParseFiles("cmd/webserver/partials/device_partial.html")
	if err != nil {
		app.logger.Error(err.Error())
		return
	}
	err = tpl.Execute(w, dp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
