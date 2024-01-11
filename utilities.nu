# different severities related to azure alert, except from unknown
const severities = ['Sev0' 'Sev1' 'Sev2' 'Sev3' 'Sev4' 'unknown']

# test json payload for different test scenarios
def alertJson [sev: string] {
    {
        "schemaId":"azureMonitorCommonAlertSchema",
        "data":{
            "essentials":{
                "alertId":"/subscriptions/9876/providers/Microsoft.AlertsManagement/alerts/b9569717-bc32-442f-add5-83a997729330",
                "alertRule":"Test-Rule-1",
                "severity":$"($sev)",
                "signalType":"Metric",
                "monitorCondition":"Fired",
                "monitoringService":"Platform",
                "alertTargetIDs":["/subscriptions/1234/resourcegroups/pipelinealertrg/providers/microsoft.compute/virtualmachines/wcus-r2-gen2"],
                "configurationItems":["wcus-r2-gen2"],
                "originAlertId":"3f2d4487-b0fc-4125-8bd5-7ad17384221e_PipeLineAlertRG_microsoft.insights_metricAlerts_WCUS-R2-Gen2_-117781227",
                "firedDateTime":"2019-03-22T13:58:24.3713213Z",
                "resolvedDateTime":"2019-03-22T14:03:16.2246313Z",
                "description":"",
                "essentialsVersion":"1.0",
                "alertContextVersion":"1.0"
            }
        }
    }
}

# caddy - generate slack notifications to #testevarsel, one for each severity, with test start and end
export def t-alert [] {
    let url = 'http://localhost/api/slack/testevarsel'
    let ct = 'application/json'

    let start = 'TestStart' | http post --content-type $ct  $url (alertJson $in) --full --allow-errors | reject headers
    let allSevs = $severities | par-each --keep-order {|sev| http post --content-type $ct $url (alertJson $sev) --full --allow-errors} | reject headers
    let end = 'TestEnd' | http post --content-type $ct $url (alertJson $in) --full --allow-errors | reject headers

    [$start $allSevs $end]
    | flatten
    | each {|r| $r.status}
}

# caddy - test rate limit of < 120 req/1min window against /api/health, should give (x - 120) "429 Too Many Requests"
export def tr-health [noOfReq: int] {
    if $noOfReq >= 1 {
        1..$noOfReq
        | par-each --keep-order {|| http get http://localhost/api/health --full --allow-errors}
        | filter {|el| $el.status == 429 }
        | reduce -f 0 {|_,acc| $acc + 1}
        | print $"No of 429 'Too Many Requests': ($in)"
    } 
}

# caddy - test different error situations, each genetating 403 Forbidden
export def te-alert [] {
    let url = "http://localhost/api/slack/testevarsel"

    let ctInv = 'invalidMediaType'
    let ctUns = 'text/xml'
    let ct = 'application/json'

    print "### test case: WAF invalid media type\n"
    '{}' | curl --header $"Content-Type: ($ctInv)" --include --data $'($in)' ($url)
    #print $"\nstatus (http post --content-type $ctInv $url 'æåø' --full --allow-errors | get status)\n"

    print "### test case: unsupported media type\n"
    'plain' | curl --header $"Content-Type: ($ctUns)" --include --data $'($in)' ($url)
    #print $"\nstatus (http post --content-type $ctUns $url 'plain' --full --allow-errors | get status)\n"

    print "### test case: cannot parse body\n"
    print $"\nstatus (http post --content-type $ct $url '' --full --allow-errors | get status)\n"

    print "### test case: unsupported schema id\n"
    print $"status (http post --content-type $ct $url {schemaId:unsupportedSchema} --full --allow-errors | get status )\n"
}

# caddy - generate slack notification to #azure-platform-alerts, severity verbose
export def p-alert [] {
    let url = 'http://localhost/api/slack/azureplatformalerts'
    let ct = 'application/json'

    'TestStart' | http post --content-type $ct  $url (alertJson $in) --full --allow-errors | reject headers
    'TestEnd' | http post --content-type $ct  $url (alertJson $in) --full --allow-errors | reject headers
}

# caddy - get health status of web server
export def h-status [] {
    http get http://localhost/api/health --full --allow-errors | reject headers
}

# set | unset required environment variables for Caddyfile, web hook secrets for related slack channels
def --env e-setup [set: bool = true] {
    print "\n### make required env. vars available\n"
    let secretStoreMap = {
        SLACK_TESTEVARSEL:['op://Development' SlackTesteVarsel 'CREDENTIAL/secret_path'],
        SLACK_AZUREPLATFORMALERTS:['op://Development' SlackAzurePlatformAlerts 'CREDENTIAL/secret_path']
        }
    let envVars = $secretStoreMap | items {|key,_| $key} | enumerate

    let statusEnvVars = $env | items {|key,_| $key } | filter {|e| $e =~ 'SLACK_*'} | enumerate
    let missing = if ($statusEnvVars | is-empty) {$envVars} else { $statusEnvVars | filter {|e| $e.item not-in $envVars.item}}
    let existing = $statusEnvVars | filter {|e| $e.item in $envVars.item}

    if $set {
        $missing.item
        | each {|v|
            let opPath = $secretStoreMap | transpose | filter {|e| $e.column0 == $v} | get column1 | first | path join
            let opSecret = $"(op read $opPath)"
            {$v:$opSecret}
        }
        | reduce -f {} {|e, acc| $acc | merge $e }
        | load-env
    } else {
        # using overlays...
        # how to remove env var in current scope??
        #$existing.item | each {|v| hide-env $v}
    }
}

# caddy - create a new release with default branch main
export def r-ca [ver: string, branch: string = "main"] {
    print "\n### do ext. tests\n"
    go test -cover ./caddy-ext/pkg/...

    print $"\n### create new relase of ext. - branch [($branch) version [($ver)]] \n"
    gh release create ($ver) --notes "wip" --target ($branch)

    cb-ca
    mb-ca
}

# caddy - build a new version and relevant extensions, on current architecture
export def cb-ca [] {
    print "\n### build custom caddy with latest of ext.\n"
    go build -o ./caddy ./cmd/caddy
}

# caddy - build multi-architecture and custom modules
export def mb-ca [] {
    print "\n### build multi architecture of custom caddy with latest of ext.\n"
    dagger run go run cmd/multibuilder/main.go
}

# caddy - start with local Caddyfile
export def u-ca [] {
    e-setup
    d-ca
    print "\n### fire up custom caddy\n"
    ./caddy start Caddyfile
}

# caddy - stop
export def d-ca [] {
    print "\n### if caddy is runing, take it(/them) down\n"
    ps | where name == caddy | get pid | each {|e| kill $e }
    #./caddy stop
}