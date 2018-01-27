package scrapbox

import (
	"net/url"
	"testing"
)

func TestHasBrackets(t *testing.T) {
	cases := []struct {
		text     string
		expected bool
	}{
		{
			text:     "hoge",
			expected: false,
		},
		{
			text:     "https://example.com Example!]",
			expected: false,
		},
		{
			text:     "[https://example.com Example!",
			expected: false,
		},
		{
			text:     "[https://example.com Example!]",
			expected: true,
		},
	}

	for _, c := range cases {
		if ok := hasBrackets(c.text); ok != c.expected {
			t.Fatalf("hasBrackets is expected to return %q when text %q is given, but actually has %q",
				c.expected, c.text, ok)
		}
	}
}

func TestTrimBrackets(t *testing.T) {
	cases := []struct {
		text        string
		expected    string
		willBeError bool
	}{
		{
			text:        "hoge",
			willBeError: true,
		},
		{
			text:        "[https://example.com Example!]",
			expected:    "https://example.com Example!",
			willBeError: false,
		},
	}

	for _, c := range cases {
		trimmedText, err := trimBrackets(c.text)

		if err != nil && !c.willBeError {
			t.Fatalf("Error occured in trimBrackets: %s", err)
		}

		if c.willBeError {
			if err == nil {
				t.Fatalf("trimBrackets is expected to be error if raw string %q is given.", c.text)
			} else {
				continue
			}
		}

		if trimmedText != c.expected {
			t.Fatalf("trimBrackets is expected to return %q when text %q is given, but actually has %q",
				c.expected, c.text, trimmedText)
		}
	}
}

func TestParseTrimmedLinkText(t *testing.T) {
	cases := []struct {
		text          string
		server        string
		project       string
		expectedURL   string
		expectedTitle string
		willBeError   bool
		description   string
	}{
		{
			text:          "http://example.com Example!",
			server:        "https://scrapbox.io",
			project:       "niboshi",
			expectedURL:   "http://example.com",
			expectedTitle: "Example!",
			willBeError:   false,
		},
		{
			text:          "Example! http://example.com",
			server:        "https://scrapbox.io",
			project:       "niboshi",
			expectedURL:   "http://example.com",
			expectedTitle: "Example!",
			willBeError:   false,
		},
		{
			text:          "Example! Other text!! http://example.com",
			server:        "https://scrapbox.io",
			project:       "niboshi",
			expectedURL:   "http://example.com",
			expectedTitle: "Example! Other text!!",
			willBeError:   false,
		},
		{
			text:          "http://example.com Example! Other text!!",
			server:        "https://scrapbox.io",
			project:       "niboshi",
			expectedURL:   "http://example.com",
			expectedTitle: "Example! Other text!!",
			willBeError:   false,
		},
		{
			text:          "Example",
			server:        "https://scrapbox.io",
			project:       "niboshi",
			expectedURL:   "https://scrapbox.io/niboshi/Example",
			expectedTitle: "Example",
			willBeError:   false,
		},
		{
			text:          "Example other text",
			server:        "https://scrapbox.io",
			project:       "niboshi",
			expectedURL:   "https://scrapbox.io/niboshi/Example%20other%20texts",
			expectedTitle: "Example other text",
			willBeError:   false,
		},
		{
			text:          "/project/page",
			server:        "https://scrapbox.io",
			project:       "niboshi",
			expectedURL:   "https://scrapbox.io/project/page",
			expectedTitle: "/project/page",
			willBeError:   false,
		},
		{
			text:          "/project",
			server:        "https://scrapbox.io",
			project:       "niboshi",
			expectedURL:   "https://scrapbox.io/project",
			expectedTitle: "/project",
			willBeError:   false,
		},
		{
			text:          "http://example.com 日本語　の　テスト",
			server:        "https://scrapbox.io",
			project:       "niboshi",
			expectedURL:   "http://example.com",
			expectedTitle: "日本語　の　テスト",
			willBeError:   false,
		},
		{
			text:          "日本語ページ",
			server:        "https://scrapbox.io",
			project:       "niboshi",
			expectedURL:   "https://scrapbox.io/niboshi/" + url.PathEscape("日本語ページ"),
			expectedTitle: "日本語ページ",
			willBeError:   false,
		},
	}

	for _, c := range cases {
		linkURL, linkTitle, err := parseTrimmedLinkText(c.server, c.project, c.text)

		if err != nil && !c.willBeError {
			t.Fatalf("Error occured in parseTrimmedLinkText: %s, test case: %q", err, c.description)
		}

		if c.willBeError {
			if err == nil {
				t.Fatalf("parseTrimmedLinkText is expected to be error if raw string %q is given."+
					" test case: %s", c.text, c.description)
			} else {
				continue
			}
		}

		if linkURL != c.expectedURL || linkTitle != c.expectedTitle {
			t.Fatalf("parseTrimmedLinkText is expected to return URL: %q "+
				"and Title: %q, when raw string %q is given, but actually URL: %q, and Title: %q test case: %s",
				c.expectedURL, c.expectedTitle, c.text, linkURL, linkTitle, c.description)
		}
	}
}

func TestNewSBLink(t *testing.T) {
	cases := []struct {
		text          string
		server        string
		project       string
		expectedURL   string
		expectedTitle string
		willBeError   bool
		description   string
	}{
		{
			text:        "http://example.com Example!",
			server:      "https://scrapbox.io",
			project:     "niboshi",
			willBeError: true,
			description: "text has no brackets",
		},
		{
			text:          "[http://example.com Example!]",
			server:        "https://scrapbox.io",
			project:       "niboshi",
			expectedURL:   "http://example.com",
			expectedTitle: "Example!",
			willBeError:   false,
		},
		{
			text:          "[Example! http://example.com]",
			server:        "https://scrapbox.io",
			project:       "niboshi",
			expectedURL:   "http://example.com",
			expectedTitle: "Example!",
			willBeError:   false,
		},
	}

	for _, c := range cases {
		link, err := NewSBLink(c.text, c.server, c.project)

		if err != nil && !c.willBeError {
			t.Fatalf("Error occured when Link creating: %s, test case: %s", err, c.description)
		}

		if c.willBeError {
			if err == nil {
				t.Fatalf("NewSBLink is expected to be error if raw string %q is given."+
					" test case: %s", c.text, c.description)
			} else {
				continue
			}
		}

		if link.Title != c.expectedTitle || link.URL != c.expectedURL {
			t.Fatalf("NewSBLink is expected to return URL: %q "+
				"and Title: %q, when raw string %q is given, but actually URL: %q, and Title: %q test case: %s",
				c.expectedURL, c.expectedTitle, c.text, link.URL, link.Title, c.description)
		}

		if link.Title != c.expectedTitle {
			t.Fatalf("Error occured when Link creating: %s", err)
		}
	}
}
