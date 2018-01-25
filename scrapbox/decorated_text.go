package scrapbox

import (
	"fmt"

	"github.com/pkg/errors"
)

type decoratedText struct {
	Symbol rune
	Text   string
}

func (d *decoratedText) String() string {
	return fmt.Sprintf("[%c %s]", d.Symbol, d.Text)
}

type BoldText decoratedText
type ItalicText decoratedText
type StrikeThroughText decoratedText

func NewBoldText(rawText string) (*BoldText, error) {
	decoratedText, err := newDecoratedText(rawText, '*', "bold")
	return (*BoldText)(decoratedText), err
}

func (t *BoldText) String() string {
	return fmt.Sprintf("[%c %s]", t.Symbol, t.Text)
}

func NewItalicText(rawText string) (*ItalicText, error) {
	decoratedText, err := newDecoratedText(rawText, '/', "italic")
	return (*ItalicText)(decoratedText), err
}

func (t *ItalicText) String() string {
	return fmt.Sprintf("[%c %s]", t.Symbol, t.Text)
}

func NewStrikeThroughText(rawText string) (*StrikeThroughText, error) {
	decoratedText, err := newDecoratedText(rawText, '-', "strike through")
	return (*StrikeThroughText)(decoratedText), err
}

func (t *StrikeThroughText) String() string {
	return fmt.Sprintf("[%c %s]", t.Symbol, t.Text)
}

func newDecoratedText(rawText string, symbol rune, decorateName string) (*decoratedText, error) {
	trimmedBoldRawText, err := trimDecoratedRawText(rawText, symbol)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("invalid text for %s text", decorateName))
	}
	return &decoratedText{Text: trimmedBoldRawText, Symbol: symbol}, nil
}
