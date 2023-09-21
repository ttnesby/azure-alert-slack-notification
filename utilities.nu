export def t-alert [] {
    '{"schemaId":"azureMonitorCommonAlertSchema","data":{"essentials":{"alertId":"/subscriptions/9876/providers/Microsoft.AlertsManagement/alerts/b9569717-bc32-442f-add5-83a997729330","alertRule":"Test-Rule-1","severity":"Sev1","signalType":"Metric","monitorCondition":"Fired","monitoringService":"Platform","alertTargetIDs":["/subscriptions/1234/resourcegroups/pipelinealertrg/providers/microsoft.compute/virtualmachines/wcus-r2-gen2"],"configurationItems":["wcus-r2-gen2"],"originAlertId":"3f2d4487-b0fc-4125-8bd5-7ad17384221e_PipeLineAlertRG_microsoft.insights_metricAlerts_WCUS-R2-Gen2_-117781227","firedDateTime":"2019-03-22T13:58:24.3713213Z","resolvedDateTime":"2019-03-22T14:03:16.2246313Z","description":"","essentialsVersion":"1.0","alertContextVersion":"1.0"}}}' |
    curl --header "Content-Type: application/json" --include --data $'($in)' http://localhost/api/slack/testevarsel
}

export def te-alert [] {
    let url = "http://localhost/api/slack/testevarsel"
    "empty body\n"
    '' | curl -X POST --include ($url)
    "\n\n"

    "wrong content type\n"
    '{}' | curl --header "Content-Type: application/xml" --include --data $'($in)' ($url)
    "\n\n"

    "wrong schema id\n"
    '{}' | curl --header "Content-Type: application/json" --include --data $'($in)' ($url)
    "\n\n"
}

export def p-alert [] {
    '{"schemaId":"azureMonitorCommonAlertSchema","data":{"essentials":{"alertId":"/subscriptions/9876/providers/Microsoft.AlertsManagement/alerts/b9569717-bc32-442f-add5-83a997729330","alertRule":"Test-Rule-1","severity":"Sev4","signalType":"Metric","monitorCondition":"Fired","monitoringService":"Platform","alertTargetIDs":["/subscriptions/1234/resourcegroups/pipelinealertrg/providers/microsoft.compute/virtualmachines/wcus-r2-gen2"],"configurationItems":["wcus-r2-gen2"],"originAlertId":"3f2d4487-b0fc-4125-8bd5-7ad17384221e_PipeLineAlertRG_microsoft.insights_metricAlerts_WCUS-R2-Gen2_-117781227","firedDateTime":"2019-03-22T13:58:24.3713213Z","resolvedDateTime":"2019-03-22T14:03:16.2246313Z","description":"","essentialsVersion":"1.0","alertContextVersion":"1.0"}}}' |
    curl --header "Content-Type: application/json" --include --data $'($in)' http://localhost/api/slack/azureplatformalerts
}

export def pe-alert [] {
    '{}' | curl --header "Content-Type: application/json" --include --data $'($in)' http://localhost/api/slack/azureplatformalerts
}

export def h-status [] {
    curl http://localhost/api/health
}

export def-env e-setup [set: bool = true] {
    if $set { 
        load-env {
            $"(op read op://Development/SlackTestNotification/CREDENTIAL/env_var)":$"(op read op://Development/SlackTestNotification/CREDENTIAL/secret_path)",
            $"(op read op://Development/SlackProdNotification/CREDENTIAL/env_var)":$"(op read op://Development/SlackProdNotification/CREDENTIAL/secret_path)",
        }
    } else {
        hide-env $"(op read op://Development/SlackTestNotification/CREDENTIAL/env_var)"
        hide-env $"(op read op://Development/SlackProdNotification/CREDENTIAL/env_var)"
    }
}

export def r-ca [ver: string, branch: string = "main"] {
    gh release create ($ver) --notes "wip" --target ($branch)
    b-ca $ver
}

export def b-ca [ver: string] {
    let ext1 = $"github.com/ttnesby/slack-block-builder/caddy-ext/azalertslacknotification@($ver)" 
    let ext2 = $"github.com/corazawaf/coraza-caddy/v2"  # waf
    let ext3 = $"github.com/mholt/caddy-ratelimit"      # rate limiter
    ~/go/bin/xcaddy build --with ($ext1) --with ($ext2) --with ($ext3)
}
export def u-ca [] {
    ./caddy start Caddyfile
}
export def d-ca [] {
    ./caddy stop
}