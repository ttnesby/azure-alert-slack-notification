
export def slack-alert-test [] {
    curl --header "Content-Type: application/json" --data $'($in)'  http://localhost/test
}

export def slack-alert-prod [] {
    curl --header "Content-Type: application/json" --data $'($in)'  http://localhost/prod
}

export def-env slack-caddy-setup [] {
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