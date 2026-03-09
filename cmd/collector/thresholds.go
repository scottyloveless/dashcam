package main

import (
	"github.com/scottyloveless/dashcam/internal/database"
)

func (app *application) evaluateThreshold(value float64, threshold database.Threshold) *database.SeverityEnum {
	var severityReturn database.SeverityEnum
	switch threshold.Direction {
	case "above":
		if value > threshold.CriticalValue {
			severityReturn = database.SeverityEnumCritical
			return &severityReturn
		} else if value > threshold.WarningValue {
			severityReturn = database.SeverityEnumWarning
			return &severityReturn
		}
		return nil
	case "below":
		if value < threshold.CriticalValue {
			severityReturn = database.SeverityEnumCritical
			return &severityReturn
		} else if value < threshold.WarningValue {
			severityReturn = database.SeverityEnumWarning
			return &severityReturn
		}
		return nil
	case "both":
		if value < threshold.WarningValue || value > threshold.CriticalValue {
			severityReturn = database.SeverityEnumCritical
			return &severityReturn
		}
		return nil
	default:
		return nil
	}
}
