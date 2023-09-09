package text

// https://api.slack.com/reference/block-kit/composition-objects#text

type Text struct {
	Type     string `json:"type"`
	Text     string `json:"text"`
	Emoji    bool   `json:"emoji"`
	Verbatim bool   `json:"verbatim"`
}

func NewText(text string) *Text {
	return &Text{
		Type:     "plain_text",
		Text:     text,
		Emoji:    true,
		Verbatim: true,
	}
}

func NewMarkDown(text string) *Text {
	return &Text{
		Type:     "mrkdwn",
		Text:     text,
		Emoji:    false,
		Verbatim: true,
	}
}
