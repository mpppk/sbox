package scrapbox

import (
	"errors"
	"fmt"
	"strings"
)

type Text fmt.Stringer

type PlainText struct {
	Text string
}

func (p *PlainText) String() string {
	return p.Text
}

func trimBrackets(text string) (string, error) {
	if !hasBrackets(text) {
		return "", errors.New("invalid texts")
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
