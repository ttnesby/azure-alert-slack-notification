package transform

// outdated test file, kept for reference
// a fix will be to update input and output

// import (
// 	"fmt"
// 	"github.com/ttnesby/slack-block-builder/pkg/azure/alert"
// 	"testing"
// )

// // iff keen on verify the visual result
// // https://app.slack.com/block-kit-builder/

// func TestVerbose(t *testing.T) {

// 	js := `{"schemaId":"azureMonitorCommonAlertSchema","data":{"essentials":{"alertId":"/subscriptions/9876/providers/Microsoft.AlertsManagement/alerts/b9569717-bc32-442f-add5-83a997729330","alertRule":"Test-Rule-1","severity":"Sev4","signalType":"Metric","monitorCondition":"Fired","monitoringService":"Platform","alertTargetIDs":["/subscriptions/1234/resourcegroups/pipelinealertrg/providers/microsoft.compute/virtualmachines/wcus-r2-gen2"],"configurationItems":["wcus-r2-gen2"],"originAlertId":"3f2d4487-b0fc-4125-8bd5-7ad17384221e_PipeLineAlertRG_microsoft.insights_metricAlerts_WCUS-R2-Gen2_-117781227","firedDateTime":"2019-03-22T13:58:24.3713213Z","resolvedDateTime":"2019-03-22T14:03:16.2246313Z","description":"","essentialsVersion":"1.0","alertContextVersion":"1.0"}}}`
// 	got := AlertToNotification(alert.Parse(js)).Json()
// 	expected := `{"blocks":[{"type":"section","text":{"type":"mrkdwn","text":"*Azure monitor alert* :rotating_light:"}},{"type":"divider"},{"type":"actions","elements":[{"type":"button","text":{"type":"plain_text","text":"View alert in Azure Monitor","emoji":true},"url":"https://portal.azure.com/#blade/Microsoft_Azure_Monitoring/AlertDetailsTemplateBlade/alertId/%2Fsubscriptions%2F9876%2Fproviders%2FMicrosoft.AlertsManagement%2Falerts%2Fb9569717-bc32-442f-add5-83a997729330","style":"primary"}]},{"type":"header","text":{"type":"plain_text","text":"Summary","emoji":true}},{"type":"divider"},{"type":"section","fields":[{"type":"mrkdwn","text":"Fired (UTC)"},{"type":"plain_text","text":"2019-03-22T13:58:24.3713213Z","emoji":true},{"type":"mrkdwn","text":"Alert name"},{"type":"plain_text","text":"Test-Rule-1","emoji":true},{"type":"mrkdwn","text":"Severity"},{"type":"plain_text","text":":speech_balloon:  4 - Verbose","emoji":true}]},{"type":"header","text":{"type":"plain_text","text":"Resources","emoji":true}},{"type":"divider"},{"type":"section","fields":[{"type":"mrkdwn","text":"\u003chttps://portal.azure.com/#@nav.no/resource/subscriptions/1234/resourcegroups/pipelinealertrg/providers/microsoft.compute/virtualmachines/wcus-r2-gen2|wcus-r2-gen2\u003e"},{"type":"plain_text","text":":link:","emoji":true}]}]}`

// 	if string(got) != expected {
// 		t.Errorf("%s", "transform failure")
// 	}
// }

// func TestUnknown(t *testing.T) {

// 	js := `{"schemaId":"azureMonitorCommonAlertSchema","data":{"essentials":{"alertId":"/subscriptions/9876/providers/Microsoft.AlertsManagement/alerts/b9569717-bc32-442f-add5-83a997729330","alertRule":"Test-Rule-1","severity":"S75","signalType":"Metric","monitorCondition":"Fired","monitoringService":"Platform","alertTargetIDs":["/subscriptions/1234/resourcegroups/pipelinealertrg/providers/microsoft.compute/virtualmachines/wcus-r2-gen2"],"configurationItems":["wcus-r2-gen2"],"originAlertId":"3f2d4487-b0fc-4125-8bd5-7ad17384221e_PipeLineAlertRG_microsoft.insights_metricAlerts_WCUS-R2-Gen2_-117781227","firedDateTime":"2019-03-22T13:58:24.3713213Z","resolvedDateTime":"2019-03-22T14:03:16.2246313Z","description":"","essentialsVersion":"1.0","alertContextVersion":"1.0"}}}`
// 	got := AlertToNotification(alert.Parse(js))
// 	expected := `{"blocks":[{"type":"section","text":{"type":"mrkdwn","text":"*Azure monitor alert* :rotating_light:"}},{"type":"divider"},{"type":"actions","elements":[{"type":"button","text":{"type":"plain_text","text":"View alert in Azure Monitor","emoji":true},"url":"https://portal.azure.com/#blade/Microsoft_Azure_Monitoring/AlertDetailsTemplateBlade/alertId/%2Fsubscriptions%2F9876%2Fproviders%2FMicrosoft.AlertsManagement%2Falerts%2Fb9569717-bc32-442f-add5-83a997729330","style":"primary"}]},{"type":"header","text":{"type":"plain_text","text":"Summary","emoji":true}},{"type":"divider"},{"type":"section","fields":[{"type":"mrkdwn","text":"Fired (UTC)"},{"type":"plain_text","text":"2019-03-22T13:58:24.3713213Z","emoji":true},{"type":"mrkdwn","text":"Alert name"},{"type":"plain_text","text":"Test-Rule-1","emoji":true},{"type":"mrkdwn","text":"Severity"},{"type":"plain_text","text":":question: unknown","emoji":true}]},{"type":"header","text":{"type":"plain_text","text":"Resources","emoji":true}},{"type":"divider"},{"type":"section","fields":[{"type":"mrkdwn","text":"\u003chttps://portal.azure.com/#@nav.no/resource/subscriptions/1234/resourcegroups/pipelinealertrg/providers/microsoft.compute/virtualmachines/wcus-r2-gen2|wcus-r2-gen2\u003e"},{"type":"plain_text","text":":link:","emoji":true}]}]}`

