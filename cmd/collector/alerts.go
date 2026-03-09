package main

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/scottyloveless/dashcam/internal/database"
)

func (app *application) evaluateAndAlert(ctx context.Context, trigger Trigger, metricname string, value float64) error {
	threshParams := database.GetActiveThresholdParams{
		DeviceID: trigger.Trigger.DeviceID,
		DeviceType: pgtype.Text{
			String: trigger.Type,
			Valid:  true,
		},
		Metric: metricname,
	}
	activeThresholds, err := app.queries.GetActiveThreshold(ctx, threshParams)
	if err == pgx.ErrNoRows {
		return nil
	} else if err != nil {
		app.logger.Error(err.Error())
		return err
	} else {
		thresh := app.evaluateThreshold(float64(value), activeThresholds)
		if thresh == nil {
			return nil
		} else {
			checkAlertParams := database.CheckAlertParams{
				DeviceID:    trigger.Trigger.DeviceID,
				AlertMetric: metricname,
			}
			alert, err := app.queries.CheckAlert(ctx, checkAlertParams)
			if err == pgx.ErrNoRows {
				alertParams := database.WriteAlertParams{
					DeviceID:    trigger.Trigger.DeviceID,
					AlertMetric: metricname,
					ThresholdID: activeThresholds.ID,
					Severity:    *thresh,
				}
				err = app.queries.WriteAlert(ctx, alertParams)
				if err != nil {
					app.logger.Error(err.Error())
					return err
				}
			} else if err != nil {
				app.logger.Error(err.Error())
				return err
			} else {
				params := database.UpdateAlertLastOccurrenceParams{
					Severity: *thresh,
					ID:       alert.ID,
				}
				err = app.queries.UpdateAlertLastOccurrence(ctx, params)
				if err != nil {
					app.logger.Error(err.Error())
					return err
				}
			}

		}
	}
	return nil
}
