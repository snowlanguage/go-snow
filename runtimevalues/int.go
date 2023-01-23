package runtimevalues

import (
	"fmt"

	"github.com/snowlanguage/go-snow/position"
	"github.com/snowlanguage/go-snow/token"
)

type RTInt struct {
	Pos         position.SEPos
	Value       int
	Environment *Environment
}

func NewRTInt(pos position.SEPos, value int, env *Environment) *RTInt {
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

func (rTInt *RTInt) Add(other RTValue, position position.SEPos) (RTValue, error) {
	switch other.GetType() {
	case RTT_INT:
		return NewRTInt(position, rTInt.Value+other.GetValue().(int), rTInt.Environment), nil
	case RTT_FLOAT:
		return NewRTFloat(position, float64(rTInt.Value)+other.GetValue().(float64), rTInt.Environment), nil
	}

	return nil, NewValueRTError(
		token.PLUS,
		rTInt,
		other,
		position,
		rTInt.Environment,
	)
}
