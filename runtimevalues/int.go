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
