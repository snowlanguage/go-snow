package token

type TokenType string

const (
	PLUS  TokenType = "PLUS"
	DASH  TokenType = "DASH"
	STAR  TokenType = "STAR"
	SLASH TokenType = "SLASH"

	INT   TokenType = "INT"
	FLOAT TokenType = "FLOAT"

	LPAREN TokenType = "LPAREN"
	RPAREN TokenType = "RPAREN"
)
