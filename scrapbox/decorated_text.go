package scrapbox

import (
	"fmt"

	"github.com/pkg/errors"
)

type TextStringer interface {
	fmt.Stringer
	GetText() string
}

type decoratedText struct {
	Symbol rune
	Text   string
}

func (d *decoratedText) String() string {
	return fmt.Sprintf("[%c %s]", d.Symbol, d.Text)
}

type NewLineText struct{}
type BulletPointText struct {
	Text  string
	Level int
}
type BoldText decoratedText
type ItalicText decoratedText
type StrikeThroughText decoratedText

func (t *BulletPointText) String() string {
	return " " + t.Text
}

func (t *BulletPointText) GetText() string {
	return t.Text
}

func NewBulletPointText(rawText string) (*BulletPointText, error) {
	for i, r := range rawText {
		if r == ' ' {
			continue
		}
		if i == 0 {
			return nil, errors.New("invalid raw text for bullet point")
		}
		return &BulletPointText{Text: string([]rune(rawText)[i:]), Level: i}, nil
	}
	return nil, errors.New("invalid raw text for bullet point")
}

type PlainText struct {
	Text string
}

func (p *PlainText) String() string {
	return p.Text
}

func (p *PlainText) GetText() string {
	return p.Text
}

func (n *NewLineText) String() string {
	return "\n"
}

func (n *NewLineText) GetText() string {
	return "\n"
}

func NewNewLineText() *NewLineText {
	return &NewLineText{}
}

func NewBoldText(text string) *BoldText {
	return (*BoldText)(&BoldText{Text: text, Symbol: '*'})
}

func NewBoldTextFromBracketsText(rawText string) (*BoldText, error) {
	decoratedText, err := newDecoratedText(rawText, '*', "bold")
	return (*BoldText)(decoratedText), err
}

func (t *BoldText) String() string {
	return fmt.Sprintf("[%c %s]", t.Symbol, t.Text)
}

func (t *BoldText) GetText() string {
	return t.Text
}

func NewItalicText(text string) *BoldText {
	return (*BoldText)(&BoldText{Text: text, Symbol: '/'})
}

func NewItalicTextFromBracketsText(rawText string) (*ItalicText, error) {
	decoratedText, err := newDecoratedText(rawText, '/', "italic")
	return (*ItalicText)(decoratedText), err
}

func (t *ItalicText) String() string {
	return fmt.Sprintf("[%c %s]", t.Symbol, t.Text)
}

func (t *ItalicText) GetText() string {
	return t.Text
}

func NewStrikeThroughText(text string) *BoldText {
	return (*BoldText)(&BoldText{Text: text, Symbol: '-'})
}

func NewStrikeThroughTextFromBracketsText(rawText string) (*StrikeThroughText, error) {
	decoratedText, err := newDecoratedText(rawText, '-', "strike through")
	return (*StrikeThroughText)(decoratedText), err
}

func (t *StrikeThroughText) String() string {
	return fmt.Sprintf("[%c %s]", t.Symbol, t.Text)
}

func (t *StrikeThroughText) GetText() string {
	return t.Text
}

func newDecoratedText(rawText string, symbol rune, decorateName string) (*decoratedText, error) {
	trimmedBoldRawText, err := trimDecoratedRawText(rawText, symbol)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("invalid text for %s text", decorateName))
	}
	return &decoratedText{Text: trimmedBoldRawText, Symbol: symbol}, nil
}
