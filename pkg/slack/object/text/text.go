package text

import "encoding/json"

// https://api.slack.com/reference/block-kit/composition-objects#text

const (
	TypePlain    = "plain_text"
	TypeMarkDown = "mrkdwn"
	EmptyText    = "M"
)

type Text struct {
	Type  string `json:"type"`            // plain_text, mrkdwn
	Text  string `json:"text"`            // min 1 and max 3 000
	Emoji *bool  `json:"emoji,omitempty"` // only relevant for plain_text
	//verbatim - not implemented
}

func NewPlain(s string) *Text {
	t := new(bool)
	*t = true

	return (&Text{
		Type:  TypePlain,
		Text:  s,
		Emoji: t,
	}).FirstN(3000)
}

func NewMarkDown(s string) *Text {
	return (&Text{
		Type: TypeMarkDown,
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
