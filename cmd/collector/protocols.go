package main

import (
	"context"
	"net/netip"

	"github.com/scottyloveless/dashcam/internal/database"
)

type Trigger struct {
	Trigger database.DevicesProtocol
	IP      netip.Addr
}

func (app *application) triggerNetworkHelper() ([]Trigger, error) {
	dp, err := app.queries.GetProtocolsDevices(context.Background())
	if err != nil {
		app.logger.Error(err.Error())
		return nil, err
	}
	// TODO: do I need to check for zero size slice here?

	var triggerSlice []Trigger

	for _, protocol := range dp {
		ip, err := app.queries.GetIPfromDeviceID(context.Background(), protocol.DeviceID)
		if err != nil {
			app.logger.Error(err.Error())
			return nil, err
		}
		triggerSlice = append(triggerSlice, Trigger{Trigger: protocol, IP: ip})
	}

	return triggerSlice, nil
}
