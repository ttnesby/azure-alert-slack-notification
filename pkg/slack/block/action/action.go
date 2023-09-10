package action

import "github.com/ttnesby/slack-block-builder/pkg/slack/object/button"

// https://api.slack.com/reference/block-kit/blocks#actions

type Action struct {
	Type     string           `json:"type"`
	Elements []*button.Button `json:"elements"` // max 25 items
}

func New(button ...*button.Button) *Action {

	elementsFirstN := func(n int) []*button.Button {
		if len(button) > n {
			return button[:n]
		} else {
			return button
		}
	}

	return &Action{
		Type:     "actions",
		Elements: elementsFirstN(25),
	}
}
