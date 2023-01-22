package token

type TokenType string

const (
	PLUS  TokenType = "PLUS"
	DASH  TokenType = "DASH"
	STAR  TokenType = "STAR"
	SLASH TokenType = "SLASH"

	LPAREN TokenType = "LPAREN"
	RPAREN TokenType = "RPAREN"

	INT        TokenType = "INT"
	FLOAT      TokenType = "FLOAT"
	STRING     TokenType = "STRING"
	IDENTIFIER TokenType = "IDENTIFIER"

	NEWLINE TokenType = "NEWLINE"

	EOF TokenType = "EOF"
)
