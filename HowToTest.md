# How to test

## Prerequisites

- [xcaddy](https://github.com/caddyserver/xcaddy) available, e.g. `go install github.com/caddyserver/xcaddy/cmd/xcaddy@latest`
- [Nushell](https://www.nushell.sh/) is active
- A password security CLI is installed, [e.g. 1Password CLI](https://developer.1password.com/docs/cli/get-started/)

## `utilities.nu`

The `./utilities.nu` defines basic utilities.

```nushell
# activate utilities
use ./utilities.nu *
```

- `e-setup [false]`, load environment dependencies, see `Caddyfile`. Default `true`, remove environment dependencies with `false`

- `r-ca vMa.Mi.Path... [branch]`, create new release and build new version of caddy. Default branch is `main`
- `b-ca vMa.Mi.Path...`, build new version of caddy

- `u-ca`, start caddy with Caddyfile
- `d-ca`, stop caddy

- `t-alert`, send test alert to caddy, received in slack#teste_varsel
- `te-alert`, send errorneous requests

- `p-alert`, send test alert to caddy, received in slack#azure-platform-alerts
- `h-status`, ask for health status

```nushell
# when entering repo root
use ./utilities.nu *

# activate env vars
e-setup

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

# stop caddy
d-ca

# do some coding and create new release, build new caddy
r-ca vx.y.z... [branch]

# goto :iteration

# remove env vars
e-setup false
```
