package text

import (
	"fmt"
	"testing"
)

func TestPlain(t *testing.T) {

	t.Parallel()

	expected := func(s string) string {
		return fmt.Sprintf(`{"type":"%s","text":"%s","emoji":true,"verbatim":false}`, TypePlain, s)
	}

	str := "hei på deg"

	text := NewPlain(str)

	if text.Json() != expected(str) {
		t.Errorf("%s", "json failure")
	}

	n := 2
	if str[:n] != text.FirstN(n).Text {
		t.Errorf("%s", "FirstN failure")
	}

	if NewPlain("").Json() != expected(EmptyText) {
		t.Errorf("%s", "empty text jason failure")
	}
}

func TestMarkDown(t *testing.T) {

	t.Parallel()

	expected := func(s string) string {
		return fmt.Sprintf(`{"type":"%s","text":"%s","emoji":false,"verbatim":false}`, TypeMarkDown, s)
	}

	str := "hei *på* deg"

	if NewMarkDown(str).Json() != expected(str) {
		t.Errorf("%s", "json failure")
	}

	if NewMarkDown("").Json() != expected(EmptyText) {
		t.Errorf("%s", "empty text jason failure")
	}
}
