package transform

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/ttnesby/azure-alert-slack-notification/caddy-ext/pkg/azure/alert"
	"github.com/ttnesby/azure-alert-slack-notification/caddy-ext/pkg/slack/block/action"
	"github.com/ttnesby/azure-alert-slack-notification/caddy-ext/pkg/slack/block/divider"
	"github.com/ttnesby/azure-alert-slack-notification/caddy-ext/pkg/slack/block/header"
	"github.com/ttnesby/azure-alert-slack-notification/caddy-ext/pkg/slack/block/section"
	"github.com/ttnesby/azure-alert-slack-notification/caddy-ext/pkg/slack/notification"
	"github.com/ttnesby/azure-alert-slack-notification/caddy-ext/pkg/slack/object/button"
	"github.com/ttnesby/azure-alert-slack-notification/caddy-ext/pkg/slack/object/text"
)

func severity(a *alert.CommonAlertSchema) notification.Severity {
	switch a.Data.Essentials.Severity {
	case "Sev0":
		return notification.SeverityCritical
	case "Sev1":
		return notification.SeverityError
	case "Sev2":
		return notification.SeverityWarning
	case "Sev3":
		return notification.SeverityInformation
	case "Sev4":
		return notification.SeverityVerbose
	case "TestStart":
		return notification.SevTestStart
	case "TestEnd":
		return notification.SevTestEnd
	default:
		return notification.SeverityUnknown
	}
}

func AlertToNotification(a *alert.CommonAlertSchema) *notification.Notification {

	alertUrl := alert.UrlAlertBlade + url.QueryEscape(a.Data.Essentials.AlertId)

	payload := notification.New().
		AddDivider(divider.New()).
		AddHeader(header.New("New monitor alert!")).
		AddDivider(divider.New()).
		AddAction(action.New(button.New("View alert in Monitor", alertUrl))).
		AddHeader(header.New("Summary")).
		AddDivider(divider.New()).
		AddSection(
			section.New().
				SetFields(
					text.NewMarkDown("Fired (UTC)"), text.NewPlain(a.Data.Essentials.FiredDateTime),
					text.NewMarkDown("Alert name"), text.NewPlain(a.Data.Essentials.AlertRule),
					text.NewMarkDown("Severity"), text.NewPlain(string(severity(a))),
				),
		)

	resourceName := func(path string) string {
		parts := strings.Split(path, "/")
		return parts[len(parts)-1]
	}

	resourceUrl := func(path string) string {
		return alert.UrlResourceBlade + path
	}

	var resources []*text.Text
	for _, r := range a.Data.Essentials.AlertTargetIDs {
		resources = append(resources, text.NewMarkDown(fmt.Sprintf("<%s|%s>", resourceUrl(r), resourceName(r))))
		resources = append(resources, text.NewPlain(string(notification.IconLink)))
	}

	if len(resources) > 0 {
		payload.
			AddHeader(header.New("Resources")).
			AddDivider(divider.New()).
			AddSection(section.New().SetFields(resources...))
	}

	return payload
}
