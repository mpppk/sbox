package scrapbox

import (
	"fmt"
	"net/url"
	"strings"
)

type Link struct {
	OriginalText string
	Server       string
	Project      string
	Title        string
	URL          string
}

func (l *Link) String() string {
	if l.OriginalText != "" {
		return l.OriginalText
	}

	return fmt.Sprintf("[%s %s]", l.Title, l.URL) // TODO related url
}

func NewSBLink(text, server, project string) (*Link, error) {
	trimmedText, err := trimBrackets(text)
	if err != nil {
		return nil, err
	}
	linkUrl, linkText, err := parseTrimmedLinkText(server, project, trimmedText)
	return &Link{OriginalText: text, Title: linkText, URL: linkUrl, Server: server, Project: project}, nil
}

func parseTrimmedLinkText(server, project, trimmedText string) (string, string, error) {
	urlAndTexts := strings.Split(trimmedText, " ")
	if strings.Contains(urlAndTexts[0], "http") {
		linkText := strings.Join(urlAndTexts[1:], " ")
		return urlAndTexts[0], linkText, nil
	}

	lastIndex := len(urlAndTexts) - 1
	if strings.Contains(urlAndTexts[lastIndex], "http") {
		linkText := strings.Join(urlAndTexts[:lastIndex], " ")
		return urlAndTexts[lastIndex], linkText, nil
	}

	if strings.HasPrefix(trimmedText, "/") {
		texts := strings.Split(trimmedText, "/")[1:]
		if len(texts) > 1 {
			escapedText := url.PathEscape(strings.Join(texts[1:], "/"))
			linkUrl := fmt.Sprintf("%s/%s/%s", server, texts[0], escapedText)
			return linkUrl, trimmedText, nil
		}
		linkUrl := fmt.Sprintf("%s/%s", server, texts[0])
		return linkUrl, trimmedText, nil
	}

	escapedText := url.PathEscape(trimmedText)
	linkUrl := fmt.Sprintf("%s/%s/%s", server, project, escapedText)
	return linkUrl, trimmedText, nil
}
