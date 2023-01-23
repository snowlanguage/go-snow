package runtimevalues

type Environment struct {
	Parent    *Environment
	vars      map[string]RTValue
	StartLine int
	Name      string
}

func NewEnvironment(parent *Environment, name string, startLine int) *Environment {
	return &Environment{
		Parent:    parent,
		StartLine: startLine,
		Name:      name,
	}
}
