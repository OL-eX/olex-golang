package main

import (
	"github.com/HerbertHe/olex-golang/analyzers"
	"syscall/js"
)

func Token(this js.Value, args []js.Value) interface{} {
	return analyzers.Token(args[0].String())
}

func main()  {
	js.Global().Set("Token", js.FuncOf(Token))
}
