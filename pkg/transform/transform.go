package transform

import (
	"fmt"
	"github.com/ttnesby/slack-block-builder/pkg/azure/alert"
	"github.com/ttnesby/slack-block-builder/pkg/slack/block/action"
	"github.com/ttnesby/slack-block-builder/pkg/slack/block/divider"
	"github.com/ttnesby/slack-block-builder/pkg/slack/block/header"
	"github.com/ttnesby/slack-block-builder/pkg/slack/block/section"
	"github.com/ttnesby/slack-block-builder/pkg/slack/notification"
	"github.com/ttnesby/slack-block-builder/pkg/slack/object/button"
	"github.com/ttnesby/slack-block-builder/pkg/slack/object/text"
	"net/url"
	"strings"
)

func severity(a *alert.CommonAlertSchema) string {
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
	default:
		return notification.SeverityUnknown
	}
}

func AlertToNotification(a *alert.CommonAlertSchema) string {

	alertUrl := alert.UrlAlertBlade + url.QueryEscape(a.Data.Essentials.AlertId)

	payload := notification.New().
		AddSection(
			section.New().
				SetText(
					text.NewMarkDown(fmt.Sprintf("*Azure monitor alert* %s", notification.IconRotatingLight)),
				),
		).
		AddDivider(divider.New()).
		AddAction(action.New(button.New("View alert in Azure Monitor", alertUrl))).
		AddHeader(header.New("Summary")).
		AddDivider(divider.New()).
		AddSection(
			section.New().
				SetFields(
					text.NewMarkDown("Fired (UTC)"), text.NewPlain(a.Data.Essentials.FiredDateTime),
					text.NewMarkDown("Alert name"), text.NewPlain(a.Data.Essentials.AlertRule),
					text.NewMarkDown("Severity"), text.NewPlain(severity(a)),
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
		resources = append(resources, text.NewPlain(notification.IconLink))
	}

	if len(resources) > 0 {
		payload.
			AddHeader(header.New("Resources")).
			AddDivider(divider.New()).
			AddSection(section.New().SetFields(resources...))
	}

	return payload.Json()
}
