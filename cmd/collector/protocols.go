package main

import (
	"context"
	"errors"
	"net/netip"

	"github.com/scottyloveless/dashcam/internal/database"
)

type Trigger struct {
	Trigger database.DevicesProtocol
	IP      netip.Addr
	Type    string
}

func (app *application) triggerNetworkHelper() ([]Trigger, error) {
	ctx := context.Background()
	dp, err := app.queries.GetProtocolsDevices(ctx)
	if err != nil {
		app.logger.Error(err.Error())
		return nil, err
	}
	if len(dp) == 0 {
		app.logger.Error("no triggers found")
		return nil, errors.New("no triggers found")
	}

	var triggerSlice []Trigger

	for _, protocol := range dp {
		dev, err := app.queries.GetIPandTypefromDeviceID(ctx, protocol.DeviceID)
		if err != nil {
			app.logger.Error(err.Error())
			return nil, err
		}
		triggerSlice = append(triggerSlice, Trigger{Trigger: protocol, IP: dev.IpAddress, Type: dev.Type})
	}

	return triggerSlice, nil
}
