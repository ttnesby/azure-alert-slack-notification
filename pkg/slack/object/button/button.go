package button

import (
	"github.com/ttnesby/slack-block-builder/pkg/slack/object/text"
)

// https://api.slack.com/reference/block-kit/block-elements#button

const (
	TypeButton = "button"
)

type Button struct {
	Type  string     `json:"type"`
	Text  *text.Text `json:"text"`            // only plain_text allowed, max 75 chars
	Url   string     `json:"url"`             // max 3000 chars
	Style string     `json:"style,omitempty"` // primary, danger
	//value, confirm, accessibility_label - not implemented
}

func New(s string, url string) *Button {

	urlFirstN := func(n int) string {
		if len(url) > n {
			return url[:n]
		} else {
			return url
		}
	}

	return &Button{
		Type:  TypeButton,
		Text:  text.NewPlain(s).FirstN(75),
		Url:   urlFirstN(3000),
		Style: "primary",
	}
}
