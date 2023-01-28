package snow

import (
	"fmt"
)

type RTBool struct {
	Pos         SEPos
	Value       bool
	Environment *Environment
}

func NewRTBool(pos SEPos, value bool, env *Environment) *RTBool {
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

func (rTBool *RTBool) Dot(other Token, position SEPos) (RTValue, error) {
	return nil, NewInvalidAttributeRTError(rTBool, other, position, rTBool.Environment)
}

func (rTBool *RTBool) SetAttribute(other string, value RTValue, position SEPos) (RTValue, error) {
	return nil, NewUnableToAssignAttributeRTError(rTBool, other, value, position, rTBool.Environment)
}

func (rTBool *RTBool) Add(other RTValue, position SEPos) (RTValue, error) {
	return nil, NewValueRTError(
		PLUS,
		rTBool,
		other,
		position,
		rTBool.Environment,
	)
}

func (rTBool *RTBool) Subtract(other RTValue, position SEPos) (RTValue, error) {
	return nil, NewValueRTError(
		DASH,
		rTBool,
		other,
		position,
		rTBool.Environment,
	)
}

func (rTBool *RTBool) Multiply(other RTValue, position SEPos) (RTValue, error) {
	return nil, NewValueRTError(
		STAR,
		rTBool,
		other,
		position,
		rTBool.Environment,
	)
}

func (rTBool *RTBool) Divide(other RTValue, position SEPos) (RTValue, error) {
	return nil, NewValueRTError(
		SLASH,
		rTBool,
		other,
		position,
		rTBool.Environment,
	)
}

func (rTBool *RTBool) Equals(other RTValue, position SEPos) (RTValue, error) {
	switch other.GetType() {
	case RTT_BOOL:
		return NewRTBool(position, rTBool.Value == other.GetValue().(bool), rTBool.Environment), nil
	}

	return NewRTBool(position, false, rTBool.Environment), nil
}

func (rTBool *RTBool) NotEquals(other RTValue, position SEPos) (RTValue, error) {
	switch other.GetType() {
	case RTT_BOOL:
		return NewRTBool(position, rTBool.Value != other.GetValue().(bool), rTBool.Environment), nil
	}

	return NewRTBool(position, true, rTBool.Environment), nil
}

func (rTBool *RTBool) GreaterThan(other RTValue, position SEPos) (RTValue, error) {
	return nil, NewValueRTError(
		GREATER_THAN,
		rTBool,
		other,
		position,
		rTBool.Environment,
	)
}

func (rTBool *RTBool) GreaterThanEquals(other RTValue, position SEPos) (RTValue, error) {
	return nil, NewValueRTError(
		GREATER_THAN_EQUALS,
		rTBool,
		other,
		position,
		rTBool.Environment,
	)
}

func (rTBool *RTBool) LessThan(other RTValue, position SEPos) (RTValue, error) {
	return nil, NewValueRTError(
		LESS_THAN,
		rTBool,
		other,
		position,
		rTBool.Environment,
	)
}

func (rTBool *RTBool) LessThanEquals(other RTValue, position SEPos) (RTValue, error) {
	return nil, NewValueRTError(
		LESS_THAN_EQUALS,
		rTBool,
		other,
		position,
		rTBool.Environment,
	)
}

func (rTBool *RTBool) Not(position SEPos) (RTValue, error) {
	return NewRTBool(position, !rTBool.Value, rTBool.Environment), nil
}

func (rTBool *RTBool) ToBool(position SEPos) (RTValue, error) {
	return NewRTBool(position, rTBool.Value, rTBool.Environment), nil
}

func (rTBool *RTBool) Call(arguments []RTValue, position SEPos, interpreter *Interpreter) (RTValue, error) {
	return nil, NewInvalidCallRTError(rTBool, position, rTBool.Environment)
}
