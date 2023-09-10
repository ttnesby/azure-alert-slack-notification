package header

import (
	"github.com/ttnesby/slack-block-builder/pkg/slack/object/text"
)

// https://api.slack.com/reference/block-kit/blocks#header

type Header struct {
	Type string      `json:"type"`
	Text *text.Plain `json:"text"` // max 150 chars
}

func New(text *text.Plain) *Header {
	return &Header{
		Type: "header",
		Text: text.FirstN(150),
	}
}
