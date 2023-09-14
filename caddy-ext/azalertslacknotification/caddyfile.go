package azalertslacknotification

import (
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	"github.com/caddyserver/caddy/v2/caddyconfig/httpcaddyfile"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
)

func init() {
	httpcaddyfile.RegisterHandlerDirective("az_alert_slack_notification", parseCaddyfile)
}

func parseCaddyfile(h httpcaddyfile.Helper) (caddyhttp.MiddlewareHandler, error) {

	an := new(AzAlertSlackNotif)
	err := an.UnmarshalCaddyfile(h.Dispenser)
	return an, err

	// rb := new(AzAlertSlackNotif)

	// for h.Next() {
	// 	// configuration should be in a block
	// 	for h.NextBlock(0) {
	// 		switch h.Val() {
	// 		case "max_size":
	// 			var sizeStr string
	// 			if !h.AllArgs(&sizeStr) {
	// 				return nil, h.ArgErr()
	// 			}
	// 			size, err := humanize.ParseBytes(sizeStr)
	// 			if err != nil {
	// 				return nil, h.Errf("parsing max_size: %v", err)
	// 			}
	// 			rb.MaxSize = int64(size)
	// 		default:
	// 			return nil, h.Errf("unrecognized servers option '%s'", h.Val())
	// 		}
	// 	}
	// }

	// return rb, nil
}

func (m *AzAlertSlackNotif) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	// for d.Next() {
	// 	if !d.Args(&m.Prefix) {
	// 		return d.ArgErr()
	// 	}
	// }
	return nil
}

var (
	_ caddyfile.Unmarshaler = (*AzAlertSlackNotif)(nil)
)
