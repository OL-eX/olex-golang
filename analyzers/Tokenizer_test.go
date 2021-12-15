package analyzers

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestTokenizer(t *testing.T) {
	content, err := ioutil.ReadFile("../test/test.tex")
	if err != nil {
		t.Fail()
	}

	token := NewTokenizer(string(content))
	r := token.Analyze()
	for _, v := range r {
		fmt.Println(v)
	}
}
