package main

import (
	"fmt"
	"github.com/ttnesby/slack-block-builder/pkg/slack/block/action"
	"github.com/ttnesby/slack-block-builder/pkg/slack/block/divider"
	"github.com/ttnesby/slack-block-builder/pkg/slack/block/header"
	"github.com/ttnesby/slack-block-builder/pkg/slack/block/section"
	"github.com/ttnesby/slack-block-builder/pkg/slack/notification"
	"github.com/ttnesby/slack-block-builder/pkg/slack/object/button"
	"github.com/ttnesby/slack-block-builder/pkg/slack/object/text"
)

func main() {
	json := notification.New().
		AddSection(
			section.New().
				SetText(
					text.NewMarkDown(fmt.Sprintf("*Azure monitor alert* %s", notification.IconRotatingLight)),
				),
		).
		AddDivider(divider.New()).
		AddAction(action.New(button.New("View alert in Azure Monitor", "https://www.vg.no"))).
		AddHeader(header.New("Summary")).
		AddDivider(divider.New()).
		AddSection(
			section.New().
				SetFields(
					text.NewMarkDown("Fired (UTC)"), text.NewPlain("2023/09/11 16:43"),
					text.NewMarkDown("Alert name"), text.NewPlain("check something..."),
					text.NewMarkDown("Severity"), text.NewPlain(notification.SeverityInformation),
					// adding a resource example
					text.NewMarkDown("<https://www.aftenposten.no|newspaper>"), text.NewPlain(notification.IconLink),
				),
		).
		Json()

	fmt.Printf("%s", json)
}
