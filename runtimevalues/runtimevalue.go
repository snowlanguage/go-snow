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
	Equals(other RTValue, position position.SEPos) (RTValue, error)
	NotEquals(other RTValue, position position.SEPos) (RTValue, error)
	GreaterThan(other RTValue, position position.SEPos) (RTValue, error)
	GreaterThanEquals(other RTValue, position position.SEPos) (RTValue, error)
	LessThan(other RTValue, position position.SEPos) (RTValue, error)
	LessThanEquals(other RTValue, position position.SEPos) (RTValue, error)
	Not(position position.SEPos) (RTValue, error)
}
