package scrapbox

import (
	"errors"
	"fmt"
	"net/url"
	"strings"
)

type SBLink struct {
	SBNode
	Title string
	URL   string
}

func NewSBLink(text, server, project string) (*SBLink, error) {
	trimmedText, err := trimBrackets(text)
	if err != nil {
		return nil, err
	}
	linkUrl, linkText, err := parseTrimmedLinkText(server, project, trimmedText)
	return &SBLink{Title: linkText, URL: linkUrl}, nil
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

	escapedText := url.PathEscape(trimmedText)
	linkUrl := fmt.Sprintf("%s/%s/%s", server, project, escapedText)
	return linkUrl, trimmedText, nil
}

func trimBrackets(text string) (string, error) {
	if !hasBrackets(text) {
		return "", errors.New("invalid text")
	}

	textList := strings.Split(text, "")
	return strings.Join(textList[1:len(textList)-1], ""), nil
}

func hasBrackets(text string) bool {
	textList := strings.Split(text, "")
	return len(textList) > 0 &&
		textList[0] == "[" &&
		textList[len(textList)-1] == "]"
}
