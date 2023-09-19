
export def alert-t [] {
    '{"schemaId":"azureMonitorCommonAlertSchema","data":{"essentials":{"alertId":"/subscriptions/9876/providers/Microsoft.AlertsManagement/alerts/b9569717-bc32-442f-add5-83a997729330","alertRule":"Test-Rule-1","severity":"Sev4","signalType":"Metric","monitorCondition":"Fired","monitoringService":"Platform","alertTargetIDs":["/subscriptions/1234/resourcegroups/pipelinealertrg/providers/microsoft.compute/virtualmachines/wcus-r2-gen2"],"configurationItems":["wcus-r2-gen2"],"originAlertId":"3f2d4487-b0fc-4125-8bd5-7ad17384221e_PipeLineAlertRG_microsoft.insights_metricAlerts_WCUS-R2-Gen2_-117781227","firedDateTime":"2019-03-22T13:58:24.3713213Z","resolvedDateTime":"2019-03-22T14:03:16.2246313Z","description":"","essentialsVersion":"1.0","alertContextVersion":"1.0"}}}' |
    curl --header "Content-Type: application/json" --include --data $'($in)'  http://localhost/transform/test
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
    gh release create ($in) --notes "wip"
    $ver | b-ca
}

export def b-ca [] {
    let ext1 = $"github.com/ttnesby/slack-block-builder/caddy-ext/azalertslacknotification@($in)" 
    ~/go/bin/xcaddy build --with ($ext1)
}
export def u-ca [] {
    ./caddy start --watch Caddyfile
}
export def d-ca [] {
    ./caddy stop
}