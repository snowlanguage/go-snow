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
	file        *file.File
	currentChar byte
	end         bool
}

func NewLexer(file *file.File) *Lexer {
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
		lexer.currentChar = 0x0
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
	return token.NewToken(tType, "", *lexer.pos.AsSEPos(lexer.file))
}

func (lexer *Lexer) isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func (lexer *Lexer) isAlpha(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || c == '_'
}

func (lexer *Lexer) isAlphaDigit(c byte) bool {
	return lexer.isAlpha(c) || lexer.isDigit(c)
}

func (lexer *Lexer) Tokenize() ([]token.Token, []error) {
	tokens := make([]token.Token, 0)
	errors := make([]error, 0)

	lexer.advance()

	for !lexer.end {
		startPos := lexer.pos

		switch lexer.currentChar {
		case '\n':
			tokens = append(tokens, *lexer.createSimpleToken(token.NEWLINE))
			lexer.advance()
		case ' ':
			lexer.advance()
		case '\t':
			lexer.advance()
		case '#':
			_, err := lexer.makeComment()
			if err != nil {
				errors = append(errors, err)
			}
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
		case '=':
			if !lexer.end && lexer.peek() == '=' {
				lexer.advance()

				tokens = append(tokens, *token.NewToken(
					token.EQUALS,
					"",
					*startPos.CreateSEPos(lexer.pos, lexer.file),
				))

				lexer.advance()
			} else {
				tokens = append(tokens, *lexer.createSimpleToken(token.SINGLE_EQUALS))
				lexer.advance()
			}
		case '<':
			if !lexer.end && lexer.peek() == '=' {
				lexer.advance()

				tokens = append(tokens, *token.NewToken(
					token.LESS_THAN_EQUALS,
					"",
					*startPos.CreateSEPos(lexer.pos, lexer.file),
				))

				lexer.advance()
			} else {
				tokens = append(tokens, *lexer.createSimpleToken(token.LESS_THAN))
				lexer.advance()
			}
		case '>':
			if !lexer.end && lexer.peek() == '=' {
				lexer.advance()

				tokens = append(tokens, *token.NewToken(
					token.GREATER_THAN_EQUALS,
					"",
					*startPos.CreateSEPos(lexer.pos, lexer.file),
				))

				lexer.advance()
			} else {
				tokens = append(tokens, *lexer.createSimpleToken(token.GREATER_THAN))
				lexer.advance()
			}
		default:
			if lexer.isDigit(lexer.currentChar) {
				tok, err := lexer.makeNumber()
				if err == nil {
					tokens = append(tokens, tok)
				} else {
					errors = append(errors, err)
				}
			} else if lexer.currentChar == '"' || lexer.currentChar == '\'' {
				tok, err := lexer.makeString()
				if err == nil {
					tokens = append(tokens, tok)
				} else {
					errors = append(errors, err)
				}
			} else if lexer.isAlpha(lexer.currentChar) {
				tok, err := lexer.makeIdentifierKeyword()
				if err == nil {
					tokens = append(tokens, tok)
				} else {
					errors = append(errors, err)
				}
			} else if lexer.currentChar == '!' && lexer.peek() == '=' {
				lexer.advance()

				tokens = append(tokens, *token.NewToken(
					token.NOT_EQUALS,
					"",
					*startPos.CreateSEPos(lexer.pos, lexer.file),
				))

				lexer.advance()
			} else {
				errors = append(errors, *snowerror.NewSnowError(
					snowerror.ILLEGAL_CHAR_ERROR,
					fmt.Sprintf("illegal character '%c'", lexer.currentChar),
					"",
					*lexer.pos.AsSEPos(lexer.file),
				))

				lexer.advance()
			}
		}
	}

	tokens = append(tokens, *lexer.createSimpleToken(token.EOF))

	return tokens, errors
}

