package analyzers

import (
	"regexp"
	"strconv"
)

type Tokenizer struct {
	Text []rune
}

type TokenizerResultsQueueType []string

func NewTokenizer(text string) *Tokenizer {
	t := []rune(text)
	return &Tokenizer{Text: t}
}

func (t *Tokenizer) Analyze() (back []TokenizerResultsQueueType) {
	for pos := 0; pos < len(t.Text); pos++ {
		c := string(t.Text[pos])
		p := strconv.Itoa(pos)
		if isAlphabet(c) {
			back = append(back, []string{"AlphabetToken", c, p})
			continue
		}

		if isNumber(c) {
			back = append(back, []string{"NumberToken", c, p})
			continue
		}

		if isSpace(c) {
			back = append(back, []string{"SpaceToken", c, p})
			continue
		}

		if isNewline(c) {
			back = append(back, []string{"NewlineToken", c, p})
			continue
		}

		if isBackslash(c) {
			back = append(back, []string{"BackslashToken", c, p})
			continue
		}

		if isComment(c) {
			back = append(back, []string{"CommentToken", c, p})
			continue
		}

		if isSpecial(c) {
			back = append(back, []string{"SpecialToken", c, p})
			continue
		}

		back = append(back, []string{"UnicodeToken", c, p})
	}

	return back
}

func isAlphabet(c string) bool {
	f, _ := regexp.MatchString("[a-zA-Z]", c)
	return f
}

func isNumber(c string) bool {
	f, _ := regexp.MatchString("[0-9]", c)
	return f
}

func isNewline(c string) bool {
	return c == "\n"
}

func isBackslash(c string) bool {
	return c == "\\"
}

func isSpace(c string) bool {
	return c == " "
}

func isComment(c string) bool {
	return c == "%"
}

func isSpecial(c string) bool {
	f, _ := regexp.MatchString("[!-$&-/:-@\\[-`{-~]", c)
	return f
}