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

- `true | env-setup`, load environment dependencies, see `Caddyfile`
- `false | env-setup`, remove environment dependencies
- `'v0.1.19' | r-ca`, create new release and build new version of caddy
- `u-ca`, start caddy with Caddyfile
- `d-ca`, stop caddy

```nushell
# when entering repo root
use ./utilities.nu *

# activate env vars
true | env-setup

# start caddy
u-ca

# fire off alert to slack test
alert-t

# stop caddy
d-ca

# do some coding and create new release, build new caddy
'v0.1.19' | r-ca

# remove env vars
false | env-setup
```

