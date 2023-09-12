package notification

import (
	"encoding/json"
	"github.com/ttnesby/slack-block-builder/pkg/slack/block/action"
	"github.com/ttnesby/slack-block-builder/pkg/slack/block/divider"
	"github.com/ttnesby/slack-block-builder/pkg/slack/block/header"
	"github.com/ttnesby/slack-block-builder/pkg/slack/block/section"
)

type Icon string
type Severity string

const (
	IconRotatingLight Icon = ":rotating_light:"
	IconLink          Icon = ":link:"

	SeverityUnknown     Severity = ":question: unknown"
	SeverityVerbose     Severity = ":speech_balloon:  4 - Verbose"
	SeverityInformation Severity = ":information_source:  3 - Information"
	SeverityWarning     Severity = ":warning:  2 - Warning"
	SeverityError       Severity = ":error:  1 - Error"
	SeverityCritical    Severity = ":severity-critical: 0 - Critical"
)

type Notification struct {
	Blocks []any `json:"blocks"`
}

func New() *Notification {
	return &Notification{}
}

type content interface {
	action.Action | divider.Divider | header.Header | section.Section
}

func add[T content](n *Notification, x *T) *Notification {
	if x == nil {
		return n
	}

	n.Blocks = append(n.Blocks, x)

	return n
}

func (n *Notification) AddSection(x *section.Section) *Notification {
	return add[section.Section](n, x)
}

func (n *Notification) AddDivider(x *divider.Divider) *Notification {
	return add[divider.Divider](n, x)
}

func (n *Notification) AddAction(x *action.Action) *Notification {
	return add[action.Action](n, x)
}

func (n *Notification) AddHeader(x *header.Header) *Notification {
	return add[header.Header](n, x)
}

func (n *Notification) Json() string {
	js, err := json.Marshal(n)
	if err != nil {
		return "{}"
	}

	return string(js)
}
