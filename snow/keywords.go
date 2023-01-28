package snow

var Keywords = map[string]TokenType{
	"not":      NOT,
	"true":     TRUE,
	"false":    FALSE,
	"var":      VAR,
	"const":    CONST,
	"while":    WHILE,
	"continue": CONTINUE,
	"break":    BREAK,
	"function": FUNCTION,
}
