package section

import (
	"encoding/json"

	"github.com/ttnesby/azure-alert-slack-notification/caddy-ext/pkg/slack/object/text"
)

type SectionType string

const (
	CSection SectionType = "section"
)

// https://api.slack.com/reference/block-kit/blocks#section

type Section struct {
	Type   SectionType  `json:"type"`
	Text   *text.Text   `json:"text,omitempty"`   // max 3000 chars
	Fields []*text.Text `json:"fields,omitempty"` // max 10 items, max 2000 chars each
	//accessory - not implemented
}

func New() *Section {
	return &Section{Type: CSection}
}

func (s *Section) SetText(txt *text.Text) *Section {
	s.Text = txt
	return s
}

func (s *Section) SetFields(fields ...*text.Text) *Section {

	lessThan10 := func() []*text.Text {
		switch noOfFields := len(fields); {
		case noOfFields > 10:
			return fields[:10]
		default:
			return fields
		}
	}

	lessThan2000 := func(f []*text.Text) []*text.Text {
		var corrected []*text.Text
		for _, t := range f {
			if t != nil {
				corrected = append(corrected, t.FirstN(2000))
			}
		}
		return corrected
	}

	s.Fields = lessThan2000(lessThan10())

	return s
}

func (s *Section) Json() string {
	js, err := json.Marshal(s)
	if err != nil {
		return "{}"
	}

	return string(js)
}
