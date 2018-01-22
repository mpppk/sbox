package scrapbox

import (
	"fmt"
	"log"
)

type Scan struct {
	server  string
	project string
	texts   []Text
}

func (s *Scan) Err(e int) {
	fmt.Printf("\n!!error!!%d\n", e)
}

func (s *Scan) Push(text Text) {
	s.texts = append(s.texts, text)
}

func (s *Scan) PushLinkFromRawText(rawText string) {
	link, err := NewSBLink(fmt.Sprintf("[%s]", rawText), s.server, s.project)
	if err != nil {
		log.Fatal(err)
	}
	s.Push(link)
}

func Parse(s string) ([]Text, error) {
	parser := &Parser{Buffer: s} // 解析対象文字の設定
	parser.Init()                // parser初期化
	parser.s.project = "niboshi"
	parser.s.server = "https://scrapbox.io"

	err := parser.Parse() // 解析
	if err != nil {
		return nil, err
	} else {
		parser.Execute() // アクション処理
	}
	return parser.s.texts, nil
}
