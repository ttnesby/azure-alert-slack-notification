package action

import "github.com/ttnesby/azure-alert-slack-notification/caddy-ext/pkg/slack/object/button"

// https://api.slack.com/reference/block-kit/blocks#actions

type ActionType string

const (
	CAction ActionType = "actions"
)

type Action struct {
	Type     ActionType       `json:"type"`
	Elements []*button.Button `json:"elements"` // max 25 items
}

func New(b ...*button.Button) *Action {

	elementsFirstN := func(n int) []*button.Button {
		if len(b) > n {
			return b[:n]
		} else {
			return b
		}
	}

	return &Action{
		Type:     CAction,
		Elements: elementsFirstN(25),
	}
}
