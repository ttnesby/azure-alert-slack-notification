package section

import (
	"encoding/json"
	"github.com/ttnesby/slack-block-builder/pkg/slack/object/text"
)

const (
	TypeSection = "section"
)

// https://api.slack.com/reference/block-kit/blocks#section

type Section struct {
	Type   string       `json:"type"`
	Text   *text.Text   `json:"text,omitempty"`   // max 3000 chars
	Fields []*text.Text `json:"fields,omitempty"` // max 10 items, max 2000 chars each
	//accessory - not implemented
}

func typeSection() *Section {
	return &Section{Type: TypeSection}
}

func New(txt *text.Text, fields ...*text.Text) *Section {

	s := typeSection()
	s.Text = txt

	max2000 := func(f []*text.Text) []*text.Text {
		for i, t := range f {
			if len(t.Text) > 2000 {
				f[i] = t.FirstN(2000)
			}
		}
		return f
	}

	first10 := func(f []*text.Text) []*text.Text {
		if len(f) <= 10 {
			return f
		} else {
			return f[:10]
		}
	}

	s.Fields = max2000(first10(fields))

	return s
}

func NewFields(key, value *text.Text) *Section {
	s := typeSection()
	s.Fields = []*text.Text{key.FirstN(2000), value.FirstN(2000)}

	return s
}

func (s *Section) AddFields(key, value *text.Text) *Section {

	// xor for text versus fields
	if len(s.Fields) <= 8 && len(s.Text.Text) == 0 {
		s.Fields = append(s.Fields, key.FirstN(2000))
		s.Fields = append(s.Fields, value.FirstN(2000))
	}

	return s
}

func (s *Section) Json() string {
	js, err := json.Marshal(s)
	if err != nil {
		return "{}"
	}

	return string(js)
}
