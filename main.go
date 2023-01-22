package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/snowlanguage/go-snow/file"
	"github.com/snowlanguage/go-snow/lexer"
	"github.com/snowlanguage/go-snow/token"
)

func logErrors(errors []error) {
	for index, err := range errors {
		if index >= 5 {
			fmt.Printf("Showing 5/%d errors\n", len(errors))
			return
		}

		fmt.Println(err)
	}
}

func runFile(file string) {
	code, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	_, errors := run(file, string(code))

	if len(errors) != 0 {
		logErrors(errors)
	}
}

func runRepl() {
	input := bufio.NewScanner(os.Stdin)

	fmt.Print("> ")
	for input.Scan() {
		code := input.Text()

		if code == "" {
			continue
		}

		tokens, errors := run("<repl>", code)

		if len(errors) != 0 {
			logErrors(errors)
		} else {
			for _, tok := range tokens {
				fmt.Print(tok.ToString() + " ")
			}

			fmt.Println()
		}

		fmt.Print("> ")
	}
}

func run(filename string, code string) ([]token.Token, []error) {
	f := file.NewFile(filename, code)
	l := lexer.NewLexer(*f)
	return l.Tokenize()
}

func main() {
	if len(os.Args) == 2 {
		runFile(os.Args[1])
	} else {
		runRepl()
	}
}
