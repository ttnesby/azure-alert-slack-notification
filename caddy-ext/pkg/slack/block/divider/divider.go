package divider

// https://api.slack.com/reference/block-kit/blocks#divider

type Divider struct {
	Type string `json:"type"`
}

func New() *Divider {
	return &Divider{
		Type: "divider",
	}
}
