package analyzers

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestLexicalAnalyzer(t *testing.T) {
	content, err := ioutil.ReadFile("../test/test.tex")
	if err != nil {
		t.Fail()
	}

	token := NewTokenizer(string(content)).Analyze()
	la := NewLexicalAnalyzer(token).Analyze()
	fmt.Println(la)
}