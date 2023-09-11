package action

import "github.com/ttnesby/slack-block-builder/pkg/slack/object/button"

// https://api.slack.com/reference/block-kit/blocks#actions

const (
	TypeAction = "actions"
)

type Action struct {
	Type     string           `json:"type"`
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
		Type:     TypeAction,
		Elements: elementsFirstN(25),
	}
}
