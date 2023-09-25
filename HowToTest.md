# How to test

## Optional prerequisites

- [Nushell](https://www.nushell.sh/) is active
- A secret manager CLI is installed, [e.g. 1Password CLI](https://developer.1password.com/docs/cli/get-started/)

It's fine without nushell and secret manager, just more manual labor. See custom commands for details.

## `utilities.nu`

The `./utilities.nu` defines basic custom commands.

```nushell
# activate utilities
use ./utilities.nu *
```

- `r-ca <version> [branch]`, create new release and build custom caddy. Default branch is `main`
- `cb-ca`, build custom caddy with latest on current architecture
- `mb-ca`, build custom caddy with latest for multi-architecture

- `u-ca`, start caddy with `./Caddyfile`
- `d-ca`, stop caddy

- `t-alert`, send test alert of different severities to caddy, received in slack#teste_varsel
- `te-alert`, send errorneous requests, triggering WAF and error handling in custom extension

- `p-alert`, send test alert to caddy, received in slack#azure-platform-alerts
- `h-status`, get health status
- `tr-health <noOfRequests>`, send noOfRequest to health api, trigger rate limiter

```nushell
# when entering repo root
use ./utilities.nu *

# :iteration
# start caddy
u-ca

# check health status
h-status

# fire off alert to slack test
t-alert
te-alert

# fire off alert to slack prod
p-alert

# do some coding and create new release, build new caddy
r-ca <version> [branch]

# goto :iteration
```
