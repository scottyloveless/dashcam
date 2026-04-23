package ninja

import (
	"context"
	"encoding/json"
	"io"
	"time"
)

type NinjaAlerts []struct {
	UID             string    `json:"uid,omitzero"`
	DeviceID        int       `json:"deviceId,omitzero"`
	Message         time.Time `json:"message,omitzero"`
	CreateTime      float64   `json:"createTime,omitzero"`
	UpdateTime      float64   `json:"updateTime,omitzero"`
	SourceType      string    `json:"sourceType,omitzero"`
	SourceConfigUID string    `json:"sourceConfigUid,omitzero"`
	SourceName      string    `json:"sourceName,omitzero"`
	Subject         string    `json:"subject,omitzero"`
	Data            struct {
		Message struct {
			Code   string `json:"code,omitzero"`
			Params struct {
				Volume    string    `json:"volume,omitzero"`
				Unit      string    `json:"unit,omitzero"`
				Threshold string    `json:"threshold,omitzero"`
				StartDate time.Time `json:"start_date,omitzero"`
				EndDate   time.Time `json:"end_date,omitzero"`
			} `json:"params,omitzero"`
		} `json:"message,omitzero"`
	} `json:"data,omitzero"`
	Priority              string `json:"priority,omitzero"`
	Severity              string `json:"severity,omitzero"`
	ConditionName         string `json:"conditionName,omitzero"`
	ConditionHealthStatus string `json:"conditionHealthStatus,omitzero"`
	UseGlobalHealthStatus bool   `json:"useGlobalHealthStatus,omitzero"`
}

func GetAlerts() ([]NinjaAlerts, error) {
	c, err := NewClient(context.Background())
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Get(c.baseURL + "/v2/alerts")
	if err != nil {
		return nil, err
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var alerts []NinjaAlerts

	err = json.Unmarshal(data, &alerts)
	if err != nil {
		return nil, err
	}

	return alerts, nil
}
