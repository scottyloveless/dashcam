package main

import (
	"net/http"
	"time"
)

func (app *application) handleExternalAlerts(w http.ResponseWriter, r *http.Request) {
	alerts, err := app.queries.ListOpenExternalAlerts(r.Context())
	if err != nil {
		app.logger.Error("list external alerts failed", "err", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	// Build view model with friendly relative timestamps
	type alertRow struct {
		ID           string
		Source       string
		DeviceName   string
		Message      string
		Severity     string
		State        string
		RelativeTime string
	}

	now := time.Now()
	rows := make([]alertRow, 0, len(alerts))
	for _, a := range alerts {
		ts := a.LastOccurrence.Time
		if !a.LastOccurrence.Valid {
			ts = a.CreatedAt.Time
		}
		rows = append(rows, alertRow{
			ID:           a.ID.String(),
			Source:       a.Source,
			DeviceName:   a.ExternalDeviceName.String,
			Message:      a.DisplayMessage,
			Severity:     string(a.Severity),
			State:        string(a.State),
			RelativeTime: relativeTime(now, ts),
		})
	}

	data := struct {
		Count  int
		Alerts []alertRow
	}{
		Count:  len(rows),
		Alerts: rows,
	}

	if err := app.templates["external_alerts_partial"].Execute(w, data); err != nil {
		app.logger.Error("render external alerts failed", "err", err)
	}
}
