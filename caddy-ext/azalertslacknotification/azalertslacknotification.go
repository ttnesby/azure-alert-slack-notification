package azalertslacknotification

import (
	"bytes"
	"fmt"
	"io"
	"mime"
	"net/http"

	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
	"github.com/ttnesby/slack-block-builder/pkg/azure/alert"
	"github.com/ttnesby/slack-block-builder/pkg/transform"
	"go.uber.org/zap"
)

const (
	ErrorNilReqBody           = "nil request or body"
	ErrorParseMediaType       = "couldn't parse media type"
	ErrorUnsupportedMediaType = "unsupported media type"
	ErrorGetBody              = "couldn't get body"
	ErrorParseBody            = "couldn't parse body"
	ErrorUnsupportedSchemaId  = "unsupported schema id"
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

	if err := an.transformedBody(r); err != nil {
		an.logger.Debug("no transformation")
		return err
	} else {
		an.logger.Debug("after transformation", zap.Object("request", caddyhttp.LoggableHTTPRequest{Request: r}))
	}

	return next.ServeHTTP(w, r)
}

func (an AzAlertSlackNotif) transformedBody(r *http.Request) error {

	if r == nil || r.Body == nil {
		an.logger.Warn(ErrorNilReqBody)
		return fmt.Errorf(ErrorNilReqBody)
	}

	mType, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))

	if err != nil {
		an.logger.Error(ErrorParseMediaType, zap.Error(err))
		return err
	}

	if !alert.ContentTypeSupported(mType) {
		an.logger.Warn(ErrorUnsupportedMediaType, zap.String("Content-Type", mType))
		return fmt.Errorf(ErrorUnsupportedMediaType+": %s", mType)
	}

	doTransform := func() (io.ReadCloser, int, error) {

		existingBodyBuf := new(bytes.Buffer)
		if _, err := io.Copy(existingBodyBuf, r.Body); err != nil {
			an.logger.Warn(ErrorGetBody, zap.Error(err))
			return nil, 0, err
		}

		an.logger.Debug("existing body", zap.String("body", existingBodyBuf.String()))

		alert, err := alert.Parse(existingBodyBuf.String())

		if err != nil {
			an.logger.Warn(ErrorParseBody, zap.Error(err))
			return nil, 0, err
		}

		if !alert.SchemaIdSupported() {
			an.logger.Warn(ErrorUnsupportedSchemaId, zap.String("SchemaId", alert.SchemaId))
			return nil, 0, fmt.Errorf(ErrorUnsupportedSchemaId+": %s", alert.SchemaId)
		}

		notificationByte := transform.AlertToNotification(alert).Json()

		an.logger.Debug("transformed body", zap.String("body", string(notificationByte)))
		notificationBuf := bytes.NewBuffer(notificationByte)

		return io.NopCloser(notificationBuf), notificationBuf.Len(), nil
	}

	// normally net/http will close the body for us, but since we
	// are replacing it with a transformed one, we have to ensure we close
	// the real body ourselves when we're done
	defer r.Body.Close()

	if readCloser, length, err := doTransform(); err != nil {
		return err
	} else {
		r.Header.Set("Content-Length", fmt.Sprintf("%d", length))
		r.ContentLength = int64(length)
		r.Body = readCloser
		r.GetBody = func() (io.ReadCloser, error) {
			return readCloser, nil
		}
	}

	return nil
}

var (
	_ caddy.Provisioner           = (*AzAlertSlackNotif)(nil)
	_ caddy.Validator             = (*AzAlertSlackNotif)(nil)
	_ caddyhttp.MiddlewareHandler = (*AzAlertSlackNotif)(nil)
)
