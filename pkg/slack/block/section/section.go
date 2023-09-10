package section

import (
	"github.com/ttnesby/slack-block-builder/pkg/slack/object/text"
)

// https://api.slack.com/reference/block-kit/blocks#section

type Section[T text.Text[T]] struct {
	Type   string `json:"type"`
	Text   *T     `json:"text"`   // max 3000 chars
	Fields []*T   `json:"fields"` // max 10 items, max 2000 chars each
	//accessory - not implemented
}

func NewText[T text.Text[T]](text *T) *Section[T] {
	return &Section[T]{
		Type: "section",
		Text: text.FirstN(3000),
	}
}

func NewFields[T text.Text[T]](text ...*T) *Section[T] {

	fields := func() []*T {
		if len(text) > 10 {
			return text[:10]
		} else {
			return text
		}
	}()

	for i, t := range fields {
		if (*t).Len() > 2000 {
			fields[i] = (*t).FirstN(2000)
		}
	}

	return &Section[T]{
		Type:   "section",
		Fields: fields,
	}
}
