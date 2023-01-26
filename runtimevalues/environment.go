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
			fmt.Sprintf("a variable with the name of '%s' has already declared on line %d", name, v.DeclarationPos.Start.Ln),
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

func (environment *Environment) Get(name string, pos position.SEPos) (RTValue, error) {
	if _, ok := environment.vars[name]; !ok {
		if environment.Parent != nil {
			return environment.Parent.Get(name, pos)
		}

		return nil, NewRuntimeError(
			snowerror.UNDEFINED_VARIABLE_ERROR,
			fmt.Sprintf("a variable with the name of '%s' could not be found", name),
			"",
			pos,
			environment,
		)
	}

	return environment.vars[name].Value, nil
}

func (environment *Environment) Set(name string, value RTValue, pos position.SEPos) (RTValue, error) {
	if _, ok := environment.vars[name]; !ok {
		if environment.Parent != nil {
			return environment.Parent.Get(name, pos)
		}

		return nil, NewRuntimeError(
			snowerror.UNDEFINED_VARIABLE_ERROR,
			fmt.Sprintf("a variable with the name of '%s' could not be found", name),
			"",
			pos,
			environment,
		)
	}

	v := environment.vars[name]
	if v.Constant {
		return nil, NewRuntimeError(
			snowerror.CONSTANT_VARIABLE_ASSIGNMENT_ERROR,
			fmt.Sprintf("the variable '%s' is a constant and can therefor not be assigned to", name),
			name,
			pos,
			environment,
		)
	}

	newVar := variable{
		Value:          value,
		Constant:       false,
		DeclarationPos: v.DeclarationPos,
	}

	environment.vars[name] = newVar

	return value, nil
}
