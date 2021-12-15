package analyzers

import (
	"github.com/HerbertHe/olex-golang/utils"
	"strings"
)

type LexicalAnalyzer struct {
	Queue []TokenizerResultsQueueType
	Pos int
}

type LexicalAnalyzerResultType = []string

func NewLexicalAnalyzer(queue []TokenizerResultsQueueType) *LexicalAnalyzer {
	return &LexicalAnalyzer{
		Queue: queue,
		Pos: 0,
	}
}

func (l *LexicalAnalyzer) next() TokenizerResultsQueueType {
	return l.Queue[l.Pos + 1]
}

func (l *LexicalAnalyzer) Analyze() []LexicalAnalyzerResultType {
	var _v []LexicalAnalyzerResultType

	for l.Pos < len(l.Queue) {
		c := l.Queue[l.Pos]

		if c[0] == "NewlineToken" {
			l.Pos++
			continue
		}
		
		switch c[0] {
		// Handle `\`
		case "BackslashToken":
			{
				_v = append(_v, []string{"CommentLiteral", "\\"})
				if l.next() != nil && l.next()[0] == "AlphabetToken" {
					_text := l.textLiteralGenerator([]string{"AlphabetToken", "UnicodeToken", "NumberToken"}, nil)
					_v = append(_v, []string{"TextLiteral", _text})
					l.Pos++
				}
				break
			}
		// Handle Special Token
		case "SpecialToken":
			{
				switch c[1] {
				case "[":
					{
						_v = append(_v, []string{"OpenBracketLiteral", "["})
						if l.next() != nil && l.next()[0] == "AlphabetToken" {
							_text := l.textLiteralGenerator([]string{"AlphabetToken", "UnicodeToken", "NumberToken", "SpaceToken"}, nil)
							_v = append(_v, []string{"TextLiteral", _text})
						} else {
							l.Pos++
						}
						break
					}

				case "]":
					{
						_v = append(_v, []string{"CloseBracketLiteral", "]"})
						l.Pos++
						break
					}

				case "{":
					{
						_v = append(_v, []string{"OpenBraceLiteral", "["})
						if l.next() != nil && l.next()[0] == "AlphabetToken" {
							_text := l.textLiteralGenerator([]string{"AlphabetToken", "UnicodeToken", "NumberToken", "SpaceToken"}, nil)
							_v = append(_v, []string{"TextLiteral", _text})
						} else {
							l.Pos++
						}
						break
					}

				case "}":
					{
						_v = append(_v, []string{"CloseBraceLiteral", "}"})
						l.Pos++
						break
					}

				// Handle `$` sign
				case "$":
					{
						if l.next() != nil && l.next()[0] == "SpecialToken" && l.next()[1] == "$" {
							_v = append(_v, []string{"DollarLiteral", "$$"})
							l.Pos += 2
							break
						}

						_v = append(_v, []string{"DollarLiteral", "$"})
						l.Pos++
						break
					}

				default:
					{
						if l.next() != nil && l.next()[0] != "NewlineToken" {
							_text := l.textLiteralGenerator([]string{"AlphabetToken", "NumberToken", "SpaceToken", "UnicodeToken", "SpecialToken"}, []string{"[", "]", "{", "}", "$"})
							_v = append(_v, []string{"TextLiteral", c[1] + _text})
							l.Pos++
							break
						}

						l.Pos++
						break
					}
				}
			}

		case "CommentToken": {
			_v = append(_v, []string{"CommentLiteral", "%"})
			if l.next() != nil && l.next()[0] != "NewlineToken" {
				_text := l.textLiteralGenerator([]string{"AlphabetToken", "BackslashToken", "CommentToken", "NumberToken", "SpaceToken", "SpecialToken", "UnicodeToken"}, nil)
				_v = append(_v, []string{"TextLiteral", strings.TrimSpace(_text)})
			} else {
				l.Pos++
			}
			break
		}

		case "NewlineToken": {
			_v = append(_v, []string{"NewlineLiteral", "\n"})
			l.Pos++
			break
		}

		case "AlphabetToken":
		case "NumberToken":
		case "UnicodeToken":
		case "SpaceToken": {
			AllowTextNext := []string{
				"AlphabetToken",
				"NumberToken",
				"SpaceToken",
				"UnicodeToken",
			}

			if l.next() != nil && utils.ValueInArray(l.next()[0], AllowTextNext) {
				_text := l.textLiteralGenerator(AllowTextNext, nil)
				if len(strings.TrimSpace(c[1] + _text)) > 0 {
					_v = append(_v, []string{"TextLiteral", c[1] + _text})
				}
			}

			l.Pos++
			break
		}

		default:
			l.Pos++
			break
		}
	}
	return _v
}

func (l *LexicalAnalyzer) textLiteralGenerator(AllowNextTokenList []string, except []string) string {
	var _v string
	if except == nil {
		except = []string{}
	}

	for l.next() != nil && l.next()[1] == "*" || (utils.ValueInArray(l.next()[0], AllowNextTokenList) && !utils.ValueInArray(l.next()[1], except)) {
		_v += l.next()[1]
		l.Pos++
	}

	return strings.TrimRight(_v, " ")
}

