package azalertslacknotification

// see https://github.com/caddyserver/caddy/blob/6f0f159ba56adeb6e2cbbb408651419b87f20856/modules/caddyhttp/replacer.go
// see https://github.com/RussellLuo/caddy-ext/blob/master/requestbodyvar/requestbodyvar.go
// see https://github.com/caddyserver/caddy/blob/master/modules/caddyhttp/requestbody/requestbody.go

import (
	"bytes"
	"io"
	"net/http"

	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
	"github.com/ttnesby/slack-block-builder/pkg/azure/alert"
	"github.com/ttnesby/slack-block-builder/pkg/transform"
	"go.uber.org/zap"
)

const (
	BodyCtxKey       caddy.CtxKey = "body"
	bodyBufferCtxKey caddy.CtxKey = "body_buffer"
)

func init() {
	caddy.RegisterModule(AzAlertSlackNotif{})
}

// AzAlertSlackNotif is middleware that transforms an Azure genric alert
// to a Slack message of a hard coded format - see transform.AlertToNotification

// of the URI matches a given prefix.
type AzAlertSlackNotif struct {
	//Prefix string `json:"prefix,omitempty"`
	logger *zap.Logger
}

// CaddyModule returns the Caddy module information.
func (AzAlertSlackNotif) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "http.handlers.az_alert_slack_notification",
		New: func() caddy.Module { return new(AzAlertSlackNotif) },
	}
}

// Provision a Zap logger to AzAlertSlackNotif.
func (an *AzAlertSlackNotif) Provision(ctx caddy.Context) error {
	an.logger = ctx.Logger(an)

	an.logger.Info("logger activated")
	return nil
}

// Validate the prefix from the module's configuration, setting the
// default prefix "." if necessary.

// func (an *AzAlertSlackNotif) Validate() error {
// 	if an.Prefix == "" {
// 		an.Prefix = "."
// 	}
// 	return nil
// }

// ServeHTTP implements the caddyhttp.MiddlewareHandler interface.
func (an AzAlertSlackNotif) ServeHTTP(w http.ResponseWriter, r *http.Request, next caddyhttp.Handler) error {

	repl := r.Context().Value(caddy.ReplacerCtxKey).(*caddy.Replacer)

	logger := an.logger.With(zap.Object("request", caddyhttp.LoggableHTTPRequest{Request: r}))

	an.TransformBody(r, repl)

	logger.Info("request", zap.Object("request", caddyhttp.LoggableHTTPRequest{Request: r}))

	return next.ServeHTTP(w, r)
}

func (an AzAlertSlackNotif) TransformBody(r *http.Request, repl *caddy.Replacer) {

	if r == nil || r.Body == nil {
		return
	}

	doTransform := func() (io.ReadCloser, error) {
		buf := new(bytes.Buffer)
		_, _ = io.Copy(buf, r.Body) // cannot do reasonable error handling

		r := transform.AlertToNotification(alert.Parse(buf.String())).Json()
		return io.NopCloser(bytes.NewBuffer(r)), nil
	}

	// normally net/http will close the body for us, but since we
	// are replacing it with a transformed one, we have to ensure we close
	// the real body ourselves when we're done
	defer r.Body.Close()

	r.Body, _ = doTransform()
	r.GetBody = func() (io.ReadCloser, error) {
		return doTransform()
	}
}

var (
	_ caddy.Provisioner = (*AzAlertSlackNotif)(nil)
	//_ caddy.Validator             = (*AzAlertSlackNotif)(nil)
	_ caddyhttp.MiddlewareHandler = (*AzAlertSlackNotif)(nil)
)
