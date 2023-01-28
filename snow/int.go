package snow

import (
	"fmt"
)

type RTInt struct {
	Pos         SEPos
	Value       int
	Environment *Environment
}

func NewRTInt(pos SEPos, value int, env *Environment) *RTInt {
	return &RTInt{
		Pos:         pos,
		Value:       value,
		Environment: env,
	}
}

func (rTInt *RTInt) ToString() string {
	return fmt.Sprintf("(INT: %d)", rTInt.Value)
}

func (rTInt *RTInt) ValueToString() string {
	return fmt.Sprintf("%d", rTInt.Value)
}

func (rTInt *RTInt) GetType() RTType {
	return RTT_INT
}

func (rTInt *RTInt) GetValue() interface{} {
	return rTInt.Value
}

func (rTInt *RTInt) GetEnvironment() *Environment {
	return rTInt.Environment
}

func (rTInt *RTInt) Dot(other Token, position SEPos) (RTValue, error) {
	return nil, NewInvalidAttributeRTError(rTInt, other, position, rTInt.Environment)
}

func (rTInt *RTInt) SetAttribute(other string, value RTValue, position SEPos) (RTValue, error) {
	return nil, NewUnableToAssignAttributeRTError(rTInt, other, value, position, rTInt.Environment)
}

func (rTInt *RTInt) Add(other RTValue, position SEPos) (RTValue, error) {
	switch other.GetType() {
	case RTT_INT:

		return NewRTInt(position, rTInt.Value+other.GetValue().(int), rTInt.Environment), nil
	case RTT_FLOAT:
		return NewRTFloat(position, float64(rTInt.Value)+other.GetValue().(float64), rTInt.Environment), nil
	}

	return nil, NewValueRTError(
		PLUS,
		rTInt,
		other,
		position,
		rTInt.Environment,
	)
}

func (rTInt *RTInt) Subtract(other RTValue, position SEPos) (RTValue, error) {
	switch other.GetType() {
	case RTT_INT:
		return NewRTInt(position, rTInt.Value-other.GetValue().(int), rTInt.Environment), nil
	case RTT_FLOAT:
		return NewRTFloat(position, float64(rTInt.Value)-other.GetValue().(float64), rTInt.Environment), nil
	}

	return nil, NewValueRTError(
		PLUS,
		rTInt,
		other,
		position,
		rTInt.Environment,
	)
}

func (rTInt *RTInt) Multiply(other RTValue, position SEPos) (RTValue, error) {
	switch other.GetType() {
	case RTT_INT:
		return NewRTInt(position, rTInt.Value*other.GetValue().(int), rTInt.Environment), nil
	case RTT_FLOAT:
		return NewRTFloat(position, float64(rTInt.Value)*other.GetValue().(float64), rTInt.Environment), nil
	}

	return nil, NewValueRTError(
		PLUS,
		rTInt,
		other,
		position,
		rTInt.Environment,
	)
}

func (rTInt *RTInt) Divide(other RTValue, position SEPos) (RTValue, error) {
	switch other.GetType() {
	case RTT_INT:
		if other.GetValue().(int) == 0 {
			return nil, NewDivisionByZeroRTError(rTInt, other, position, rTInt.Environment)
		}

		return NewRTFloat(position, float64(rTInt.Value)/float64(other.GetValue().(int)), rTInt.Environment), nil
	case RTT_FLOAT:
		if other.GetValue().(float64) == 0 {
			return nil, NewDivisionByZeroRTError(rTInt, other, position, rTInt.Environment)
		}

		return NewRTFloat(position, float64(rTInt.Value)/other.GetValue().(float64), rTInt.Environment), nil
	}

	return nil, NewValueRTError(
		PLUS,
		rTInt,
		other,
		position,
		rTInt.Environment,
	)
}

func (rTInt *RTInt) Equals(other RTValue, position SEPos) (RTValue, error) {
	switch other.GetType() {
	case RTT_INT:
		return NewRTBool(position, rTInt.Value == other.GetValue().(int), rTInt.Environment), nil
	case RTT_FLOAT:
		return NewRTBool(position, float64(rTInt.Value) == other.GetValue().(float64), rTInt.Environment), nil
	}

	return NewRTBool(position, false, rTInt.Environment), nil
}

func (rTInt *RTInt) NotEquals(other RTValue, position SEPos) (RTValue, error) {
	switch other.GetType() {
	case RTT_INT:
		return NewRTBool(position, rTInt.Value != other.GetValue().(int), rTInt.Environment), nil
	case RTT_FLOAT:
		return NewRTBool(position, float64(rTInt.Value) != other.GetValue().(float64), rTInt.Environment), nil
	}

	return NewRTBool(position, true, rTInt.Environment), nil
}

func (rTInt *RTInt) GreaterThan(other RTValue, position SEPos) (RTValue, error) {
	switch other.GetType() {
	case RTT_INT:
		return NewRTBool(position, rTInt.Value > other.GetValue().(int), rTInt.Environment), nil
	case RTT_FLOAT:
		return NewRTBool(position, float64(rTInt.Value) > other.GetValue().(float64), rTInt.Environment), nil
	}

	return nil, NewValueRTError(
		GREATER_THAN,
		rTInt,
		other,
		position,
		rTInt.Environment,
	)
}

func (rTInt *RTInt) GreaterThanEquals(other RTValue, position SEPos) (RTValue, error) {
	switch other.GetType() {
	case RTT_INT:
		return NewRTBool(position, rTInt.Value >= other.GetValue().(int), rTInt.Environment), nil
	case RTT_FLOAT:
		return NewRTBool(position, float64(rTInt.Value) >= other.GetValue().(float64), rTInt.Environment), nil
	}

	return nil, NewValueRTError(
		GREATER_THAN_EQUALS,
		rTInt,
		other,
		position,
		rTInt.Environment,
	)
}

func (rTInt *RTInt) LessThan(other RTValue, position SEPos) (RTValue, error) {
	switch other.GetType() {
	case RTT_INT:
		return NewRTBool(position, rTInt.Value < other.GetValue().(int), rTInt.Environment), nil
	case RTT_FLOAT:
		return NewRTBool(position, float64(rTInt.Value) < other.GetValue().(float64), rTInt.Environment), nil
	}

	return nil, NewValueRTError(
		LESS_THAN,
		rTInt,
		other,
		position,
		rTInt.Environment,
	)
}

func (rTInt *RTInt) LessThanEquals(other RTValue, position SEPos) (RTValue, error) {
	switch other.GetType() {
	case RTT_INT:
		return NewRTBool(position, rTInt.Value <= other.GetValue().(int), rTInt.Environment), nil
	case RTT_FLOAT:
		return NewRTBool(position, float64(rTInt.Value) <= other.GetValue().(float64), rTInt.Environment), nil
	}

	return nil, NewValueRTError(
		LESS_THAN_EQUALS,
		rTInt,
		other,
		position,
		rTInt.Environment,
	)
}

func (rTInt *RTInt) Not(position SEPos) (RTValue, error) {
	if rTInt.Value == 0 {
		return NewRTBool(position, true, rTInt.Environment), nil
	}

	return NewRTBool(position, false, rTInt.Environment), nil
}

func (rTInt *RTInt) ToBool(position SEPos) (RTValue, error) {
	if rTInt.Value == 0 {
		return NewRTBool(position, false, rTInt.Environment), nil
	}

	return NewRTBool(position, true, rTInt.Environment), nil
}

func (rTInt *RTInt) Call(arguments []RTValue, position SEPos, interpreter *Interpreter) (RTValue, error) {
	return nil, NewInvalidCallRTError(rTInt, position, rTInt.Environment)
}
