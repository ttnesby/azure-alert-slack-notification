package section

import (
	"fmt"
	"github.com/ttnesby/slack-block-builder/pkg/slack/object/text"
	"testing"
)

func TestText(t *testing.T) {

	t.Parallel()

	expected := func(s string) string {
		return fmt.Sprintf(`{"type":"%s","text":%s}`, TypeSection, s)
	}
	txt := text.NewPlain("hei på deg")

	if NewText(txt).Json() != expected(txt.Json()) {
		t.Errorf("%s", "json failure")
	}

	txtM := text.NewMarkDown("hei *på* deg")

	if NewText(txtM).Json() != expected(txtM.Json()) {
		t.Errorf("%s", "json failure")
	}

	// cannot add fields to text section, returning unchanged struct
	sec := NewText(txt).AddFields(text.NewMarkDown("*key*"), text.NewPlain("value"))
	if sec.Json() != expected(txt.Json()) {
		t.Errorf("%s", "json failure")
	}
}

func TestFields(t *testing.T) {
	t.Parallel()

	expected := func(s string) string {
		return fmt.Sprintf(`{"type":"%s","fields":%s}`, TypeSection, s)
	}

	to := []*text.Text{text.NewMarkDown("*k1*"), text.NewPlain("v1")}

	toJson := func(to []*text.Text) string {
		json := ""
		for _, t := range to {
			if json == "" {
				json += t.Json()
			} else {
				json += fmt.Sprintf(",%s", t.Json())
			}

		}
		return fmt.Sprintf("[%s]", json)
	}

	sec := NewFields(to[0], to[1])

	if sec.Json() != expected(toJson(to)) {
		t.Errorf("%s", "json failure")
	}

	fmt.Printf("%s\n", sec.Json())
	fmt.Printf("%s\n", expected(toJson(to)))
}
