package scrapbox

import (
	"fmt"
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	server := "https://scrapbox.io"
	project := "testproject"

	plainTextCases := []string{
		"foobar",
		"not link]",
		"[not link",
	}

	for _, c := range plainTextCases {
		parsedTexts, err := Parse(c, server, project)
		if err != nil {
			t.Fatalf("Unexpected error occured in Parse if text %s is given: %s", c, err)
		}

		concatenateText := ""
		for _, text := range parsedTexts {
			typeStr := "*scrapbox.PlainText"
			if reflect.TypeOf(text).String() != typeStr {
				t.Fatalf("Parse() is expected to return %q if text %q is given", typeStr, c)
			}
			concatenateText += text.String()
		}

		if concatenateText != c {
			t.Fatalf("If structs that returned from Parse() is joined as string, "+
				"it is expected to be same as argument(%q), but actually it got %q",
				c, concatenateText)
		}
	}

	linkTextCases := []struct {
		text     string
		expected Link
	}{
		{
			text: "[link]",
			expected: Link{
				Server:  server,
				Project: project,
				Title:   "link",
				URL:     fmt.Sprintf("%s/%s/link", server, project),
			},
		},
		{
			text: "[*NotBold]",
			expected: Link{
				Server:  server,
				Project: project,
				Title:   "*NotBold",
				URL:     fmt.Sprintf("%s/%s/%s", server, project, "%2ANotBold"),
			},
		},
		{
			text: "[https://sample.com sample link]",
			expected: Link{
				Server:  server,
				Project: project,
				Title:   "sample link",
				URL:     "https://sample.com",
			},
		},
	}

	for _, c := range linkTextCases {
		parsedTexts, err := Parse(c.text, server, project)
		if err != nil {
			t.Fatalf("Unexpected error occured in Parse if text %s is given: %s", c.text, err)
		}

		parsedText, ok := parsedTexts[0].(*Link)
		if !ok {
			t.Fatalf("Parse() is expected to return *scrapbox.Link if text %q is given", c.text)
		}

		if parsedText.String() != c.text {
			t.Fatalf("If structs that returned from Parse() is joined as string, "+
				"it is expected to be same as argument(%q), but actually it got %q",
				c, parsedText)
		}

		if parsedText.Server != c.expected.Server {
			t.Fatalf("Link.Server that returned from Parse() is expected to be %q, if text %q is given, "+
				"but actually it is %q",
				c.expected.Server, c.text, parsedText.Server)
		}

		if parsedText.Project != c.expected.Project {
			t.Fatalf("Link.Project that returned from Parse() is expected to be %q, if text %q is given, "+
				"but actually it is %q",
				c.expected.Project, c.text, parsedText.Project)
		}

		if parsedText.Title != c.expected.Title {
			t.Fatalf("Link.Title that returned from Parse() is expected to be %q, if text %q is given, "+
				"but actually it is %q",
				c.expected.Title, c.text, parsedText.Title)
		}

		if parsedText.URL != c.expected.URL {
			t.Fatalf("Link.URL that returned from Parse() is expected to be %q, if text %q is given, "+
				"but actually it is %q",
				c.expected.URL, c.text, parsedText.URL)
		}

	}

	cases := []struct {
		texts    string
		textType string
	}{
		{
			texts: "[* Bold]",
		},
		{
			texts: "[/ Italic]",
		},
		{
			texts: "[- Strike Through]",
		},
		{
			texts: "[https://sample.com sample link]and[* Bold]Text " +
				"and[/ Italic]text and [- Strike]text",
		},
	}

	for _, c := range cases {
		parsedTexts, err := Parse(c.texts, server, project)
		if err != nil {
			t.Fatalf("Unexpected error occured in Parse: %s", err)
		}

		concatenateText := ""
		for _, text := range parsedTexts {
			fmt.Printf("%#v\n", text)
			concatenateText += text.String()
		}

		if concatenateText != c.texts {
			t.Fatalf("If structs that returned from Parse() is joined as string, "+
				"it is expected to be same as argument(%q), but actually it got %q",
				c.texts, concatenateText)
		}
	}
}
