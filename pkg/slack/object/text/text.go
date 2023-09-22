package text

import "encoding/json"

// https://api.slack.com/reference/block-kit/composition-objects#text

type TextType string

const (
	Plain    TextType = "plain_text"
	MarkDown TextType = "mrkdwn"

	EmptyText = "M"
)

type Text struct {
	Type  TextType `json:"type"`            // plain_text, mrkdwn
	Text  string   `json:"text"`            // min 1 and max 3 000
	Emoji *bool    `json:"emoji,omitempty"` // only relevant for plain_text
	//verbatim - not implemented
}

func NewPlain(s string) *Text {
	t := new(bool)
	*t = true

	return (&Text{
		Type:  Plain,
		Text:  s,
		Emoji: t,
	}).FirstN(3000)
}

func NewMarkDown(s string) *Text {
	return (&Text{
		Type: MarkDown,
		Text: s,
	}).FirstN(3000)
}

func (t *Text) FirstN(n int) *Text {
	switch l := len(t.Text); {
	case l == 0: // min
		t.Text = EmptyText
	case l > n: // max
		t.Text = t.Text[:n]
	}

	return t
}

func (t *Text) Json() string {
	js, err := json.Marshal(t)
	if err != nil {
		return "{}"
	}

	return string(js)
}
