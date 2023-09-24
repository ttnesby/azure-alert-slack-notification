package text

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestText(t *testing.T) {
	t.Parallel()

	const (
		ie = ""
		it = "hei på deg!"
	)

	expect := func(t TextType, s string) string {
		if t == Plain {
			return fmt.Sprintf(`{"type":"%s","text":"%s","emoji":true}`, t, s)
		} else {
			return fmt.Sprintf(`{"type":"%s","text":"%s"}`, t, s)
		}
	}

	expectJsonString := func(t TextType, s string) string {
		js, _ := json.Marshal(s)
		if t == Plain {
			return fmt.Sprintf(`{"type":"%s","text":%s,"emoji":true}`, t, string(js))
		} else {
			return fmt.Sprintf(`{"type":"%s","text":%s}`, t, string(js))
		}
	}

	testData := []struct {
		input  string
		fun    func(string) *Text
		expect string
	}{
		{
			input:  ie,
			fun:    NewPlain,
			expect: expect(Plain, EmptyText),
		},
		{
			input:  it,
			fun:    NewPlain,
			expect: expect(Plain, it),
		},
		{
			input:  long3003,
			fun:    NewPlain,
			expect: expectJsonString(Plain, long3000),
		},
		{
			input:  ie,
			fun:    NewMarkDown,
			expect: expect(MarkDown, EmptyText),
		},
		{
			input:  it,
			fun:    NewMarkDown,
			expect: expect(MarkDown, it),
		},
		{
			input:  long3003,
			fun:    NewMarkDown,
			expect: expectJsonString(MarkDown, long3000),
		},
	}

	for i, td := range testData {
		actual := td.fun(td.input).Json()
		if actual != td.expect {
			t.Errorf("Test %d: Got %s but should have: %s", i, actual, td.expect)
		}
	}
}
