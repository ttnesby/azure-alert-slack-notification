# How to test

Assumptions

- Caddy available, e.g. `brew install caddy`
- [Nushell](https://www.nushell.sh/) is active
- A password security CLI is installed, [e.g. 1Password CLI](https://developer.1password.com/docs/cli/get-started/)

The `slack-notification-config.nu` defines the basic configuration

- set/unset a couple of environment variables used by `Caddyfile`
- set a couple of aliases for easy `curl` usage

Being in `repo root/caddy-test`;

```nushell
# load configuration
use ./slack-notification-config.nu *

# activate env vars
true | slack-caddy-setup

# start caddy
caddy start ./Caddyfile

# send test alert to test
go run ../main.go | slack-alert-test

# result should be ok and a related notification in #test_varsel

# unset env vars
false | slack-caddy-setup

```
