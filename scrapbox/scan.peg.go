package scrapbox

//go:generate peg scrapbox/scan.peg

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

const endSymbol rune = 1114112

/* The rule types inferred from the grammar are below. */
type pegRule uint8

const (
	ruleUnknown pegRule = iota
	ruleroot
	ruleEOT
	ruleexpression
	ruleliteral
	rulebracketText
	ruleinnerDecoratedText
	ruleplainTextLine
	rulePegText
	ruleAction0
	ruleAction1
	ruleAction2
	ruleAction3
	ruleAction4
	ruleAction5
	ruleAction6
	ruleAction7
	ruleAction8
	ruleAction9
	ruleAction10
)

var rul3s = [...]string{
	"Unknown",
	"root",
	"EOT",
	"expression",
	"literal",
	"bracketText",
	"innerDecoratedText",
	"plainTextLine",
	"PegText",
	"Action0",
	"Action1",
	"Action2",
	"Action3",
	"Action4",
	"Action5",
	"Action6",
	"Action7",
	"Action8",
	"Action9",
	"Action10",
}

type token32 struct {
	pegRule
	begin, end uint32
}

func (t *token32) String() string {
	return fmt.Sprintf("\x1B[34m%v\x1B[m %v %v", rul3s[t.pegRule], t.begin, t.end)
}

type node32 struct {
	token32
	up, next *node32
}

func (node *node32) print(pretty bool, buffer string) {
	var print func(node *node32, depth int)
	print = func(node *node32, depth int) {
		for node != nil {
			for c := 0; c < depth; c++ {
				fmt.Printf(" ")
			}
			rule := rul3s[node.pegRule]
			quote := strconv.Quote(string(([]rune(buffer)[node.begin:node.end])))
			if !pretty {
				fmt.Printf("%v %v\n", rule, quote)
			} else {
				fmt.Printf("\x1B[34m%v\x1B[m %v\n", rule, quote)
			}
			if node.up != nil {
				print(node.up, depth+1)
			}
			node = node.next
		}
	}
	print(node, 0)
}

func (node *node32) Print(buffer string) {
	node.print(false, buffer)
}

func (node *node32) PrettyPrint(buffer string) {
	node.print(true, buffer)
}

type tokens32 struct {
	tree []token32
}

func (t *tokens32) Trim(length uint32) {
	t.tree = t.tree[:length]
}

func (t *tokens32) Print() {
	for _, token := range t.tree {
		fmt.Println(token.String())
	}
}

func (t *tokens32) AST() *node32 {
	type element struct {
		node *node32
		down *element
	}
	tokens := t.Tokens()
	var stack *element
	for _, token := range tokens {
		if token.begin == token.end {
			continue
		}
		node := &node32{token32: token}
		for stack != nil && stack.node.begin >= token.begin && stack.node.end <= token.end {
			stack.node.next = node.up
			node.up = stack.node
			stack = stack.down
		}
		stack = &element{node: node, down: stack}
	}
	if stack != nil {
		return stack.node
	}
	return nil
}

func (t *tokens32) PrintSyntaxTree(buffer string) {
	t.AST().Print(buffer)
}

func (t *tokens32) PrettyPrintSyntaxTree(buffer string) {
	t.AST().PrettyPrint(buffer)
}

func (t *tokens32) Add(rule pegRule, begin, end, index uint32) {
	if tree := t.tree; int(index) >= len(tree) {
		expanded := make([]token32, 2*len(tree))
		copy(expanded, tree)
		t.tree = expanded
	}
	t.tree[index] = token32{
		pegRule: rule,
		begin:   begin,
		end:     end,
	}
}

func (t *tokens32) Tokens() []token32 {
	return t.tree
}

type Parser struct {
	s Scan // parserが自動生成するフィールド変数と区別するために
	//   敢えて埋め込みを行っていない。

	Buffer string
	buffer []rune
	rules  [20]func() bool
	parse  func(rule ...int) error
	reset  func()
	Pretty bool
	tokens32
}

func (p *Parser) Parse(rule ...int) error {
	return p.parse(rule...)
}

func (p *Parser) Reset() {
	p.reset()
}

type textPosition struct {
	line, symbol int
}

type textPositionMap map[int]textPosition

func translatePositions(buffer []rune, positions []int) textPositionMap {
	length, translations, j, line, symbol := len(positions), make(textPositionMap, len(positions)), 0, 1, 0
	sort.Ints(positions)

search:
	for i, c := range buffer {
		if c == '\n' {
			line, symbol = line+1, 0
		} else {
			symbol++
		}
		if i == positions[j] {
			translations[positions[j]] = textPosition{line, symbol}
			for j++; j < length; j++ {
				if i != positions[j] {
					continue search
				}
			}
			break search
		}
	}

	return translations
}

