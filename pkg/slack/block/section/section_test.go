package section

import (
	"encoding/json"
	"fmt"
	text2 "github.com/ttnesby/slack-block-builder/pkg/slack/object/text"
	"testing"
)

func TestPlain(t *testing.T) {

	t.Parallel()

	fail := func() {
		t.Errorf("%s", "failure")
	}

	testText := "hei på deg"
	n := 2
	expected := fmt.Sprintf(`{"type":"plain_text","text":"%s","emoji":true,"verbatim":false}`, testText)

	text := NewText[text.Plain](text2.NewPlain(testText))
	got, err := json.Marshal(text)

	if err != nil {
		fail()
	}

	if string(got) != expected {
		fail()
	}

	if len(testText) != text.Len() {
		fail()
	}

	if testText[:n] != text.FirstN(n).Text {
		fail()
	}
}
