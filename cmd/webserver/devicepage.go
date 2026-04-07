package main

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

func (app *application) devicePageHandler(w http.ResponseWriter, r *http.Request) {
	params := r.PathValue("id")

	parsedUUID, err := uuid.Parse(params)
	if err != nil {
		app.logger.Error(err.Error())
		return
	}

	pgTypeUUID := pgtype.UUID{
		Bytes: parsedUUID,
		Valid: true,
	}

	ctx := context.Background()
	deviceQuery, err := app.queries.GetOneDeviceInfo(ctx, pgTypeUUID)
	if err != nil {
		app.logger.Error(err.Error())
		return
	}

	metrics, err := app.queries.GetAllMetricsForOneDeviceByID(ctx, deviceQuery.ID)
	if err != nil {
		app.logger.Error(err.Error())
		return
	}

	if err = app.templates["deviceInfo"].Execute(w, metrics); err != nil {
		app.logger.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
