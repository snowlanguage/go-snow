package runtimevalues

import (
	"github.com/snowlanguage/go-snow/position"
)

type RTValue interface {
	ToString() string
	ValueToString() string
	GetType() RTType
	GetValue() interface{}
	GetEnvironment() *Environment
	Add(other RTValue, position position.SEPos) (RTValue, error)
	Subtract(other RTValue, position position.SEPos) (RTValue, error)
	Multiply(other RTValue, position position.SEPos) (RTValue, error)
	Divide(other RTValue, position position.SEPos) (RTValue, error)
}
