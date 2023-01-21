package token

import "github.com/snowlanguage/go-snow/position"

type Token struct {
	TType TokenType
	Value string
	Pos   position.SEPos
}

func NewToken(tToken TokenType, value string, pos position.SEPos) *Token {
	return &Token{
		TType: tToken,
		Value: value,
		Pos:   pos,
	}
}
