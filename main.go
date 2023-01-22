package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/snowlanguage/go-snow/file"
	"github.com/snowlanguage/go-snow/lexer"
	parsevals "github.com/snowlanguage/go-snow/parseVals"
	"github.com/snowlanguage/go-snow/parser"
)

func logErrors(errors []error) {
	for index, err := range errors {
		if index >= 5 {
			fmt.Printf("Showing 5/%d errors\n", len(errors))
			return
		}

		fmt.Println(err.Error())
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

			if len(tokens) != 0 {
				fmt.Println()
			}
		}

		fmt.Print("> ")
	}
}

func run(filename string, code string) ([]parsevals.Stmt, []error) {
	f := file.NewFile(filename, code)
	l := lexer.NewLexer(f)

	t, err := l.Tokenize()

	if len(err) != 0 {
		return nil, err
	}

	for _, tok := range t {
		fmt.Println("token", tok.ToString())
	}

	p := parser.NewParser(t, f)
	s, err2 := p.Parse()

	if err2 != nil {
		var errArray = []error{err2}
		return nil, errArray
	}

	return s, nil
}

func main() {
	if len(os.Args) == 2 {
		runFile(os.Args[1])
	} else {
		runRepl()
	}
}
