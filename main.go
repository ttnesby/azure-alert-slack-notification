package main

import (
	"fmt"
	"github.com/ttnesby/slack-block-builder/pkg/azure/alert"
	"github.com/ttnesby/slack-block-builder/pkg/transform"
)

func main() {
	json := `
{
  "schemaId": "azureMonitorCommonAlertSchema",
  "data": {
    "essentials": {
      "alertId": "/subscriptions/9876/providers/Microsoft.AlertsManagement/alerts/b9569717-bc32-442f-add5-83a997729330",
      "alertRule": "Test-Rule-1",
      "severity": "Sev0",
      "signalType": "Metric",
      "monitorCondition": "Fired",
      "monitoringService": "Platform",
      "alertTargetIDs": [
        "/subscriptions/1234/resourcegroups/pipelinealertrg/providers/microsoft.compute/virtualmachines/wcus-r2-gen2",
		"/subscriptions/1234/resourcegroups/pipelinealertrg/providers/microsoft.compute/virtualmachines/wcus-r2-gen3",
        "/subscriptions/1234/resourcegroups/pipelinealertrg/providers/microsoft.compute/virtualmachines/wcus-r2-gen4",
		"/subscriptions/1234/resourcegroups/pipelinealertrg/providers/microsoft.compute/virtualmachines/wcus-r2-gen5",
        "/subscriptions/1234/resourcegroups/pipelinealertrg/providers/microsoft.compute/virtualmachines/wcus-r2-gen6"
      ],
      "configurationItems": [
        "wcus-r2-gen2","wcus-r2-gen3","wcus-r2-gen4","wcus-r2-gen5","wcus-r2-gen6"
      ],
      "originAlertId": "3f2d4487-b0fc-4125-8bd5-7ad17384221e_PipeLineAlertRG_microsoft.insights_metricAlerts_WCUS-R2-Gen2_-117781227",
      "firedDateTime": "2019-03-22T13:58:24.3713213Z",
      "resolvedDateTime": "2019-03-22T14:03:16.2246313Z",
      "description": "",
      "essentialsVersion": "1.0",
      "alertContextVersion": "1.0"
    }
  }
}
`

	fmt.Printf("%s", transform.AlertToNotification(alert.Parse(json)))
}
