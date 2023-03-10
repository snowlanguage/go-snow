package snow

import (
	"fmt"
)

type Token struct {
	TType TokenType
	Value string
	Pos   SEPos
}

func NewToken(tToken TokenType, value string, pos SEPos) *Token {
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
	case LCURLYBRACKET:
		str = "{"
	case RCURLYBRACKET:
		str = "}"
	case EOF:
		str = "(EOF)"
	case SINGLE_EQUALS:
		str = "="
	case EQUALS:
		str = "=="
	case GREATER_THAN:
		str = ">"
	case GREATER_THAN_EQUALS:
		str = ">="
	case LESS_THAN:
		str = "<"
	case LESS_THAN_EQUALS:
		str = "<="
	case NEWLINE:
		str = "(NEWLINE)"
	case DOT:
		str = "."
	case COMMA:
		str = ","
	case INT:
		str = fmt.Sprintf("(INT: %s)", token.Value)
	case FLOAT:
		str = fmt.Sprintf("(FLOAT: %s)", token.Value)
	case STRING:
		str = fmt.Sprintf("(STRING: \"%s\")", token.Value)
	case IDENTIFIER:
		str = fmt.Sprintf("(IDENTIFIER: %s)", token.Value)
	default:
		str = fmt.Sprintf("(KEYWORD: %s)", token.Value)
	}

	return str
}
