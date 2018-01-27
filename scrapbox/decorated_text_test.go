package scrapbox

import (
	"testing"
)

func TestNewBoldText(t *testing.T) {
	cases := []struct {
		rawText        string
		expectedText   string
		expectedString string
		willBeError    bool
	}{
		{
			rawText:     "* foobar]",
			willBeError: true,
		},
		{
			rawText:     "[ foobar]",
			willBeError: true,
		},
		{
			rawText:     "[* foobar",
			willBeError: true,
		},
		{
			rawText:     "[*foobar]",
			willBeError: true,
		},
		{
			rawText:        "[* foobar]",
			expectedText:   "foobar",
			expectedString: "[* foobar]",
		},
		{
			rawText:        "[* foo bar]",
			expectedText:   "foo bar",
			expectedString: "[* foo bar]",
		},
	}

	for _, c := range cases {
		boldText, err := NewBoldTextFromBracketsText(c.rawText)
		if err != nil && !c.willBeError {
			t.Fatalf("Unexpected error occured in NewBoldText: %s", err)
		}

		if c.willBeError {
			if err == nil {
				t.Fatalf("NewBoldText is expectedText to be error if argumet %q is given.", c.rawText)
			} else {
				continue
			}
		}

		if boldText.Text != c.expectedText {
			t.Fatalf("BoldText.Text is expected to return %q when rawText %q is given, but actually has %q",
				c.expectedText, c.rawText, boldText)
		}

		if boldText.String() != c.expectedString {
			t.Fatalf("BoldText.String() is expected to return %q when rawText %q is given, but actually has %q",
				c.expectedText, c.rawText, boldText)
		}
	}
}
