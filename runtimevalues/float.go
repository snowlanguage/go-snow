package runtimevalues

import (
	"fmt"

	"github.com/snowlanguage/go-snow/position"
	"github.com/snowlanguage/go-snow/token"
)

type RTFloat struct {
	Pos         position.SEPos
	Value       float64
	Environment *Environment
}

func NewRTFloat(pos position.SEPos, value float64, env *Environment) *RTFloat {
	return &RTFloat{
		Pos:         pos,
		Value:       value,
		Environment: env,
	}
}

func (rTFloat *RTFloat) ToString() string {
	return fmt.Sprintf("(FLOAT: %f)", rTFloat.Value)
}

func (rTFloat *RTFloat) ValueToString() string {
	return fmt.Sprintf("%f", rTFloat.Value)
}

func (rTFloat *RTFloat) GetType() RTType {
	return RTT_FLOAT
}

func (rTFloat *RTFloat) GetValue() interface{} {
	return rTFloat.Value
}

func (rTFloat *RTFloat) GetEnvironment() *Environment {
	return rTFloat.Environment
}

func (rTFloat *RTFloat) Add(other RTValue, position position.SEPos) (RTValue, error) {
	fmt.Println(rTFloat.Value + other.GetValue().(float64))
	fmt.Println("here!!!")
	switch other.GetType() {
	case RTT_FLOAT:
		return NewRTFloat(position, rTFloat.Value+other.GetValue().(float64), rTFloat.Environment), nil
	case RTT_INT:
		return NewRTFloat(position, rTFloat.Value+float64(other.GetValue().(int)), rTFloat.Environment), nil
	}

	return nil, NewValueRTError(
		token.PLUS,
		rTFloat,
		other,
		position,
		rTFloat.Environment,
	)
}

func (rTFloat *RTFloat) Subtract(other RTValue, position position.SEPos) (RTValue, error) {
	switch other.GetType() {
	case RTT_FLOAT:
		return NewRTFloat(position, rTFloat.Value-other.GetValue().(float64), rTFloat.Environment), nil
	case RTT_INT:
		return NewRTFloat(position, rTFloat.Value-float64(other.GetValue().(int)), rTFloat.Environment), nil
	}

	return nil, NewValueRTError(
		token.PLUS,
		rTFloat,
		other,
		position,
		rTFloat.Environment,
	)
}

func (rTFloat *RTFloat) Multiply(other RTValue, position position.SEPos) (RTValue, error) {
	switch other.GetType() {
	case RTT_FLOAT:
		return NewRTFloat(position, rTFloat.Value*other.GetValue().(float64), rTFloat.Environment), nil
	case RTT_INT:
		return NewRTFloat(position, rTFloat.Value*float64(other.GetValue().(int)), rTFloat.Environment), nil
	}

	return nil, NewValueRTError(
		token.PLUS,
		rTFloat,
		other,
		position,
		rTFloat.Environment,
	)
}

func (rTFloat *RTFloat) Divide(other RTValue, position position.SEPos) (RTValue, error) {
	switch other.GetType() {
	case RTT_FLOAT:
		if other.GetValue().(float64) == 0 {
			return nil, NewDivisionByZeroRTError(rTFloat, other, position, rTFloat.Environment)
		}

		return NewRTFloat(position, rTFloat.Value/other.GetValue().(float64), rTFloat.Environment), nil
	case RTT_INT:
		if other.GetValue().(int) == 0 {
			return nil, NewDivisionByZeroRTError(rTFloat, other, position, rTFloat.Environment)
		}

		return NewRTFloat(position, rTFloat.Value/float64(other.GetValue().(int)), rTFloat.Environment), nil
	}

	return nil, NewValueRTError(
		token.PLUS,
		rTFloat,
		other,
		position,
		rTFloat.Environment,
	)
}
