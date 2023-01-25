package runtimevalues

import (
	"fmt"

	"github.com/snowlanguage/go-snow/position"
	snowerror "github.com/snowlanguage/go-snow/snowError"
)

type variable struct {
	Value          RTValue
	Constant       bool
	DeclarationPos position.SEPos
}

type Environment struct {
	Parent    *Environment
	vars      map[string]variable
	StartLine int
	Name      string
}

func NewEnvironment(parent *Environment, name string, startLine int) *Environment {
	return &Environment{
		Parent:    parent,
		vars:      make(map[string]variable, 0),
		StartLine: startLine,
		Name:      name,
	}
}

func (environment *Environment) Declare(constant bool, name string, value RTValue, pos position.SEPos) error {
	if v, ok := environment.vars[name]; ok {
		return NewRuntimeError(
			snowerror.VARIABLE_ALREADY_DECLARED_ERROR,
			fmt.Sprintf("a variable with the name '%s' has already declared on line %d", name, v.DeclarationPos.Start.Ln),
			"",
			pos,
			environment,
		)
	}

	environment.vars[name] = variable{
		Value:          value,
		Constant:       constant,
		DeclarationPos: pos,
	}

	return nil
}
