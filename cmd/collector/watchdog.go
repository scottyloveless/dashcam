package main

import "context"

func (app *application) watchdog() {
	ctx := context.Background()

	alerts, err := app.queries.GetAlerts(ctx)
	if err != nil {
		app.logger.Error(err.Error())
		return
	}

	for _, alert := range alerts {
	}
}
