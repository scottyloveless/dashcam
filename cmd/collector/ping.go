package main

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	probing "github.com/prometheus-community/pro-bing"
	"github.com/scottyloveless/dashcam/internal/database"
)

func (app *application) collectPing(trigger Trigger) error {
	ctx := context.Background()
	pinger, err := probing.NewPinger(trigger.IP.String())
	if err != nil {
		app.logger.Error(err.Error())
		return err
	}

	pinger.Count = 4
	pinger.Timeout = 1 * time.Second

	requestedTime := time.Now()

	err = pinger.Run()

	receivedTime := time.Now()

	if err != nil {
		app.logger.Error(err.Error())
		return err
	}

	stats := pinger.Statistics()

	statsPayload := database.WritePingParams{
		MetricName: "packet_loss",
		Value:      stats.PacketLoss,
		DeviceID:   trigger.Trigger.DeviceID,
		RequestedAt: pgtype.Timestamptz{
			Time:             requestedTime,
			InfinityModifier: 0,
			Valid:            true,
		},
		ReceivedAt: pgtype.Timestamptz{
			Time:             receivedTime,
			InfinityModifier: 0,
			Valid:            true,
		},
		MetricName_2: "rtt_avg",
		Value_2:      float64(stats.AvgRtt.Milliseconds()),
		DeviceID_2:   trigger.Trigger.DeviceID,
		RequestedAt_2: pgtype.Timestamptz{
			Time:             requestedTime,
			InfinityModifier: 0,
			Valid:            true,
		},
		ReceivedAt_2: pgtype.Timestamptz{
			Time:             receivedTime,
			InfinityModifier: 0,
			Valid:            true,
		},
	}

	err = app.queries.WritePing(ctx, statsPayload)
	if err != nil {
		app.logger.Error(err.Error())
		return err
	}

	err = app.evaluateAndAlert(ctx, trigger, "rtt_avg", float64(stats.AvgRtt.Milliseconds()))
	if err != nil {
		app.logger.Error(err.Error())
		return err
	}
	err = app.evaluateAndAlert(ctx, trigger, "packet_loss", stats.PacketLoss)
	if err != nil {
		app.logger.Error(err.Error())
		return err
	}
	return nil
}
