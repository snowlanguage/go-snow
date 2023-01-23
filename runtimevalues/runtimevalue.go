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
}
