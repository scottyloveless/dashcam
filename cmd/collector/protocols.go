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
	ctx := context.Background()
	dp, err := app.queries.GetProtocolsDevices(ctx)
	if err != nil {
		app.logger.Error(err.Error())
		return nil, err
	}
	// TODO: do I need to check for zero size slice here?

	var triggerSlice []Trigger

	for _, protocol := range dp {
		ip, err := app.queries.GetIPfromDeviceID(ctx, protocol.DeviceID)
		if err != nil {
			app.logger.Error(err.Error())
			return nil, err
		}
		triggerSlice = append(triggerSlice, Trigger{Trigger: protocol, IP: ip})
	}

	return triggerSlice, nil
}
