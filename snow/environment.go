package snow

import (
	"fmt"
)

type variable struct {
	Value          RTValue
	Constant       bool
	DeclarationPos SEPos
}

type Environment struct {
	Parent    *Environment
	vars      map[string]variable
	StartLine int
	FileName  string
	Name      string
	IsFile    bool
}

func NewEnvironment(parent *Environment, name string, startLine int, fileName string, isFile bool) *Environment {
	return &Environment{
		Parent:    parent,
		vars:      make(map[string]variable, 0),
		StartLine: startLine,
		FileName:  fileName,
		Name:      name,
		IsFile:    isFile,
	}
}

func (environment *Environment) Declare(constant bool, name string, value RTValue, pos SEPos) error {
	if v, ok := environment.vars[name]; ok {
		return NewRuntimeError(
			VARIABLE_ALREADY_DECLARED_ERROR,
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

func (environment *Environment) Get(name string, pos SEPos, env *Environment) (RTValue, error) {
	if _, ok := environment.vars[name]; !ok {
		if environment.Parent != nil {
			return environment.Parent.Get(name, pos, env)
		}

		return nil, NewRuntimeError(
			UNDEFINED_VARIABLE_ERROR,
			fmt.Sprintf("a variable with the name of '%s' could not be found", name),
			"",
			pos,
			env,
		)
	}

	return environment.vars[name].Value, nil
}

func (environment *Environment) Set(name string, value RTValue, env *Environment, pos SEPos) (RTValue, error) {
	if _, ok := environment.vars[name]; !ok {
		if environment.Parent != nil {
			return environment.Parent.Set(name, value, env, pos)
		}

		return nil, NewRuntimeError(
			UNDEFINED_VARIABLE_ERROR,
			fmt.Sprintf("a variable with the name of '%s' could not be found", name),
			"",
			pos,
			env,
		)
	}

	v := environment.vars[name]
	if v.Constant {
		return nil, NewRuntimeError(
			CONSTANT_VARIABLE_ASSIGNMENT_ERROR,
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
