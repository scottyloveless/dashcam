package main

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/scottyloveless/dashcam/internal/database"
)

func (app *application) watchdog(alert database.GetAlertsRow, ctx context.Context) {
	reqPayload := database.GetLastFiveMetricsByDeviceIDParams{
		MetricName: alert.AlertMetric,
		DeviceID:   alert.DeviceID,
	}

	ipAndType, err := app.queries.GetIPandTypefromDeviceID(ctx, alert.DeviceID)
	if err != nil {
		app.logger.Error(err.Error())
		return
	}

	lastFive, err := app.queries.GetLastFiveMetricsByDeviceID(ctx, reqPayload)
	if err != nil {
		app.logger.Error(err.Error())
		return
	}

	if len(lastFive) < 5 {
		return
	}

	threshParams := database.GetActiveThresholdParams{
		DeviceID: alert.DeviceID,
		DeviceType: pgtype.Text{
			String: ipAndType.Type,
			Valid:  true,
		},
		Metric: alert.AlertMetric,
	}

	thresh, err := app.queries.GetActiveThreshold(ctx, threshParams)
	if err != nil {
		app.logger.Error(err.Error())
		return
	}

	for _, metric := range lastFive {
		sev := app.evaluateThreshold(metric.Value, thresh)
		if sev != nil {
			return
		}
	}

	err = app.queries.ClearAlert(ctx, alert.ID)
	if err != nil {
		app.logger.Error(err.Error())
	}
}
