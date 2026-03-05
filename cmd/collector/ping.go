package main

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	probing "github.com/prometheus-community/pro-bing"
	"github.com/scottyloveless/dashcam/internal/database"
)

func (app *application) collectPing(trigger Trigger) error {
	pinger, err := probing.NewPinger(trigger.IP.String())
	if err != nil {
		app.logger.Error(err.Error())
		return err
	}

	pinger.Count = 5
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
		Value:      stats.PacketLoss / 100,
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
	}

	statsPayload2 := database.WritePingParams{
		MetricName: "rtt_avg",
		Value:      float64(stats.AvgRtt.Milliseconds()),
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
	}

	err = app.queries.WritePing(context.Background(), statsPayload)
	if err != nil {
		app.logger.Error(err.Error())
		return err
	}

	err = app.queries.WritePing(context.Background(), statsPayload2)
	if err != nil {
		app.logger.Error(err.Error())
		return err
	}
	return nil
}
