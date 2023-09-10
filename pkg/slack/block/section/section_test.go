package section

import (
	"encoding/json"
	"fmt"
	"github.com/ttnesby/slack-block-builder/pkg/slack/object/text"
	"testing"
)

func TestText(t *testing.T) {

	t.Parallel()

	fail := func() {
		t.Errorf("%s", "failure")
	}

	testText := "hei på deg"
	expected := fmt.Sprintf(`{"type":"section","text":{"type":"plain_text","text":"%s","emoji":true,"verbatim":false}}`, testText)

	sec := NewText(text.NewPlain(testText))
	got, err := json.Marshal(sec)

	if err != nil {
		fail()
	}

	if string(got) != expected {
		fail()
	}

	fmt.Printf("%s", got)
}
