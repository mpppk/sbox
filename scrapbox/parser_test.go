package scrapbox

import (
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	cases := []struct {
		texts       string
		expected    []Text
		willBeError bool
	}{
		{
			texts:       "foobar",
			willBeError: false,
		},
		{
			texts:       "[* Bold]",
			willBeError: false,
		},
		{
			texts:       "[*NotBold]",
			willBeError: false,
		},
		{
			texts:       "[link]",
			willBeError: false,
		},
		{
			texts:       "not link]",
			willBeError: false,
		},
		{
			texts:       "[not link",
			willBeError: false,
		},
		{
			texts:       "[https://sample.com sample link]",
			willBeError: false,
		},
		{
			texts:       "[https://sample.com sample link]and[* Bold]Text",
			willBeError: false,
		},
	}

	for _, c := range cases {
		parsedTexts, err := Parse(c.texts)
		if err != nil && !c.willBeError {
			t.Fatalf("Unexpected error occured in Parse: %s", err)
		}

		if c.willBeError {
			if err == nil {
				t.Fatalf("Parse is expected to be error if argumet %q is given.", c.texts)
			} else {
				continue
			}
		}

		parsedTextsStr := ""
		for _, text := range parsedTexts {
			fmt.Printf("%#v\n", text)
			parsedTextsStr += text.String()
		}

		if parsedTextsStr != c.texts {
			t.Fatalf("If structs that returned from Parse() is joined as string, "+
				"it is expected to be same as argument(%q), but actually it got %q",
				c.texts, parsedTextsStr)
		}
	}
}
