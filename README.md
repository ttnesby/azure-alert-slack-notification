# azure-alert-slack-notification

Is a prototype for an Azure container app receiving azure alerts, transforms them to slack notification, and inverse proxy them to relevant slack channel.

## Components

- [Caddy server](https://github.com/caddyserver/caddy)
- [Transformer as Caddy custom extension](https://caddyserver.com/docs/extending-caddy)
- [OWASP WAF as Caddy custom module](https://caddyserver.com/docs/modules/http.handlers.waf#github.com/corazawaf/coraza-caddy)
- [Rate limiter as Caddy custom module](https://caddyserver.com/docs/modules/http.handlers.rate_limit#github.com/mholt/caddy-ratelimit)

An example of the [caddy configuration ](./Caddyfile).

## Repo structure

The repo is multi-module with three different modules

- `caddy-ext`, transformer logic
- `cmd/caddy`, build caddy for current architecture
- `cmd/multibuilder`, build caddy for multi-architecture

## How to test

See [HowToTest](./HowToTest.md)
