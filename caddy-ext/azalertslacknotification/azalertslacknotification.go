package azalertslacknotification

// see https://github.com/caddyserver/caddy/blob/6f0f159ba56adeb6e2cbbb408651419b87f20856/modules/caddyhttp/replacer.go
// see https://github.com/RussellLuo/caddy-ext/blob/master/requestbodyvar/requestbodyvar.go
// see https://github.com/caddyserver/caddy/blob/master/modules/caddyhttp/requestbody/requestbody.go

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"strings"

	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
	"github.com/ttnesby/slack-block-builder/pkg/azure/alert"
	"github.com/ttnesby/slack-block-builder/pkg/transform"
	"go.uber.org/zap"
)

const (
	fullReqBodyReplPrefix               = "http.request.body."
	shortReqBodyReplPrefix              = "body."
	bodyBufferCtxKey       caddy.CtxKey = "body_buffer"
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
func (an AzAlertSlackNotif) ServeHTTP(w http.ResponseWriter, r *http.Request,
	next caddyhttp.Handler) error {

	an.logger.Info("received request")

	transformer := func(key string) (any, bool) {

		key, ok := parseKey(key)
		if !ok || key == "" {
			an.logger.Error("invalid var", zap.String("key", key))
			return nil, false
		}

		// First of all, try to get the value from the buffered JSON body, if any.
		buf, ok := r.Context().Value(bodyBufferCtxKey).(*bytes.Buffer)
		if ok {
			an.logger.Debug("got from the buffer", zap.String("key", key))
			return string(transform.AlertToNotification(alert.Parse(buf.String())).Json()), true
		}

		an.logger.Debug("got from the body", zap.String("key", key))

		// Otherwise, try to get the value by reading the request body.
		if r == nil || r.Body == nil {
			return "", true
		}
		// Close the real body since we will replace it with a fake one.
		defer r.Body.Close()

		// Copy the request body.
		buf = new(bytes.Buffer)
		if _, err := io.Copy(buf, r.Body); err != nil {
			return "", true
		}

		// Replace the real body with buffered data.
		r.Body = io.NopCloser(buf)

		// Add the buffered JSON body into the context for the request.
		ctx := context.WithValue(r.Context(), bodyBufferCtxKey, buf)
		r = r.WithContext(ctx)

		return "", true

	}

	// if r == nil || r.Body == nil {
	// 	return next.ServeHTTP(w, r)
	// }

	// // normally net/http will close the body for us, but since we
	// // are replacing it with a transformed one, we have to ensure we close
	// // the real body ourselves when we're done

	// defer r.Body.Close()

	// // read the request body into a buffer (can't pool because we
	// // don't know its lifetime and would have to make a copy anyway)

	// buf := new(bytes.Buffer)
	// if _, err := io.Copy(buf, r.Body); err != nil {
	// 	return err
	// }

	// an.logger.Info("received azure alert", zap.String("body", buf.String()))

	// //slackMsg := transform.AlertToNotification(alert.Parse(buf.String())).Json()
	// smBuf := bytes.NewBuffer(transform.AlertToNotification(alert.Parse(buf.String())).Json())

	// an.logger.Info("transformed to slack notification", zap.String("body", string(smBuf.Bytes())))

	// // must set content length before body https://github.com/caddyserver/caddy/issues/5485
	// // r.Header.Set("Content-Length", fmt.Sprint(smBuf.Cap()))
	// // an.logger.Info("new content length is set", zap.Int("Content-Length", smBuf.Cap()))

	// // replace real body with buffered data
	// r.Body = io.NopCloser(smBuf)

	// // Add the buffered JSON body into the context for the request.
	// ctx := context.WithValue(r.Context(), bodyBufferCtxKey, smBuf)
	// r = r.WithContext(ctx)
	// an.logger.Info("request context updated with new body")

	repl := r.Context().Value(caddy.ReplacerCtxKey).(*caddy.Replacer)
	repl.Map(transformer)

	return next.ServeHTTP(w, r)
}

func parseKey(s string) (string, bool) {
	switch {
	case strings.HasPrefix(s, fullReqBodyReplPrefix):
		return s[len(fullReqBodyReplPrefix):], true
	case strings.HasPrefix(s, shortReqBodyReplPrefix):
		return s[len(shortReqBodyReplPrefix):], true
	default:
		// unrecognized
		return "", false
	}
}

var (
	_ caddy.Provisioner = (*AzAlertSlackNotif)(nil)
	//_ caddy.Validator             = (*AzAlertSlackNotif)(nil)
	_ caddyhttp.MiddlewareHandler = (*AzAlertSlackNotif)(nil)
)
