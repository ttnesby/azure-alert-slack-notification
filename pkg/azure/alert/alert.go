package alert

import (
	"encoding/json"
	"fmt"
)

// https://learn.microsoft.com/en-us/azure/azure-monitor/alerts/alerts-common-schema#sample-alert-payload

const (
	UrlAlertBlade    = "https://portal.azure.com/#blade/Microsoft_Azure_Monitoring/AlertDetailsTemplateBlade/alertId/"
	UrlResourceBlade = "https://portal.azure.com/#@nav.no/resource"
)

type Content struct {
	AlertId             string   `json:"alertId"`
	AlertRule           string   `json:"alertRule"`
	Severity            string   `json:"severity"`
	SignalType          string   `json:"signalType"`
	MonitorCondition    string   `json:"monitorCondition"`
	MonitoringService   string   `json:"monitoringService"`
	AlertTargetIDs      []string `json:"alertTargetIDs"`
	ConfigurationItems  []string `json:"configurationItems"`
	OriginAlertId       string   `json:"originAlertId"`
	FiredDateTime       string   `json:"firedDateTime"`
	ResolvedDateTime    string   `json:"resolvedDateTime"`
	Description         string   `json:"description"`
	EssentialsVersion   string   `json:"essentialsVersion"`
	AlertContextVersion string   `json:"alertContextVersion"`
}

type Essentials struct {
	Essentials Content `json:"essentials"`
}

type CommonAlertSchema struct {
	SchemaId string     `json:"schemaId"`
	Data     Essentials `json:"data"`
}

func Parse(s string) *CommonAlertSchema {
	var alert CommonAlertSchema
	err := json.Unmarshal([]byte(s), &alert)

	if err != nil {
		fmt.Println(fmt.Errorf("unmarshal failed for %s", s))
	}

	return &alert
}
