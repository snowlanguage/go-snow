package main

import (
	"fmt"

	"github.com/snowlanguage/go-snow/file"
	"github.com/snowlanguage/go-snow/lexer"
)

func main() {
	f := file.NewFile("<repl>", "123.4.567")
	l := lexer.NewLexer(*f)
	tokens, errors := l.Tokenize()

	if len(errors) != 0 {
		for _, err := range errors {
			fmt.Println(err)
		}
	} else {
		for _, token := range tokens {
			fmt.Print(token.ToString())
		}
	}
}
