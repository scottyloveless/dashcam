package main

import (
	"context"
	"html/template"
	"net/http"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/julienschmidt/httprouter"
)

func (app *application) devicePageHandler(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	unparsedUUID := params.ByName("id")
	parsedUUID, err := uuid.Parse(unparsedUUID)
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

	tpl, err := template.ParseFiles("cmd/webserver/templates/device_template.html")
	if err != nil {
		app.logger.Error(err.Error())
		return
	}

	if err = tpl.Execute(w, metrics); err != nil {
		app.logger.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