// 	if fmt.Sprintf("%s", got) != expected {
// 		t.Errorf("%s", "transform failure")
// 	}

// 	fmt.Printf("%s\n", expected)
// 	fmt.Printf("%s\n", fmt.Sprintf("%s", got))

// }

// func Test5Resources(t *testing.T) {

// 	js := `{"schemaId":"azureMonitorCommonAlertSchema","data":{"essentials":{"alertId":"/subscriptions/9876/providers/Microsoft.AlertsManagement/alerts/b9569717-bc32-442f-add5-83a997729330","alertRule":"Test-Rule-1","severity":"Sev0","signalType":"Metric","monitorCondition":"Fired","monitoringService":"Platform","alertTargetIDs":["/subscriptions/1234/resourcegroups/pipelinealertrg/providers/microsoft.compute/virtualmachines/wcus-r2-gen2","/subscriptions/1234/resourcegroups/pipelinealertrg/providers/microsoft.compute/virtualmachines/wcus-r2-gen3","/subscriptions/1234/resourcegroups/pipelinealertrg/providers/microsoft.compute/virtualmachines/wcus-r2-gen4","/subscriptions/1234/resourcegroups/pipelinealertrg/providers/microsoft.compute/virtualmachines/wcus-r2-gen5","/subscriptions/1234/resourcegroups/pipelinealertrg/providers/microsoft.compute/virtualmachines/wcus-r2-gen6"],"configurationItems":["wcus-r2-gen2","wcus-r2-gen3","wcus-r2-gen4","wcus-r2-gen5","wcus-r2-gen6"],"originAlertId":"3f2d4487-b0fc-4125-8bd5-7ad17384221e_PipeLineAlertRG_microsoft.insights_metricAlerts_WCUS-R2-Gen2_-117781227","firedDateTime":"2019-03-22T13:58:24.3713213Z","resolvedDateTime":"2019-03-22T14:03:16.2246313Z","description":"","essentialsVersion":"1.0","alertContextVersion":"1.0"}}}`
// 	got := AlertToNotification(alert.Parse(js)).Json()
// 	expected := `{"blocks":[{"type":"section","text":{"type":"mrkdwn","text":"*Azure monitor alert* :rotating_light:"}},{"type":"divider"},{"type":"actions","elements":[{"type":"button","text":{"type":"plain_text","text":"View alert in Azure Monitor","emoji":true},"url":"https://portal.azure.com/#blade/Microsoft_Azure_Monitoring/AlertDetailsTemplateBlade/alertId/%2Fsubscriptions%2F9876%2Fproviders%2FMicrosoft.AlertsManagement%2Falerts%2Fb9569717-bc32-442f-add5-83a997729330","style":"primary"}]},{"type":"header","text":{"type":"plain_text","text":"Summary","emoji":true}},{"type":"divider"},{"type":"section","fields":[{"type":"mrkdwn","text":"Fired (UTC)"},{"type":"plain_text","text":"2019-03-22T13:58:24.3713213Z","emoji":true},{"type":"mrkdwn","text":"Alert name"},{"type":"plain_text","text":"Test-Rule-1","emoji":true},{"type":"mrkdwn","text":"Severity"},{"type":"plain_text","text":":severity-critical: 0 - Critical","emoji":true}]},{"type":"header","text":{"type":"plain_text","text":"Resources","emoji":true}},{"type":"divider"},{"type":"section","fields":[{"type":"mrkdwn","text":"\u003chttps://portal.azure.com/#@nav.no/resource/subscriptions/1234/resourcegroups/pipelinealertrg/providers/microsoft.compute/virtualmachines/wcus-r2-gen2|wcus-r2-gen2\u003e"},{"type":"plain_text","text":":link:","emoji":true},{"type":"mrkdwn","text":"\u003chttps://portal.azure.com/#@nav.no/resource/subscriptions/1234/resourcegroups/pipelinealertrg/providers/microsoft.compute/virtualmachines/wcus-r2-gen3|wcus-r2-gen3\u003e"},{"type":"plain_text","text":":link:","emoji":true},{"type":"mrkdwn","text":"\u003chttps://portal.azure.com/#@nav.no/resource/subscriptions/1234/resourcegroups/pipelinealertrg/providers/microsoft.compute/virtualmachines/wcus-r2-gen4|wcus-r2-gen4\u003e"},{"type":"plain_text","text":":link:","emoji":true},{"type":"mrkdwn","text":"\u003chttps://portal.azure.com/#@nav.no/resource/subscriptions/1234/resourcegroups/pipelinealertrg/providers/microsoft.compute/virtualmachines/wcus-r2-gen5|wcus-r2-gen5\u003e"},{"type":"plain_text","text":":link:","emoji":true},{"type":"mrkdwn","text":"\u003chttps://portal.azure.com/#@nav.no/resource/subscriptions/1234/resourcegroups/pipelinealertrg/providers/microsoft.compute/virtualmachines/wcus-r2-gen6|wcus-r2-gen6\u003e"},{"type":"plain_text","text":":link:","emoji":true}]}]}`

// 	if string(got) != expected {
// 		t.Errorf("%s", "transform failure")
// 	}
// }
