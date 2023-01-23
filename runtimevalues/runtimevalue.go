package runtimevalues

import "github.com/snowlanguage/go-snow/position"

type RTValue interface {
	ToString() string
	GetType() RTType
	GetValue() interface{}
	Add(other RTValue, position position.SEPos) (RTValue, error)
}
