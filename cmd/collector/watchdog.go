package main

// import (
// 	"context"
// 	"time"
// )
//
// func (app *application) watchdog() {
// 	ctx := context.Background()
//
// 	ticker := time.NewTicker(5 * time.Second)
// 	defer ticker.Stop()
//
// 	for range ticker.C {
// 		alerts, err := app.queries.GetAlerts(ctx)
// 		if err != nil {
// 			app.logger.Error(err.Error())
// 			return
// 		}
//
// 		for _, alert := range alerts {
// 		}
// 	}
// 	// 1. ticker or sleep - i'm thinking ticker
// 	// 2. fetch list of alerts
// 	// 3. fetch last 5-10 metrics for that device and metric name, sort descending
// 	// 4. if last metrics look good, clear the alert
// 	// 5. if last metrics look bad, keep alert
// }
