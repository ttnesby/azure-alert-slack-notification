

export def notif-t [] {
    '{"blocks":[{"type":"divider"},{"type":"header","text":{"type":"plain_text","text":"New monitor alert!","emoji":true}},{"type":"divider"},{"type":"actions","elements":[{"type":"button","text":{"type":"plain_text","text":"View alert in Monitor","emoji":true},"url":"https://portal.azure.com/#blade/Microsoft_Azure_Monitoring/AlertDetailsTemplateBlade/alertId/%2Fsubscriptions%2F9876%2Fproviders%2FMicrosoft.AlertsManagement%2Falerts%2Fb9569717-bc32-442f-add5-83a997729330","style":"primary"}]},{"type":"header","text":{"type":"plain_text","text":"Summary","emoji":true}},{"type":"divider"},{"type":"section","fields":[{"type":"mrkdwn","text":"Fired (UTC)"},{"type":"plain_text","text":"2019-03-22T13:58:24.3713213Z","emoji":true},{"type":"mrkdwn","text":"Alert name"},{"type":"plain_text","text":"Test-Rule-1","emoji":true},{"type":"mrkdwn","text":"Severity"},{"type":"plain_text","text":":speech_balloon:  Verbose(4)","emoji":true}]},{"type":"header","text":{"type":"plain_text","text":"Resources","emoji":true}},{"type":"divider"},{"type":"section","fields":[{"type":"mrkdwn","text":"\u003chttps://portal.azure.com/#@nav.no/resource/subscriptions/1234/resourcegroups/pipelinealertrg/providers/microsoft.compute/virtualmachines/wcus-r2-gen2|wcus-r2-gen2\u003e"},{"type":"plain_text","text":":link:","emoji":true}]}]}' |
    curl --header "Content-Type: application/json" --include --data $'($in)'  http://localhost:80/direct/test
}

export def alert-fixed-t [] {
    '{"schemaId":"azureMonitorCommonAlertSchema","data":{"essentials":{"alertId":"/subscriptions/9876/providers/Microsoft.AlertsManagement/alerts/b9569717-bc32-442f-add5-83a997729330","alertRule":"Test-Rule-1","severity":"Sev4","signalType":"Metric","monitorCondition":"Fired","monitoringService":"Platform","alertTargetIDs":["/subscriptions/1234/resourcegroups/pipelinealertrg/providers/microsoft.compute/virtualmachines/wcus-r2-gen2"],"configurationItems":["wcus-r2-gen2"],"originAlertId":"3f2d4487-b0fc-4125-8bd5-7ad17384221e_PipeLineAlertRG_microsoft.insights_metricAlerts_WCUS-R2-Gen2_-117781227","firedDateTime":"2019-03-22T13:58:24.3713213Z","resolvedDateTime":"2019-03-22T14:03:16.2246313Z","description":"","essentialsVersion":"1.0","alertContextVersion":"1.0aaaaaaaaabbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbaaaaaaaaabbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbaaaaaaaaabbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbaaaaaaaaabbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbaaaaaaaaabbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbaaaaaaaaabbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbaaaaaaaaabbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbaaaaaaaaabbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbaaaaaaaaabbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbaaaaaaaaabbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbaaaaaaaaabbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbaaaaaaaaabbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbaaaaaaaaabbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbaaaaaaaaabbbbbbbbbbbbbbbbbbb"}}}' |
    curl --header "Content-Type: application/json" --include --data $'($in)'  http://localhost:80/transform/test
}

export def alert-t [] {
    '{"schemaId":"azureMonitorCommonAlertSchema","data":{"essentials":{"alertId":"/subscriptions/9876/providers/Microsoft.AlertsManagement/alerts/b9569717-bc32-442f-add5-83a997729330","alertRule":"Test-Rule-1","severity":"Sev4","signalType":"Metric","monitorCondition":"Fired","monitoringService":"Platform","alertTargetIDs":["/subscriptions/1234/resourcegroups/pipelinealertrg/providers/microsoft.compute/virtualmachines/wcus-r2-gen2"],"configurationItems":["wcus-r2-gen2"],"originAlertId":"3f2d4487-b0fc-4125-8bd5-7ad17384221e_PipeLineAlertRG_microsoft.insights_metricAlerts_WCUS-R2-Gen2_-117781227","firedDateTime":"2019-03-22T13:58:24.3713213Z","resolvedDateTime":"2019-03-22T14:03:16.2246313Z","description":"","essentialsVersion":"1.0","alertContextVersion":"1.0"}}}' |
    curl --header "Content-Type: application/json" --include --data $'($in)'  http://localhost:80/transform/test
}

export def alert-p [] {
    curl --header "Content-Type: application/json" --data $'($in)'  http://localhost/transform/prod
}

export def-env env-setup [] {
    if $in { 
        load-env {
            $"(op read op://Development/SlackTestNotification/CREDENTIAL/env_var)":$"(op read op://Development/SlackTestNotification/CREDENTIAL/secret_path)",
            $"(op read op://Development/SlackProdNotification/CREDENTIAL/env_var)":$"(op read op://Development/SlackProdNotification/CREDENTIAL/secret_path)",
        }
    } else {
        hide-env $"(op read op://Development/SlackTestNotification/CREDENTIAL/env_var)"
        hide-env $"(op read op://Development/SlackProdNotification/CREDENTIAL/env_var)"
    }
}

export def r-ca [] {
    let ver = $in
    gh release create ($ver) --notes "wip"
    b-ca $ver
}

export def b-ca [ver: string] {
    let ext1 = $"github.com/ttnesby/slack-block-builder/caddy-ext/azalertslacknotification@($ver)" 
    ~/go/bin/xcaddy build --with ($ext1)
}
export def u-ca [] {
    ./caddy start Caddyfile
}
export def d-ca [] {
    ./caddy stop
}