package runtimevalues

import (
	"fmt"

	"github.com/snowlanguage/go-snow/position"
	"github.com/snowlanguage/go-snow/token"
)

type RTBool struct {
	Pos         position.SEPos
	Value       bool
	Environment *Environment
}

func NewRTBool(pos position.SEPos, value bool, env *Environment) *RTBool {
	return &RTBool{
		Pos:         pos,
		Value:       value,
		Environment: env,
	}
}

func (rTBool *RTBool) ToString() string {
	return fmt.Sprintf("(BOOL: %t)", rTBool.Value)
}

func (rTBool *RTBool) ValueToString() string {
	if rTBool.Value {
		return "true"
	} else {
		return "false"
	}
}

func (rTBool *RTBool) GetType() RTType {
	return RTT_BOOL
}

func (rTBool *RTBool) GetValue() interface{} {
	return rTBool.Value
}

func (rTBool *RTBool) GetEnvironment() *Environment {
	return rTBool.Environment
}

func (rTBool *RTBool) Dot(other token.Token, position position.SEPos) (RTValue, error) {
	return nil, NewInvalidAttributeRTError(rTBool, other, position, rTBool.Environment)
}

func (rTBool *RTBool) SetAttribute(other string, value RTValue, position position.SEPos) (RTValue, error) {
	return nil, NewUnableToAssignAttributeError(rTBool, other, value, position, rTBool.Environment)
}

func (rTBool *RTBool) Add(other RTValue, position position.SEPos) (RTValue, error) {
	return nil, NewValueRTError(
		token.PLUS,
		rTBool,
		other,
		position,
		rTBool.Environment,
	)
}

func (rTBool *RTBool) Subtract(other RTValue, position position.SEPos) (RTValue, error) {
	return nil, NewValueRTError(
		token.DASH,
		rTBool,
		other,
		position,
		rTBool.Environment,
	)
}

func (rTBool *RTBool) Multiply(other RTValue, position position.SEPos) (RTValue, error) {
	return nil, NewValueRTError(
		token.STAR,
		rTBool,
		other,
		position,
		rTBool.Environment,
	)
}

func (rTBool *RTBool) Divide(other RTValue, position position.SEPos) (RTValue, error) {
	return nil, NewValueRTError(
		token.SLASH,
		rTBool,
		other,
		position,
		rTBool.Environment,
	)
}

func (rTBool *RTBool) Equals(other RTValue, position position.SEPos) (RTValue, error) {
	switch other.GetType() {
	case RTT_BOOL:
		return NewRTBool(position, rTBool.Value == other.GetValue().(bool), rTBool.Environment), nil
	}

	return NewRTBool(position, false, rTBool.Environment), nil
}

func (rTBool *RTBool) NotEquals(other RTValue, position position.SEPos) (RTValue, error) {
	switch other.GetType() {
	case RTT_BOOL:
		return NewRTBool(position, rTBool.Value != other.GetValue().(bool), rTBool.Environment), nil
	}

	return NewRTBool(position, true, rTBool.Environment), nil
}

func (rTBool *RTBool) GreaterThan(other RTValue, position position.SEPos) (RTValue, error) {
	return nil, NewValueRTError(
		token.GREATER_THAN,
		rTBool,
		other,
		position,
		rTBool.Environment,
	)
}

func (rTBool *RTBool) GreaterThanEquals(other RTValue, position position.SEPos) (RTValue, error) {
	return nil, NewValueRTError(
		token.GREATER_THAN_EQUALS,
		rTBool,
		other,
		position,
		rTBool.Environment,
	)
}

func (rTBool *RTBool) LessThan(other RTValue, position position.SEPos) (RTValue, error) {
	return nil, NewValueRTError(
		token.LESS_THAN,
		rTBool,
		other,
		position,
		rTBool.Environment,
	)
}

func (rTBool *RTBool) LessThanEquals(other RTValue, position position.SEPos) (RTValue, error) {
	return nil, NewValueRTError(
		token.LESS_THAN_EQUALS,
		rTBool,
		other,
		position,
		rTBool.Environment,
	)
}

func (rTBool *RTBool) Not(position position.SEPos) (RTValue, error) {
	return NewRTBool(position, !rTBool.Value, rTBool.Environment), nil
}
