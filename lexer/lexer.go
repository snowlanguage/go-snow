package lexer

import (
	"fmt"

	"github.com/snowlanguage/go-snow/file"
	"github.com/snowlanguage/go-snow/position"
	snowerror "github.com/snowlanguage/go-snow/snowError"
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

func (lexer *Lexer) peek() byte {
	if lexer.pos.Idx+1 >= len(lexer.file.Code) {
		return 0x0
	} else {
		return lexer.file.Code[lexer.pos.Idx+1]
	}
}

func (lexer *Lexer) isPeekEnd() bool {
	if lexer.pos.Idx+1 >= len(lexer.file.Code) {
		return true
	} else {
		return false
	}
}

func (lexer *Lexer) createSimpleToken(tType token.TokenType) *token.Token {
	return token.NewToken(tType, "", *lexer.pos.AsSEPos(&lexer.file))
}

func (lexer *Lexer) isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func (lexer *Lexer) isAlpha(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || c == '_'
}

func (lexer *Lexer) Tokenize() ([]token.Token, []error) {
	tokens := make([]token.Token, 0)
	errors := make([]error, 0)

	lexer.advance()

	for !lexer.end {
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
		case '(':
			tokens = append(tokens, *lexer.createSimpleToken(token.LPAREN))
			lexer.advance()
		case ')':
			tokens = append(tokens, *lexer.createSimpleToken(token.RPAREN))
			lexer.advance()
		default:
			if lexer.isDigit(lexer.currentChar) {
				tok, err := lexer.makeNumber()
				if err == nil {
					tokens = append(tokens, tok)
				} else {
					errors = append(errors, err)
				}
			} else {
				errors = append(errors, *snowerror.NewSnowError(
					snowerror.ILLEGAL_CHAR_ERROR,
					fmt.Sprintf("illegal character '%c'", lexer.currentChar),
					"",
					*lexer.pos.AsSEPos(&lexer.file),
				))

				lexer.advance()
			}
		}
	}

	return tokens, errors
}

func (lexer *Lexer) makeNumber() (token.Token, error) {
	startPos := lexer.pos
	numberStr := string(lexer.currentChar)
	isFloat := false

	lexer.advance()

	for !lexer.end && (lexer.isDigit(lexer.currentChar) || (lexer.currentChar == '.' && lexer.isDigit(lexer.peek()))) {
		if lexer.currentChar == '.' {
			if isFloat && !lexer.isDigit(lexer.peek()) {
				return *token.NewToken(
					token.FLOAT,
					numberStr,
					*startPos.CreateSEPos(lexer.pos, &lexer.file),
				), nil
			} else if isFloat {
				break
			}

			isFloat = true
		}

		numberStr += string(lexer.currentChar)

		lexer.advance()
	}

	if lexer.currentChar == '.' && (lexer.isPeekEnd() || lexer.peek() == '\n') {
		pos := lexer.pos

		lexer.advance()

		if !isFloat {
			return token.Token{}, snowerror.NewSnowError(
				snowerror.TRAILING_DOT_ERROR,
				"Trailing dots are not allowed",
				fmt.Sprintf("To define a float add a zero after: '%s.0'", numberStr),
				*pos.AsSEPos(&lexer.file),
			)
		} else {
			return token.Token{}, snowerror.NewSnowError(
				snowerror.MULTIPLE_DOTS_ERROR,
				"More than one dot while defining a float is not allowed",
				fmt.Sprintf("Remove the dot: '%s'", numberStr),
				*pos.AsSEPos(&lexer.file),
			)
		}
	} else if isFloat && lexer.currentChar == '.' && lexer.isDigit(lexer.peek()) {
		pos := lexer.pos

		lexer.advance()

		return token.Token{}, snowerror.NewSnowError(
			snowerror.MULTIPLE_DOTS_ERROR,
			"More than one dot while defining a float is not allowed",
			fmt.Sprintf("Remove the dot: '%s'", numberStr),
			*pos.AsSEPos(&lexer.file),
		)
	}

	if isFloat {
		return *token.NewToken(
			token.FLOAT,
			numberStr,
			*startPos.CreateSEPos(lexer.pos, &lexer.file),
		), nil
	}

	return *token.NewToken(
		token.INT,
		numberStr,
		*startPos.CreateSEPos(lexer.pos, &lexer.file),
	), nil
}
