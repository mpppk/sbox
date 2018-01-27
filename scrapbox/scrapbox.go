package scrapbox

import (
	"errors"
	"strings"
)

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

func isValidDecoratedRawText(rawText string, decoratedSymbol rune) bool {
	trimmedRawText, err := trimBrackets(rawText)
	if err != nil {
		return false
	}
	trimmedRawTextList := []rune(trimmedRawText)
	return len(trimmedRawTextList) > 0 &&
		trimmedRawTextList[0] == decoratedSymbol &&
		trimmedRawTextList[1] == ' '
}

func trimDecoratedRawText(rawText string, decoratedSymbol rune) (string, error) {
	if !isValidDecoratedRawText(rawText, decoratedSymbol) {
		return "", errors.New("invalid text for bold text")
	}

	trimmedRawText, _ := trimBrackets(rawText) // error is already checked
	return strings.Replace(
		strings.Replace(trimmedRawText, string(decoratedSymbol), "", 1),
		" ", "", 1), nil
}
