package lexer

import (
	"fmt"

	"github.com/snowlanguage/go-snow/file"
	"github.com/snowlanguage/go-snow/position"
	"github.com/snowlanguage/go-snow/token"
)

type Lexer struct {
	pos         position.SimplePos
	file        file.File
	currentChar byte
	end         bool
}

func NewLexer(file file.File) *Lexer {
	return &Lexer{
		pos:  *position.NewSimplePos(-1, 1, -1),
		file: file,
	}
}

func (lexer *Lexer) advance() {
	lexer.pos.Idx += 1
	lexer.pos.Col += 1

	if lexer.currentChar == '\n' {
		lexer.pos.Ln += 1
		lexer.pos.Col = 0
	}

	if lexer.pos.Idx >= len(lexer.file.Code) {
		lexer.end = true
	} else {
		lexer.currentChar = lexer.file.Code[lexer.pos.Idx]
	}
}

func (lexer *Lexer) createSimpleToken(tType token.TokenType) *token.Token {
	return token.NewToken(token.PLUS, "", *lexer.pos.AsSEPos(&lexer.file))
}

func (lexer *Lexer) Tokenize() []token.Token {
	tokens := make([]token.Token, 0)

	lexer.advance()

	for !lexer.end {
		fmt.Println(lexer.end, string(lexer.currentChar))
		switch lexer.currentChar {
		case '+':
			tokens = append(tokens, *lexer.createSimpleToken(token.PLUS))
			lexer.advance()
		case '-':
			tokens = append(tokens, *lexer.createSimpleToken(token.DASH))
			lexer.advance()
		case '*':
			tokens = append(tokens, *lexer.createSimpleToken(token.STAR))
			lexer.advance()
		case '/':
			tokens = append(tokens, *lexer.createSimpleToken(token.SLASH))
			lexer.advance()
		default:
			lexer.advance()
		}

	}

	return tokens
}
