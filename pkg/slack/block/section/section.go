package section

import (
	"github.com/ttnesby/slack-block-builder/pkg/slack/object/text"
)

// https://api.slack.com/reference/block-kit/blocks#section

type Section struct {
	Type   string       `json:"type"`
	Text   *text.Text   `json:"text"`
	Fields []*text.Text `json:"fields"`
}
