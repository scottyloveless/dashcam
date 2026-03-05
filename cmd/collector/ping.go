package main

import (
	probing "github.com/prometheus-community/pro-bing"
)

func (app *application) sendPing() *probing.Statistics {
	pinger, err := probing.NewPinger("google.com")
	if err != nil {
		app.logger.Error(err.Error())
		return nil
	}

	pinger.Count = 5
	err = pinger.Run()
	if err != nil {
		app.logger.Error(err.Error())
		return nil
	}

	return pinger.Statistics()
}
