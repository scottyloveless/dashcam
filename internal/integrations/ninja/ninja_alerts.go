package ninja

import (
	"encoding/json"
	"io"
	"strconv"
	"time"
)

type NinjaAlerts struct {
	UID             string  `json:"uid,omitzero"`
	DeviceID        int     `json:"deviceId,omitzero"`
	DeviceName      string  `json:"device_name,omitzero"`
	Message         string  `json:"message,omitzero"`
	CreateTime      float64 `json:"createTime,omitzero"`
	UpdateTime      float64 `json:"updateTime,omitzero"`
	SourceType      string  `json:"sourceType,omitzero"`
	SourceConfigUID string  `json:"sourceConfigUid,omitzero"`
	SourceName      string  `json:"sourceName,omitzero"`
	Subject         string  `json:"subject,omitzero"`
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

type NinjaDevice struct {
	ID             int      `json:"id,omitzero"`
	UID            string   `json:"uid,omitzero"`
	OrganizationID int      `json:"organizationId,omitzero"`
	LocationID     int      `json:"locationId,omitzero"`
	NodeClass      string   `json:"nodeClass,omitzero"`
	NodeRoleID     int      `json:"nodeRoleId,omitzero"`
	RolePolicyID   int      `json:"rolePolicyId,omitzero"`
	PolicyID       int      `json:"policyId,omitzero"`
	ApprovalStatus string   `json:"approvalStatus,omitzero"`
	Offline        bool     `json:"offline,omitzero"`
	SystemName     string   `json:"systemName,omitzero"`
	DNSName        string   `json:"dnsName,omitzero"`
	Created        float64  `json:"created,omitzero"`
	LastContact    float64  `json:"lastContact,omitzero"`
	LastUpdate     float64  `json:"lastUpdate,omitzero"`
	IPAddresses    []string `json:"ipAddresses,omitzero"`
	MacAddresses   []string `json:"macAddresses,omitzero"`
	PublicIP       string   `json:"publicIP,omitzero"`
	Os             struct {
		Manufacturer            string  `json:"manufacturer,omitzero"`
		Name                    string  `json:"name,omitzero"`
		Architecture            string  `json:"architecture,omitzero"`
		LastBootTime            float64 `json:"lastBootTime,omitzero"`
		BuildNumber             string  `json:"buildNumber,omitzero"`
		ReleaseID               string  `json:"releaseId,omitzero"`
		ServicePackMajorVersion int     `json:"servicePackMajorVersion,omitzero"`
		ServicePackMinorVersion int     `json:"servicePackMinorVersion,omitzero"`
		Locale                  string  `json:"locale,omitzero"`
		Language                string  `json:"language,omitzero"`
		NeedsReboot             bool    `json:"needsReboot,omitzero"`
	} `json:"os,omitzero"`
	System struct {
		Name                string `json:"name,omitzero"`
		Manufacturer        string `json:"manufacturer,omitzero"`
		Model               string `json:"model,omitzero"`
		BiosSerialNumber    string `json:"biosSerialNumber,omitzero"`
		SerialNumber        string `json:"serialNumber,omitzero"`
		Domain              string `json:"domain,omitzero"`
		DomainRole          string `json:"domainRole,omitzero"`
		NumberOfProcessors  int    `json:"numberOfProcessors,omitzero"`
		TotalPhysicalMemory int64  `json:"totalPhysicalMemory,omitzero"`
		VirtualMachine      bool   `json:"virtualMachine,omitzero"`
		ChassisType         string `json:"chassisType,omitzero"`
	} `json:"system,omitzero"`
	Memory struct {
		Capacity int64 `json:"capacity,omitzero"`
	} `json:"memory,omitzero"`
	Volumes []struct {
		Name         string `json:"name,omitzero"`
		Label        string `json:"label,omitzero"`
		DeviceType   string `json:"deviceType,omitzero"`
		FileSystem   string `json:"fileSystem,omitzero"`
		AutoMount    bool   `json:"autoMount,omitzero"`
		Compressed   bool   `json:"compressed,omitzero"`
		Capacity     int64  `json:"capacity,omitzero"`
		FreeSpace    int64  `json:"freeSpace,omitzero"`
		SerialNumber string `json:"serialNumber,omitzero"`
	} `json:"volumes,omitzero"`
	LastLoggedInUser string `json:"lastLoggedInUser,omitzero"`
	DeviceType       string `json:"deviceType,omitzero"`
}

func (c *Client) GetAlerts() ([]NinjaAlerts, error) {
	resp, err := c.HTTPClient.Get(c.BaseURL + "/v2/alerts")
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

	for i := range alerts {
		resp, err := c.HTTPClient.Get(c.BaseURL + "/v2/device/" + strconv.Itoa(alerts[i].DeviceID))
		if err != nil {
			return nil, err
		}

		data, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		var device NinjaDevice

		err = json.Unmarshal(data, &device)
		if err != nil {
			return nil, err
		}

		alert := &alerts[i]
		alert.DeviceName = device.DNSName
	}

	return alerts, nil
}
