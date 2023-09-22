package transform

import (
	"fmt"
	"testing"

	"github.com/ttnesby/slack-block-builder/pkg/azure/alert"
)

// iff keen on verify the visual result
// https://app.slack.com/block-kit-builder/

func TestTransform(t *testing.T) {

	alertSevs := []string{"Sev0", "Sev1", "Sev2", "Sev3", "Sev4", "unknown"}

	alertJson := func(sev string) string {
		return fmt.Sprintf(`{
			"schemaId":"azureMonitorCommonAlertSchema",
			"data":{
				"essentials":{
					"alertId":"alertId-1",
					"alertRule":"Test-Rule-1",
					"severity":"%s",
					"signalType":"Metric",
					"monitorCondition":"Fired",
					"monitoringService":"Platform",
					"alertTargetIDs":["resourceId-1"],
					"configurationItems":["resourceId-1-lastElem"],
					"originAlertId":"alertid-2",
					"firedDateTime":"2019-03-22T13:58:24.3713213Z",
					"resolvedDateTime":"2019-03-22T14:03:16.2246313Z",
					"description":"",
					"essentialsVersion":"1.0",
					"alertContextVersion":"1.0"
				}
			}
		}`, sev)
	}

	expectedNotificationJson := func(sev string) string {
		return fmt.Sprintf(`{"blocks":[{"type":"divider"},{"type":"header","text":{"type":"plain_text","text":"New monitor alert!","emoji":true}},{"type":"divider"},{"type":"actions","elements":[{"type":"button","text":{"type":"plain_text","text":"View alert in Monitor","emoji":true},"url":"https://portal.azure.com/#blade/Microsoft_Azure_Monitoring/AlertDetailsTemplateBlade/alertId/alertId-1","style":"primary"}]},{"type":"header","text":{"type":"plain_text","text":"Summary","emoji":true}},{"type":"divider"},{"type":"section","fields":[{"type":"mrkdwn","text":"Fired (UTC)"},{"type":"plain_text","text":"2019-03-22T13:58:24.3713213Z","emoji":true},{"type":"mrkdwn","text":"Alert name"},{"type":"plain_text","text":"Test-Rule-1","emoji":true},{"type":"mrkdwn","text":"Severity"},{"type":"plain_text","text":"%s","emoji":true}]},{"type":"header","text":{"type":"plain_text","text":"Resources","emoji":true}},{"type":"divider"},{"type":"section","fields":[{"type":"mrkdwn","text":"\u003chttps://portal.azure.com/#@nav.no/resourceresourceId-1|resourceId-1\u003e"},{"type":"plain_text","text":":link:","emoji":true}]}]}`, sev)
	}

	for _, s := range alertSevs {
		alert, err := alert.Parse(alertJson(s))
		if err != nil {
			t.Error(err)
		}
		expected := expectedNotificationJson(string(severity(alert)))
		got := AlertToNotification(alert).Json()

		if string(got) != expected {
			t.Errorf("Severity %s, got %s, expected %s", s, got, expected)
		}
	}
}
