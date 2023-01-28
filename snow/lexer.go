package snow

import (
	"fmt"
)

type Lexer struct {
	pos         SimplePos
	file        *File
	currentChar byte
	end         bool
}

func NewLexer(file *File) *Lexer {
	return &Lexer{
		pos:  *NewSimplePos(-1, 1, -1),
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

func (lexer *Lexer) createSimpleToken(tType TokenType) *Token {
	return NewToken(tType, "", *lexer.pos.AsSEPos(lexer.file))
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

func (lexer *Lexer) Tokenize() ([]Token, []error) {
	tokens := make([]Token, 0)
	errors := make([]error, 0)

	lexer.advance()

	for !lexer.end {
		startPos := lexer.pos

		switch lexer.currentChar {
		case '\n':
			tokens = append(tokens, *lexer.createSimpleToken(NEWLINE))
			lexer.advance()
		case ' ':
			lexer.advance()
		case '\t':
			lexer.pos.Col += 2
			lexer.advance()
		case '#':
			_, err := lexer.makeComment()
			if err != nil {
				errors = append(errors, err)
			}
		case '.':
			tokens = append(tokens, *lexer.createSimpleToken(DOT))
			lexer.advance()
		case ',':
			tokens = append(tokens, *lexer.createSimpleToken(COMMA))
			lexer.advance()
		case '+':
			tokens = append(tokens, *lexer.createSimpleToken(PLUS))
			lexer.advance()
		case '-':
			tokens = append(tokens, *lexer.createSimpleToken(DASH))
			lexer.advance()
		case '*':
			tokens = append(tokens, *lexer.createSimpleToken(STAR))
			lexer.advance()
		case '/':
			tokens = append(tokens, *lexer.createSimpleToken(SLASH))
			lexer.advance()
		case '(':
			tokens = append(tokens, *lexer.createSimpleToken(LPAREN))
			lexer.advance()
		case ')':
			tokens = append(tokens, *lexer.createSimpleToken(RPAREN))
			lexer.advance()
		case '{':
			tokens = append(tokens, *lexer.createSimpleToken(LCURLYBRACKET))
			lexer.advance()
		case '}':
			tokens = append(tokens, *lexer.createSimpleToken(RCURLYBRACKET))
			lexer.advance()
		case '=':
			if !lexer.end && lexer.peek() == '=' {
				lexer.advance()

				tokens = append(tokens, *NewToken(
					EQUALS,
					"",
					*startPos.CreateSEPos(lexer.pos, lexer.file),
				))

				lexer.advance()
			} else {
				tokens = append(tokens, *lexer.createSimpleToken(SINGLE_EQUALS))
				lexer.advance()
			}
		case '<':
			if !lexer.end && lexer.peek() == '=' {
				lexer.advance()

				tokens = append(tokens, *NewToken(
					LESS_THAN_EQUALS,
					"",
					*startPos.CreateSEPos(lexer.pos, lexer.file),
				))

				lexer.advance()
			} else {
				tokens = append(tokens, *lexer.createSimpleToken(LESS_THAN))
				lexer.advance()
			}
		case '>':
			if !lexer.end && lexer.peek() == '=' {
				lexer.advance()

				tokens = append(tokens, *NewToken(
					GREATER_THAN_EQUALS,
					"",
					*startPos.CreateSEPos(lexer.pos, lexer.file),
				))

				lexer.advance()
			} else {
				tokens = append(tokens, *lexer.createSimpleToken(GREATER_THAN))
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

				tokens = append(tokens, *NewToken(
					NOT_EQUALS,
					"",
					*startPos.CreateSEPos(lexer.pos, lexer.file),
				))

				lexer.advance()
			} else {
				errors = append(errors, *NewSnowError(
					ILLEGAL_CHAR_ERROR,
					fmt.Sprintf("illegal character '%c'", lexer.currentChar),
					"",
					*lexer.pos.AsSEPos(lexer.file),
				))

				lexer.advance()
			}
		}
	}

	tokens = append(tokens, *lexer.createSimpleToken(EOF))

	return tokens, errors
}

func (lexer *Lexer) makeNumber() (Token, error) {
	startPos := lexer.pos
	endPos := startPos
	numberStr := string(lexer.currentChar)
	isFloat := false

	lexer.advance()

	for !lexer.end && (lexer.isDigit(lexer.currentChar) || (lexer.currentChar == '.' && lexer.isDigit(lexer.peek()))) {
		if lexer.currentChar == '.' {
			if isFloat && !lexer.isDigit(lexer.peek()) {
				return *NewToken(
					FLOAT,
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
			return Token{}, NewSnowError(
				TRAILING_DOT_ERROR,
				"Trailing dots are not allowed",
				fmt.Sprintf("To define a float add a zero after: '%s.0'", numberStr),
				*pos.AsSEPos(lexer.file),
			)
		} else {
			return Token{}, NewSnowError(
				MULTIPLE_DOTS_ERROR,
				"More than one dot while defining a float is not allowed",
				fmt.Sprintf("Remove the dot: '%s'", numberStr),
				*pos.AsSEPos(lexer.file),
			)
		}
	} else if isFloat && lexer.currentChar == '.' && lexer.isDigit(lexer.peek()) {
		pos := lexer.pos

		lexer.advance()

		return Token{}, NewSnowError(
			MULTIPLE_DOTS_ERROR,
			"More than one dot while defining a float is not allowed",
			fmt.Sprintf("Remove the dot: '%s'", numberStr),
			*pos.AsSEPos(lexer.file),
		)
	}

	if isFloat {
		return *NewToken(
			FLOAT,
			numberStr,
			*startPos.CreateSEPos(endPos, lexer.file),
		), nil
	}

	tok := *NewToken(
		INT,
		numberStr,
		*startPos.CreateSEPos(endPos, lexer.file),
	)

	return tok, nil
}

func (lexer *Lexer) makeString() (Token, error) {
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
		return Token{}, NewSnowError(
			UNTERMINATED_STRING_ERROR,
			"the string was never closed",
			fmt.Sprintf("Add a closing quote to the end of the string: %s", str),
			*startPos.CreateSEPos(endPos, lexer.file),
		)
	}

	endPos = lexer.pos
	lexer.advance()

	return *NewToken(
		STRING,
		strValue,
		*startPos.CreateSEPos(endPos, lexer.file),
	), nil
}

func (lexer *Lexer) makeIdentifierKeyword() (Token, error) {
	startPos := lexer.pos
	endPos := startPos
	valueStr := string(lexer.currentChar)

	lexer.advance()

	for lexer.isAlphaDigit(lexer.currentChar) && !lexer.end {
		valueStr += string(lexer.currentChar)

		endPos = lexer.pos
		lexer.advance()
	}

	tType, ok := Keywords[valueStr]
	if !ok {
		tType = IDENTIFIER
	}

	return *NewToken(
		tType,
		valueStr,
		*startPos.CreateSEPos(endPos, lexer.file),
	), nil
}

func (lexer *Lexer) makeComment() (Token, error) {
	inline := false
	startPos := lexer.pos
	var lastPos SimplePos

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

			return Token{}, nil
		} else if !inline && lexer.currentChar == '\n' {
			return Token{}, nil
		}

		lastPos = lexer.pos

		lexer.advance()
	}

	if inline && lexer.end {
		return Token{}, NewSnowError(
			UNTERMINATED_INLINE_COMMENT_ERROR,
			"the inline comment was never closed",
			"Add '/#' to close the inline comment",
			*startPos.CreateSEPos(lastPos, lexer.file),
		)
	}

	return Token{}, nil
}
