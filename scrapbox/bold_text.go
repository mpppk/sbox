package scrapbox

import (
	"errors"
	"fmt"
	"strings"
)

type BoldText Text

func NewBoldText(rawText string) (*BoldText, error) {
	trimmedBoldRawText, err := trimBoldRawText(rawText)
	if err != nil {
		return nil, errors.New("invalid text for bold text")
	}
	return &BoldText{Text: trimmedBoldRawText}, nil
}

func (b *BoldText) String() string {
	return fmt.Sprintf("[* %s]", b.Text)
}

func isValidBoldRawText(rawText string) bool {
	trimmedRawText, err := trimBrackets(rawText)
	if err != nil {
		return false
	}
	trimmedRawTextList := strings.Split(trimmedRawText, "")
	return len(trimmedRawTextList) > 0 &&
		trimmedRawTextList[0] == "*" &&
		trimmedRawTextList[1] == " "
}

func trimBoldRawText(rawText string) (string, error) {
	if !isValidBoldRawText(rawText) {
		return "", errors.New("invalid text for bold text")
	}

	trimmedRawText, _ := trimBrackets(rawText) // error is already checked
	return strings.Replace(
		strings.Replace(trimmedRawText, "*", "", 1),
		" ", "", 1), nil
}
