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

// IconRotatingLight Icon = ":rotating_light:"
const (
	IconLink Icon = ":link:"

	// custom severity codes for testing
	SevTestStart Severity = ":test: START"
	SevTestEnd   Severity = ":test: END"

	SeverityUnknown     Severity = ":question: unknown"
	SeverityVerbose     Severity = ":speech_balloon:  Verbose(4)"
	SeverityInformation Severity = ":information_source:  Information(3)"
	SeverityWarning     Severity = ":warning:  Warning(2)"
	SeverityError       Severity = ":error:  Error(1)"
	SeverityCritical    Severity = ":severity-critical:  Critical(0)"
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

func (n *Notification) Json() []byte {
	js, err := json.Marshal(n)
	if err != nil {
		return []byte("{}")
	}

	return js
}
