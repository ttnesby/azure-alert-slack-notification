package section

import (
	"fmt"
	"testing"

	"github.com/ttnesby/azure-alert-slack-notification/caddy-ext/pkg/slack/object/text"
)

func listJson(a []*text.Text) string {
	json := ""
	for _, t := range a {
		if json == "" {
			json += t.Json()
		} else {
			json += fmt.Sprintf(",%s", t.Json())
		}

	}
	return fmt.Sprintf("[%s]", json)
}

func TestText(t *testing.T) {
	t.Parallel()

	expected := func(s string) string {
		return fmt.Sprintf(`{"type":"%s","text":%s}`, CSection, s)
	}
	txt := text.NewPlain("hei på deg")

	if New().SetText(txt).Json() != expected(txt.Json()) {
		t.Errorf("%s", "json failure")
	}

	txtM := text.NewMarkDown("hei *på* deg")

	if New().SetText(txtM).Json() != expected(txtM.Json()) {
		t.Errorf("%s", "json failure")
	}
}

func TestFields(t *testing.T) {
	t.Parallel()

	expected := func(s string) string {
		return fmt.Sprintf(`{"type":"%s","fields":%s}`, CSection, s)
	}

	to := []*text.Text{
		text.NewMarkDown("*k1*"),
		text.NewPlain("v2"),
		text.NewMarkDown("*k3*"),
		text.NewPlain("v4"),
		text.NewMarkDown("*k5*"),
		text.NewPlain("v6"),
		text.NewMarkDown("*k7*"),
		text.NewPlain("v8"),
		text.NewMarkDown("*k9*"),
		text.NewPlain("v10"),
		text.NewMarkDown("*k11*"),
		text.NewPlain("v12"),
	}

	// >10 elements returns exactly 10 elements
	sec := New().SetFields(to...)

	if sec.Json() != expected(listJson(to[:10])) {
		t.Errorf("%s", "json failure")
	}

	// < 10 returns the same number of elements
	sec = New().SetFields(to[:4]...)

	if sec.Json() != expected(listJson(to[:4])) {
		t.Errorf("%s", "json failure")
	}
}
