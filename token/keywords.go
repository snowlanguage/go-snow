package token

var Keywords = map[string]TokenType{
	"not":   NOT,
	"true":  TRUE,
	"false": FALSE,
	"var":   VAR,
	"const": CONST,
}
