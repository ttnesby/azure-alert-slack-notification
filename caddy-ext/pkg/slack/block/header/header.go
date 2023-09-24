package header

import (
	"github.com/ttnesby/slack-block-builder/pkg/slack/object/text"
)

type HeaderType string

const (
	CHeader HeaderType = "header"
)

// https://api.slack.com/reference/block-kit/blocks#header

type Header struct {
	Type HeaderType `json:"type"`
	Text *text.Text `json:"text"` // only plain_text and max 150 chars
}

func New(title string) *Header {

	return &Header{
		Type: CHeader,
		Text: text.NewPlain(title).FirstN(150),
	}
}
