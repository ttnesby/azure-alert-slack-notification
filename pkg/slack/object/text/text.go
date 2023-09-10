package text

// https://api.slack.com/reference/block-kit/composition-objects#text

type base struct {
	Type     string `json:"type"` // plain_text, mrkdwn
	Text     string `json:"text"` // min 1 and max 3 000
	Emoji    bool   `json:"emoji"`
	Verbatim bool   `json:"verbatim"`
}

type Plain struct{ base }
type MarkDown struct{ base }

type Text[T Plain | MarkDown] interface {
	Plain | MarkDown

	Len() int
	FirstN(n int) *T
}

func NewPlain(text string) *Plain {
	return (&Plain{
		base{
			Type:     "plain_text",
			Text:     text,
			Emoji:    true,
			Verbatim: false,
		},
	}).FirstN(3000)
}

func NewMarkDown(text string) *MarkDown {
	return (&MarkDown{
		base{
			Type:     "mrkdwn",
			Text:     text,
			Emoji:    false,
			Verbatim: false,
		},
	}).FirstN(3000)
}

func (b *base) Len() int {
	return len(b.Text)
}

func (b *base) FirstN(n int) *base {
	if b.Len() > n {
		b.Text = b.Text[:n]
	}

	return b
}

func (p *Plain) FirstN(n int) *Plain {
	return &Plain{*p.base.FirstN(n)}
}

func (m *MarkDown) FirstN(n int) *MarkDown {
	return &MarkDown{*m.base.FirstN(n)}
}
