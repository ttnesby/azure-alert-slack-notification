package azalertslacknotification

import (
	"bytes"
	"fmt"
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

	an.logger.Debug("logger activated")
	return nil
}

func (an *AzAlertSlackNotif) Validate() error {
	// nothing to validate
	return nil
}

// ServeHTTP implements the caddyhttp.MiddlewareHandler interface.
func (an AzAlertSlackNotif) ServeHTTP(w http.ResponseWriter, r *http.Request, next caddyhttp.Handler) error {

	an.logger.Debug("before transformation", zap.Object("request", caddyhttp.LoggableHTTPRequest{Request: r}))

	an.TransformBody(r)

	an.logger.Debug("after transformation", zap.Object("request", caddyhttp.LoggableHTTPRequest{Request: r}))

	return next.ServeHTTP(w, r)
}

func (an AzAlertSlackNotif) TransformBody(r *http.Request) {

	if r == nil || r.Body == nil {
		return
	}

	// verify Content-Type application/json

	doTransform := func() (io.ReadCloser, int, error) {
		buf := new(bytes.Buffer)
		_, _ = io.Copy(buf, r.Body) // reasonable error handling, not yet...

		an.logger.Debug("before body", zap.String("body", buf.String()))

		// verify "schemaId":"azureMonitorCommonAlertSchema"
		r := transform.AlertToNotification(alert.Parse(buf.String())).Json()

		an.logger.Debug("transformed body", zap.String("body", string(r)))
		transformedBuf := bytes.NewBuffer(r)

		return io.NopCloser(transformedBuf), transformedBuf.Len(), nil
	}

	// normally net/http will close the body for us, but since we
	// are replacing it with a transformed one, we have to ensure we close
	// the real body ourselves when we're done
	defer r.Body.Close()

	readCloser, length, err := doTransform()

	r.Header.Set("Content-Length", fmt.Sprintf("%d", length))
	r.ContentLength = int64(length)
	r.Body = readCloser
	r.GetBody = func() (io.ReadCloser, error) {
		return readCloser, err
	}
}

var (
	_ caddy.Provisioner           = (*AzAlertSlackNotif)(nil)
	_ caddy.Validator             = (*AzAlertSlackNotif)(nil)
	_ caddyhttp.MiddlewareHandler = (*AzAlertSlackNotif)(nil)
)
