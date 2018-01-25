package scrapbox

import (
	"errors"
	"fmt"
)

type decoratedText struct {
	Symbol rune
	Text   string
}

type BoldText decoratedText

func NewBoldText(rawText string) (*BoldText, error) {
	trimmedBoldRawText, err := trimDecoratedRawText(rawText, '*')
	if err != nil {
		return nil, errors.New("invalid text for bold text")
	}
	return &BoldText{Text: trimmedBoldRawText, Symbol: '*'}, nil
}

func (b *BoldText) String() string {
	return fmt.Sprintf("[%c %s]", b.Symbol, b.Text)
}
