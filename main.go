package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/snowlanguage/go-snow/file"
	"github.com/snowlanguage/go-snow/interpreter"
	"github.com/snowlanguage/go-snow/lexer"
	"github.com/snowlanguage/go-snow/parser"
	"github.com/snowlanguage/go-snow/runtimevalues"
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

	env := runtimevalues.NewEnvironment(nil, file, 1)

	vals, errors := run(file, string(code), env)

	if len(errors) != 0 {
		logErrors(errors)
	}

	for _, val := range vals {
		fmt.Println(val.ValueToString())
	}
}

func runRepl() {
	input := bufio.NewReader(os.Stdin)

	env := runtimevalues.NewEnvironment(nil, "<repl>", 1)

	code := "a"
	fmt.Print("> ")
	for code != "" {
		codeByte, _, err := input.ReadLine()
		if err != nil {
			log.Fatalln(err)
		}

		code = string(codeByte)

		if code == "" {
			continue
		}

		values, errors := run("<repl>", code, env)

		if len(errors) != 0 {
			logErrors(errors)
		} else {
			for _, val := range values {
				fmt.Print(val.ValueToString() + " ")
			}

			if len(values) != 0 {
				fmt.Println()
			}
		}

		fmt.Print("> ")
	}
}

func run(filename string, code string, e *runtimevalues.Environment) ([]runtimevalues.RTValue, []error) {
	f := file.NewFile(filename, code)
	l := lexer.NewLexer(f)

	// fmt.Println("Tokenizing")

	t, err := l.Tokenize()

	if len(err) != 0 {
		return nil, err
	}

	for _, tok := range t {
		fmt.Println("token", tok.ToString())
	}

	p := parser.NewParser(t, f)

	// fmt.Println("Parsing")

	s, err2 := p.Parse()

	if err2 != nil {
		var errArray = []error{err2}
		return nil, errArray
	}

	for _, stmt := range s {
		fmt.Println("parser", stmt.ToString())
	}

	i := interpreter.NewInterpreter(s, f, e)

	// fmt.Println("Interpreting")

	v, err3 := i.Interpret()

	if err3 != nil {
		var errArray = []error{err3}
		return nil, errArray
	}

	return v, nil
}

func main() {
	if len(os.Args) == 2 {
		runFile(os.Args[1])
	} else {
		runRepl()
	}
}
