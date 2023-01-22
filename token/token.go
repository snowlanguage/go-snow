package token

import (
	"fmt"

	"github.com/snowlanguage/go-snow/position"
)

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

func (token *Token) ToString() string {
	var str string

	switch token.TType {
	case PLUS:
		str = "+"
	case DASH:
		str = "-"
	case STAR:
		str = "*"
	case SLASH:
		str = "/"
	case LPAREN:
		str = "("
	case RPAREN:
		str = ")"
	case INT:
		str = fmt.Sprintf("(INT: %s)", token.Value)
	case FLOAT:
		str = fmt.Sprintf("(FLOAT: %s)", token.Value)
	}

	return str
}