type parseError struct {
	p   *Parser
	max token32
}

func (e *parseError) Error() string {
	tokens, error := []token32{e.max}, "\n"
	positions, p := make([]int, 2*len(tokens)), 0
	for _, token := range tokens {
		positions[p], p = int(token.begin), p+1
		positions[p], p = int(token.end), p+1
	}
	translations := translatePositions(e.p.buffer, positions)
	format := "parse error near %v (line %v symbol %v - line %v symbol %v):\n%v\n"
	if e.p.Pretty {
		format = "parse error near \x1B[34m%v\x1B[m (line %v symbol %v - line %v symbol %v):\n%v\n"
	}
	for _, token := range tokens {
		begin, end := int(token.begin), int(token.end)
		error += fmt.Sprintf(format,
			rul3s[token.pegRule],
			translations[begin].line, translations[begin].symbol,
			translations[end].line, translations[end].symbol,
			strconv.Quote(string(e.p.buffer[begin:end])))
	}

	return error
}

func (p *Parser) PrintSyntaxTree() {
	if p.Pretty {
		p.tokens32.PrettyPrintSyntaxTree(p.Buffer)
	} else {
		p.tokens32.PrintSyntaxTree(p.Buffer)
	}
}

func (p *Parser) Execute() {
	buffer, _buffer, text, begin, end := p.Buffer, p.buffer, "", 0, 0
	for _, token := range p.Tokens() {
		switch token.pegRule {

		case rulePegText:
			begin, end = int(token.begin), int(token.end)
			text = string(_buffer[begin:end])

		case ruleAction0:
			p.s.Err(begin)
		case ruleAction1:
			p.s.Err(begin)
		case ruleAction2:

			b, _ := NewBulletPointText(text)
			p.s.Push(b)

		case ruleAction3:

			p.s.Push(&PlainText{Text: text})

		case ruleAction4:

			p.s.Push(&PlainText{Text: text})

		case ruleAction5:

			p.s.Push(NewNewLineText())

		case ruleAction6:

			p.s.Push(&PlainText{Text: text})

		case ruleAction7:

			p.s.Push(NewBoldText(text))

		case ruleAction8:

			p.s.Push(NewItalicText(text))

		case ruleAction9:

			p.s.Push(NewStrikeThroughText(text))

		case ruleAction10:

			p.s.PushLinkFromRawText(text)

		}
	}
	_, _, _, _, _ = buffer, _buffer, text, begin, end
}

