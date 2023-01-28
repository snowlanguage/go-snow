package token

type TokenType string

const (
	PLUS  TokenType = "PLUS"
	DASH  TokenType = "DASH"
	STAR  TokenType = "STAR"
	SLASH TokenType = "SLASH"

	LPAREN        TokenType = "LPAREN"
	RPAREN        TokenType = "RPAREN"
	LCURLYBRACKET TokenType = "LCURLYBRACKET"
	RCURLYBRACKET TokenType = "RCURLYBRACKET"

	SINGLE_EQUALS       TokenType = "SINGLE_EQUALS"
	EQUALS              TokenType = "EQUALS"
	NOT_EQUALS          TokenType = "NOT_EQUALS"
	GREATER_THAN        TokenType = "GREATER_THAN"
	GREATER_THAN_EQUALS TokenType = "GREATER_THAN_EQUALS"
	LESS_THAN           TokenType = "LESS_THAN"
	LESS_THAN_EQUALS    TokenType = "LESS_THAN_EQUALS"

	INT        TokenType = "INT"
	FLOAT      TokenType = "FLOAT"
	STRING     TokenType = "STRING"
	IDENTIFIER TokenType = "IDENTIFIER"

	NOT      TokenType = "NOT"
	TRUE     TokenType = "TRUE"
	FALSE    TokenType = "FALSE"
	VAR      TokenType = "VAR"
	CONST    TokenType = "CONST"
	WHILE    TokenType = "WHILE"
	CONTINUE TokenType = "CONTINUE"
	BREAK    TokenType = "BREAK"

	DOT TokenType = "DOT"

	NEWLINE TokenType = "NEWLINE"

	EOF TokenType = "EOF"

	PLACEHOLDER TokenType = "PLACEHOLDER"
)
