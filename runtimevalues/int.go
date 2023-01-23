package runtimevalues

import (
	"fmt"

	"github.com/snowlanguage/go-snow/position"
)

type RTInt struct {
	Pos   position.SEPos
	Value int
}

func NewRTInt(pos position.SEPos, value int) *RTInt {
	return &RTInt{
		Pos:   pos,
		Value: value,
	}
}

func (rTInt *RTInt) ToString() string {
	return fmt.Sprintf("(INT: %d)", rTInt.Value)
}

func (rTInt *RTInt) GetType() RTType {
	return RTT_INT
}

func (rTInt *RTInt) GetValue() interface{} {
	return rTInt.Value
}

func (rTInt *RTInt) Add(other RTValue, position position.SEPos) (RTValue, error) {
	// TODO FIX
	return NewRTInt(position, rTInt.Value+other.GetValue().(int)), nil
}
