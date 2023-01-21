package main

import (
	"fmt"

	"github.com/snowlanguage/go-snow/file"
	"github.com/snowlanguage/go-snow/lexer"
)

func main() {
	f := file.NewFile("<repl>", "abc+/+-*-")
	l := lexer.NewLexer(*f)
	fmt.Println(l.Tokenize())
}
