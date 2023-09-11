package section

import (
	"fmt"
	"github.com/ttnesby/slack-block-builder/pkg/slack/object/text"
	"testing"
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
		return fmt.Sprintf(`{"type":"%s","text":%s}`, TypeSection, s)
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
		return fmt.Sprintf(`{"type":"%s","fields":%s}`, TypeSection, s)
	}

	to := []*text.Text{text.NewMarkDown("*k1*"), text.NewPlain("v1")}
	sec := New().SetFields(to[0], to[1])

	if sec.Json() != expected(listJson(to)) {
		t.Errorf("%s", "json failure")
	}
}