func (lexer *Lexer) makeNumber() (token.Token, error) {
	startPos := lexer.pos
	endPos := startPos
	numberStr := string(lexer.currentChar)
	isFloat := false

	lexer.advance()

	for !lexer.end && (lexer.isDigit(lexer.currentChar) || (lexer.currentChar == '.' && lexer.isDigit(lexer.peek()))) {
		if lexer.currentChar == '.' {
			if isFloat && !lexer.isDigit(lexer.peek()) {
				return *token.NewToken(
					token.FLOAT,
					numberStr,
					*startPos.CreateSEPos(endPos, lexer.file),
				), nil
			} else if isFloat {
				break
			}

			isFloat = true
		}

		numberStr += string(lexer.currentChar)

		endPos = lexer.pos

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
				*pos.AsSEPos(lexer.file),
			)
		} else {
			return token.Token{}, snowerror.NewSnowError(
				snowerror.MULTIPLE_DOTS_ERROR,
				"More than one dot while defining a float is not allowed",
				fmt.Sprintf("Remove the dot: '%s'", numberStr),
				*pos.AsSEPos(lexer.file),
			)
		}
	} else if isFloat && lexer.currentChar == '.' && lexer.isDigit(lexer.peek()) {
		pos := lexer.pos

		lexer.advance()

		return token.Token{}, snowerror.NewSnowError(
			snowerror.MULTIPLE_DOTS_ERROR,
			"More than one dot while defining a float is not allowed",
			fmt.Sprintf("Remove the dot: '%s'", numberStr),
			*pos.AsSEPos(lexer.file),
		)
	}

	if isFloat {
		return *token.NewToken(
			token.FLOAT,
			numberStr,
			*startPos.CreateSEPos(endPos, lexer.file),
		), nil
	}

	tok := *token.NewToken(
		token.INT,
		numberStr,
		*startPos.CreateSEPos(endPos, lexer.file),
	)

	return tok, nil
}

func (lexer *Lexer) makeString() (token.Token, error) {
	startPos := lexer.pos
	startChar := lexer.currentChar
	strValue := ""
	endPos := startPos

	lexer.advance()

	for lexer.currentChar != startChar && !lexer.end && lexer.currentChar != '\n' {
		strValue += string(lexer.currentChar)

		endPos = lexer.pos

		lexer.advance()
	}

	if lexer.end || lexer.currentChar == '\n' {
		str := ""
		if startChar == '"' {
			str = fmt.Sprintf("'\"%s\"'", strValue)
		} else {
			str = fmt.Sprintf("\"'%s'\"", strValue)
		}
		return token.Token{}, snowerror.NewSnowError(
			snowerror.UNTERMINATED_STRING_ERROR,
			"the string was never closed",
			fmt.Sprintf("Add a closing quote to the end of the string: %s", str),
			*startPos.CreateSEPos(endPos, lexer.file),
		)
	}

	endPos = lexer.pos
	lexer.advance()

	return *token.NewToken(
		token.STRING,
		strValue,
		*startPos.CreateSEPos(endPos, lexer.file),
	), nil
}

func (lexer *Lexer) makeIdentifierKeyword() (token.Token, error) {
	startPos := lexer.pos
	endPos := startPos
	valueStr := string(lexer.currentChar)

	lexer.advance()

	for lexer.isAlphaDigit(lexer.currentChar) && !lexer.end {
		valueStr += string(lexer.currentChar)

		endPos = lexer.pos
		lexer.advance()
	}

	tType, ok := token.Keywords[valueStr]
	if !ok {
		tType = token.IDENTIFIER
	}

	return *token.NewToken(
		tType,
		valueStr,
		*startPos.CreateSEPos(endPos, lexer.file),
	), nil
}

func (lexer *Lexer) makeComment() (token.Token, error) {
	inline := false
	startPos := lexer.pos
	var lastPos position.SimplePos

	lexer.advance()

	lastPos = lexer.pos

	if lexer.currentChar == '/' {
		inline = true
		lexer.advance()
	}

	for !lexer.end {
		if inline && lexer.currentChar == '/' && lexer.peek() == '#' {
			lexer.advance()
			lexer.advance()

			return token.Token{}, nil
		} else if !inline && lexer.currentChar == '\n' {
			return token.Token{}, nil
		}

		lastPos = lexer.pos

		lexer.advance()
	}

	if inline && lexer.end {
		return token.Token{}, snowerror.NewSnowError(
			snowerror.UNTERMINATED_INLINE_COMMENT_ERROR,
			"the inline comment was never closed",
			"Add '/#' to close the inline comment",
			*startPos.CreateSEPos(lastPos, lexer.file),
		)
	}

	return token.Token{}, nil
}