func (p *Parser) Init() {
	var (
		max                  token32
		position, tokenIndex uint32
		buffer               []rune
	)
	p.reset = func() {
		max = token32{}
		position, tokenIndex = 0, 0

		p.buffer = []rune(p.Buffer)
		if len(p.buffer) == 0 || p.buffer[len(p.buffer)-1] != endSymbol {
			p.buffer = append(p.buffer, endSymbol)
		}
		buffer = p.buffer
	}
	p.reset()

	_rules := p.rules
	tree := tokens32{tree: make([]token32, math.MaxInt16)}
	p.parse = func(rule ...int) error {
		r := 1
		if len(rule) > 0 {
			r = rule[0]
		}
		matches := p.rules[r]()
		p.tokens32 = tree
		if matches {
			p.Trim(tokenIndex)
			return nil
		}
		return &parseError{p, max}
	}

	add := func(rule pegRule, begin uint32) {
		tree.Add(rule, begin, position, tokenIndex)
		tokenIndex++
		if begin != position && position > max.end {
			max = token32{rule, begin, position}
		}
	}

	matchDot := func() bool {
		if buffer[position] != endSymbol {
			position++
			return true
		}
		return false
	}

	/*matchChar := func(c byte) bool {
		if buffer[position] == c {
			position++
			return true
		}
		return false
	}*/

	_rules = [...]func() bool{
		nil,
		/* 0 root <- <((expression EOT) / (expression <.+> Action0 EOT) / (<.+> Action1 EOT))> */
		func() bool {
			position0, tokenIndex0 := position, tokenIndex
			{
				position1 := position
				{
					position2, tokenIndex2 := position, tokenIndex
					if !_rules[ruleexpression]() {
						goto l3
					}
					if !_rules[ruleEOT]() {
						goto l3
					}
					goto l2
				l3:
					position, tokenIndex = position2, tokenIndex2
					if !_rules[ruleexpression]() {
						goto l4
					}
					{
						position5 := position
						if !matchDot() {
							goto l4
						}
					l6:
						{
							position7, tokenIndex7 := position, tokenIndex
							if !matchDot() {
								goto l7
							}
							goto l6
						l7:
							position, tokenIndex = position7, tokenIndex7
						}
						add(rulePegText, position5)
					}
					if !_rules[ruleAction0]() {
						goto l4
					}
					if !_rules[ruleEOT]() {
						goto l4
					}
					goto l2
				l4:
					position, tokenIndex = position2, tokenIndex2
					{
						position8 := position
						if !matchDot() {
							goto l0
						}
					l9:
						{
							position10, tokenIndex10 := position, tokenIndex
							if !matchDot() {
								goto l10
							}
							goto l9
						l10:
							position, tokenIndex = position10, tokenIndex10
						}
						add(rulePegText, position8)
					}
					if !_rules[ruleAction1]() {
						goto l0
					}
					if !_rules[ruleEOT]() {
						goto l0
					}
				}
			l2:
				add(ruleroot, position1)
			}
			return true
		l0:
			position, tokenIndex = position0, tokenIndex0
			return false
		},
		/* 1 EOT <- <!.> */
		func() bool {
			position11, tokenIndex11 := position, tokenIndex
			{
				position12 := position
				{
					position13, tokenIndex13 := position, tokenIndex
					if !matchDot() {
						goto l13
					}
					goto l11
				l13:
					position, tokenIndex = position13, tokenIndex13
				}
				add(ruleEOT, position12)
			}
			return true
		l11:
			position, tokenIndex = position11, tokenIndex11
			return false
		},
		/* 2 expression <- <literal+> */
		func() bool {
			position14, tokenIndex14 := position, tokenIndex
			{
				position15 := position
				if !_rules[ruleliteral]() {
					goto l14
				}
			l16:
				{
					position17, tokenIndex17 := position, tokenIndex
					if !_rules[ruleliteral]() {
						goto l17
					}
					goto l16
				l17:
					position, tokenIndex = position17, tokenIndex17
				}
				add(ruleexpression, position15)
			}
			return true
		l14:
			position, tokenIndex = position14, tokenIndex14
			return false
		},
		/* 3 literal <- <(bracketText / (<(' '+ plainTextLine)> Action2) / (<(!('[' / '\n') .)+> Action3) / (<(!(']' / '\n') .)+> Action4) / ('\n' Action5))> */
		func() bool {
			position18, tokenIndex18 := position, tokenIndex
			{
				position19 := position
				{
					position20, tokenIndex20 := position, tokenIndex
					if !_rules[rulebracketText]() {
						goto l21
					}
					goto l20
				l21:
					position, tokenIndex = position20, tokenIndex20
					{
						position23 := position
						if buffer[position] != rune(' ') {
							goto l22
						}
						position++
					l24:
						{
							position25, tokenIndex25 := position, tokenIndex
							if buffer[position] != rune(' ') {
								goto l25
							}
							position++
							goto l24
						l25:
							position, tokenIndex = position25, tokenIndex25
						}
						if !_rules[ruleplainTextLine]() {
							goto l22
						}
						add(rulePegText, position23)
					}
					if !_rules[ruleAction2]() {
						goto l22
					}
					goto l20
				l22:
					position, tokenIndex = position20, tokenIndex20
					{
						position27 := position
						{
							position30, tokenIndex30 := position, tokenIndex
							{
								position31, tokenIndex31 := position, tokenIndex
								if buffer[position] != rune('[') {
									goto l32
								}
								position++
								goto l31
							l32:
								position, tokenIndex = position31, tokenIndex31
								if buffer[position] != rune('\n') {
									goto l30
								}
								position++
							}
						l31:
							goto l26
						l30:
							position, tokenIndex = position30, tokenIndex30
						}
						if !matchDot() {
							goto l26
						}
					l28:
						{
							position29, tokenIndex29 := position, tokenIndex
							{
								position33, tokenIndex33 := position, tokenIndex
								{
									position34, tokenIndex34 := position, tokenIndex
									if buffer[position] != rune('[') {
										goto l35
									}
									position++
									goto l34
								l35:
									position, tokenIndex = position34, tokenIndex34
									if buffer[position] != rune('\n') {
										goto l33
									}
									position++
								}
							l34:
								goto l29
							l33:
								position, tokenIndex = position33, tokenIndex33
							}
							if !matchDot() {
								goto l29
							}
							goto l28
						l29:
							position, tokenIndex = position29, tokenIndex29
						}
						add(rulePegText, position27)
					}
					if !_rules[ruleAction3]() {
						goto l26
					}
					goto l20
				l26:
					position, tokenIndex = position20, tokenIndex20
					{
						position37 := position
						{
							position40, tokenIndex40 := position, tokenIndex
							{
								position41, tokenIndex41 := position, tokenIndex
								if buffer[position] != rune(']') {
									goto l42
								}
								position++
								goto l41
							l42:
								position, tokenIndex = position41, tokenIndex41
								if buffer[position] != rune('\n') {
									goto l40
								}
								position++
							}
						l41:
							goto l36
						l40:
							position, tokenIndex = position40, tokenIndex40
						}
						if !matchDot() {
							goto l36
						}
					l38:
						{
							position39, tokenIndex39 := position, tokenIndex
							{
								position43, tokenIndex43 := position, tokenIndex
								{
									position44, tokenIndex44 := position, tokenIndex
									if buffer[position] != rune(']') {
										goto l45
									}
									position++
									goto l44
								l45:
									position, tokenIndex = position44, tokenIndex44
									if buffer[position] != rune('\n') {
										goto l43
									}
									position++
								}
							l44:
								goto l39
							l43:
								position, tokenIndex = position43, tokenIndex43
							}
							if !matchDot() {
								goto l39
							}
							goto l38
						l39:
							position, tokenIndex = position39, tokenIndex39
						}
						add(rulePegText, position37)
					}
					if !_rules[ruleAction4]() {
						goto l36
					}
					goto l20
				l36:
					position, tokenIndex = position20, tokenIndex20
					if buffer[position] != rune('\n') {
						goto l18
					}
					position++
					if !_rules[ruleAction5]() {
						goto l18
					}
				}
			l20:
				add(ruleliteral, position19)
			}
			return true
		l18:
			position, tokenIndex = position18, tokenIndex18
			return false
		},
		/* 4 bracketText <- <(('[' innerDecoratedText ']') / (<('[' plainTextLine)> Action6))> */
		func() bool {
			position46, tokenIndex46 := position, tokenIndex
			{
				position47 := position
				{
					position48, tokenIndex48 := position, tokenIndex
					if buffer[position] != rune('[') {
						goto l49
					}
					position++
					if !_rules[ruleinnerDecoratedText]() {
						goto l49
					}
					if buffer[position] != rune(']') {
						goto l49
					}
					position++
					goto l48
				l49:
					position, tokenIndex = position48, tokenIndex48
					{
						position50 := position
						if buffer[position] != rune('[') {
							goto l46
						}
						position++
						if !_rules[ruleplainTextLine]() {
							goto l46
						}
						add(rulePegText, position50)
					}
					if !_rules[ruleAction6]() {
						goto l46
					}
				}
			l48:
				add(rulebracketText, position47)
			}
			return true
		l46:
			position, tokenIndex = position46, tokenIndex46
			return false
		},
		/* 5 innerDecoratedText <- <(('*' ' ' <plainTextLine> Action7) / ('/' ' ' <plainTextLine> Action8) / ('-' ' ' <plainTextLine> Action9) / (<plainTextLine> Action10))> */
		func() bool {
			position51, tokenIndex51 := position, tokenIndex
			{
				position52 := position
				{
					position53, tokenIndex53 := position, tokenIndex
					if buffer[position] != rune('*') {
						goto l54
					}
					position++
					if buffer[position] != rune(' ') {
						goto l54
					}
					position++
					{
						position55 := position
						if !_rules[ruleplainTextLine]() {
							goto l54
						}
						add(rulePegText, position55)
					}
					if !_rules[ruleAction7]() {
						goto l54
					}
					goto l53
				l54:
					position, tokenIndex = position53, tokenIndex53
					if buffer[position] != rune('/') {
						goto l56
					}
					position++
					if buffer[position] != rune(' ') {
						goto l56
					}
					position++
					{
						position57 := position
						if !_rules[ruleplainTextLine]() {
							goto l56
						}
						add(rulePegText, position57)
					}
					if !_rules[ruleAction8]() {
						goto l56
					}
					goto l53
				l56:
					position, tokenIndex = position53, tokenIndex53
					if buffer[position] != rune('-') {
						goto l58
					}
					position++
					if buffer[position] != rune(' ') {
						goto l58
					}
					position++
					{
						position59 := position
						if !_rules[ruleplainTextLine]() {
							goto l58
						}
						add(rulePegText, position59)
					}
					if !_rules[ruleAction9]() {
						goto l58
					}
					goto l53
				l58:
					position, tokenIndex = position53, tokenIndex53
					{
						position60 := position
						if !_rules[ruleplainTextLine]() {
							goto l51
						}
						add(rulePegText, position60)
					}
					if !_rules[ruleAction10]() {
						goto l51
					}
				}
			l53:
				add(ruleinnerDecoratedText, position52)
			}
			return true
		l51:
			position, tokenIndex = position51, tokenIndex51
			return false
		},
		/* 6 plainTextLine <- <(!(']' / '[' / '\n') .)+> */
		func() bool {
			position61, tokenIndex61 := position, tokenIndex
			{
				position62 := position
				{
					position65, tokenIndex65 := position, tokenIndex
					{
						position66, tokenIndex66 := position, tokenIndex
						if buffer[position] != rune(']') {
							goto l67
						}
						position++
						goto l66
					l67:
						position, tokenIndex = position66, tokenIndex66
						if buffer[position] != rune('[') {
							goto l68
						}
						position++
						goto l66
					l68:
						position, tokenIndex = position66, tokenIndex66
						if buffer[position] != rune('\n') {
							goto l65
						}
						position++
					}
				l66:
					goto l61
				l65:
					position, tokenIndex = position65, tokenIndex65
				}
				if !matchDot() {
					goto l61
				}
			l63:
				{
					position64, tokenIndex64 := position, tokenIndex
					{
						position69, tokenIndex69 := position, tokenIndex
						{
							position70, tokenIndex70 := position, tokenIndex
							if buffer[position] != rune(']') {
								goto l71
							}
							position++
							goto l70
						l71:
							position, tokenIndex = position70, tokenIndex70
							if buffer[position] != rune('[') {
								goto l72
							}
							position++
							goto l70
						l72:
							position, tokenIndex = position70, tokenIndex70
							if buffer[position] != rune('\n') {
								goto l69
							}
							position++
						}
					l70:
						goto l64
					l69:
						position, tokenIndex = position69, tokenIndex69
					}
					if !matchDot() {
						goto l64
					}
					goto l63
				l64:
					position, tokenIndex = position64, tokenIndex64
				}
				add(ruleplainTextLine, position62)
			}
			return true
		l61:
			position, tokenIndex = position61, tokenIndex61
			return false
		},
		nil,
		/* 9 Action0 <- <{p.s.Err(begin)}> */
		func() bool {
			{
				add(ruleAction0, position)
			}
			return true
		},
		/* 10 Action1 <- <{p.s.Err(begin)}> */
		func() bool {
			{
				add(ruleAction1, position)
			}
			return true
		},
		/* 11 Action2 <- <{
		    b, _ := NewBulletPointText(text)
		    p.s.Push(b)
		}> */
		func() bool {
			{
				add(ruleAction2, position)
			}
			return true
		},
		/* 12 Action3 <- <{
		    p.s.Push(&PlainText{Text: text})
		}> */
		func() bool {
			{
				add(ruleAction3, position)
			}
			return true
		},
		/* 13 Action4 <- <{
		    p.s.Push(&PlainText{Text: text})
		}> */
		func() bool {
			{
				add(ruleAction4, position)
			}
			return true
		},
		/* 14 Action5 <- <{
		    p.s.Push(NewNewLineText())
		}> */
		func() bool {
			{
				add(ruleAction5, position)
			}
			return true
		},
		/* 15 Action6 <- <{
		    p.s.Push(&PlainText{Text: text})
		}> */
		func() bool {
			{
				add(ruleAction6, position)
			}
			return true
		},
		/* 16 Action7 <- <{
		    p.s.Push(NewBoldText(text))
		}> */
		func() bool {
			{
				add(ruleAction7, position)
			}
			return true
		},
		/* 17 Action8 <- <{
		    p.s.Push(NewItalicText(text))
		}> */
		func() bool {
			{
				add(ruleAction8, position)
			}
			return true
		},
		/* 18 Action9 <- <{
		    p.s.Push(NewStrikeThroughText(text))
		}> */
		func() bool {
			{
				add(ruleAction9, position)
			}
			return true
		},
		/* 19 Action10 <- <{
		    p.s.PushLinkFromRawText(text)
		}> */
		func() bool {
			{
				add(ruleAction10, position)
			}
			return true
		},
	}
	p.rules = _rules
}
