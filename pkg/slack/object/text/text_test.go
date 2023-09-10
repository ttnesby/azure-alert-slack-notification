package text

import (
	"encoding/json"
	"fmt"
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

	text := NewPlain(testText)
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

func TestMarkDown(t *testing.T) {

	t.Parallel()

	fail := func() {
		t.Errorf("%s", "failure")
	}

	testText := "hei på deg"
	n := 2
	expected := fmt.Sprintf(`{"type":"mrkdwn","text":"%s","emoji":false,"verbatim":false}`, testText)

	text := NewMarkDown(testText)
	got, err := json.Marshal(text)

	if err != nil {
		t.Errorf("failed!")
	}

	if string(got) != expected {
		t.Errorf("failure")
	}

	if len(testText) != text.Len() {
		fail()
	}

	if testText[:n] != text.FirstN(n).Text {
		fail()
	}
}
