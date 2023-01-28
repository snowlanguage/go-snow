package snow

import (
	"fmt"
	"strconv"
)

type Parser struct {
	tokens       []Token
	file         *File
	currentToken Token
	index        int
	inBlock      int
	inLoop       int
}

func NewParser(tokens []Token, file *File) *Parser {
	return &Parser{
		tokens: tokens,
		file:   file,
		index:  -1,
	}
}

func (parser *Parser) advance() {
	if parser.currentToken.TType != EOF {
		parser.index++
		parser.currentToken = parser.tokens[parser.index]
	}
}

func (parser *Parser) peek() Token {
	if parser.currentToken.TType != EOF {
		return parser.tokens[parser.index+1]
	}

	return parser.currentToken
}

func (parser *Parser) consume(tType TokenType) error {
	if parser.currentToken.TType != tType {
		pos := parser.currentToken.Pos
		tType2 := parser.currentToken.TType

		parser.advance()

		return NewSnowError(
			EXPECTED_TOKEN_ERROR,
			fmt.Sprintf("expected token of type '%s', not token of type '%s'", tType, tType2),
			"",
			pos,
		)
	}

	parser.advance()

	return nil
}

func (parser *Parser) binary(t1 TokenType, t2 TokenType, t3 TokenType, t4 TokenType, function func() (Expr, error)) (Expr, error) {
	startPos := parser.currentToken.Pos.Start

	left, err := function()
	if err != nil {
		return nil, err
	}

	for parser.currentToken.TType == t1 || parser.currentToken.TType == t2 || parser.currentToken.TType == t3 || parser.currentToken.TType == t4 {
		opToken := parser.currentToken
		parser.advance()

		right, err := function()
		if err != nil {
			return nil, err
		}

		left = NewBinaryExpr(
			left,
			right,
			opToken,
			*startPos.CreateSEPos(right.GetPosition().End, parser.currentToken.Pos.File),
		)
	}

	return left, nil
}

func (parser *Parser) Parse() ([]Stmt, error) {
	statements := make([]Stmt, 0)

	parser.advance()

	for parser.currentToken.TType != EOF {
		if parser.currentToken.TType == NEWLINE {
			parser.advance()

			continue
		}

		stmt, err := parser.deceleration()
		if err != nil {
			return nil, err
		}

		statements = append(statements, stmt)
	}

	return statements, nil
}

func (parser *Parser) deceleration() (Stmt, error) {
	if parser.currentToken.TType == NEWLINE {
		return nil, nil
	}

	if parser.currentToken.TType == VAR || parser.currentToken.TType == CONST {
		varDeclStmt, err := parser.varDeclStmt()
		if err != nil {
			return nil, err
		}

		return varDeclStmt, nil
	} else if parser.currentToken.TType == FUNCTION {
		fmt.Println("yes")
		return parser.functionDeclStmt()
	}

	statement, err := parser.statement()
	if err != nil {
		return nil, err
	}

	return statement, nil
}

func (parser *Parser) functionDeclStmt() (Stmt, error) {
	startPos := parser.currentToken.Pos.Start

	err := parser.consume(FUNCTION)
	if err != nil {
		return nil, err
	}

	fmt.Println("yes2")

	name := parser.currentToken
	err = parser.consume(IDENTIFIER)
	if err != nil {
		return nil, err
	}

	fmt.Println("yes3")

	err = parser.consume(LPAREN)
	if err != nil {
		return nil, err
	}

	parameters := make([]Token, 0)

	for parser.currentToken.TType != EOF && parser.currentToken.TType != RPAREN {
		parameter := parser.currentToken

		err = parser.consume(IDENTIFIER)
		if err != nil {
			return nil, err
		}

		if parser.currentToken.TType != RPAREN {
			err = parser.consume(COMMA)
			if err != nil {
				return nil, err
			}
		}

		parameters = append(parameters, parameter)
	}

	err = parser.consume(RPAREN)
	if err != nil {
		return nil, err
	}

	fmt.Println("yes4")

	block, err := parser.blockStatement()
	if err != nil {
		return nil, err
	}

	fmt.Println("yes5")

	return NewFunctionDeclStmt(name.Value, parameters, block, *startPos.CreateSEPos(block.GetPos().End, block.GetPos().File)), nil
}

func (parser *Parser) varDeclStmt() (Stmt, error) {
	startTok := parser.currentToken

	parser.advance()

	if parser.currentToken.TType != IDENTIFIER {
		return nil, NewUnexpectedTokenError(IDENTIFIER, parser.currentToken)
	}

	identifier := parser.currentToken

	parser.advance()

	err := parser.consume(SINGLE_EQUALS)
	if err != nil {
		return nil, err
	}

	expr, err := parser.expression()
	if err != nil {
		return nil, err
	}

	if parser.currentToken.TType != EOF {
		err = parser.consume(NEWLINE)
		if err != nil {
			return nil, err
		}
	}

	return NewVarDeclStmt(startTok, identifier, expr, *startTok.Pos.Start.CreateSEPos(expr.GetPosition().End, startTok.Pos.File)), nil
}

func (parser *Parser) statement() (Stmt, error) {
	if parser.currentToken.TType == LCURLYBRACKET {
		return parser.blockStatement()
	} else if parser.currentToken.TType == WHILE {
		return parser.whileStatement()
	} else if parser.currentToken.TType == BREAK {
		return parser.breakStmt()
	} else if parser.currentToken.TType == CONTINUE {
		return parser.continueStmt()
	}

	statement, err := parser.expressionStmt()
	if err != nil {
		return nil, err
	}

	return statement, nil
}

func (parser *Parser) whileStatement() (Stmt, error) {
	startPos := parser.currentToken.Pos.Start

	err := parser.consume(WHILE)
	if err != nil {
		return nil, err
	}

	expr, err := parser.expression()
	if err != nil {
		return nil, err
	}

	parser.inLoop += 1

	stmt, err := parser.statement()
	if err != nil {
		return nil, err
	}

	parser.inLoop -= 1

	return NewWhileStmt(stmt, expr, *startPos.CreateSEPos(stmt.GetPos().End, stmt.GetPos().File)), nil
}

func (parser *Parser) blockStatement(params ...string) (Stmt, error) {
	startPos := parser.currentToken.Pos.Start
	file := parser.currentToken.Pos.File
	statements := make([]Stmt, 0)

	err := parser.consume(LCURLYBRACKET)
	if err != nil {
		return nil, err
	}

	for parser.currentToken.TType == NEWLINE {
		parser.advance()
	}

	parser.inBlock += 1

	for parser.currentToken.TType != RCURLYBRACKET && parser.currentToken.TType != EOF {
		statement, err := parser.deceleration()
		if err != nil {
			return nil, err
		}

		statements = append(statements, statement)
	}

	parser.inBlock -= 1

	if parser.currentToken.TType != RCURLYBRACKET {
		return nil, NewUnexpectedTokenError(RCURLYBRACKET, parser.currentToken)
	}

	parser.advance()

	endPos := parser.currentToken.Pos

	parser.advance()

	name := ""
	if len(params) != 0 {
		name = params[0]
	}

	return NewBlockStmt(statements, name, *startPos.CreateSEPos(endPos.End, file)), nil
}

func (parser *Parser) expressionStmt() (Stmt, error) {
	pos := parser.currentToken.Pos

	expression, err := parser.expression()
	if err != nil {
		return nil, err
	}

	if parser.currentToken.TType != EOF && !(parser.inBlock != 0 && parser.currentToken.TType == RCURLYBRACKET) {
		err = parser.consume(NEWLINE)
		if err != nil {
			return nil, err
		}
	}

	return NewExpressionStmt(expression, pos), nil
}

func (parser *Parser) breakStmt() (Stmt, error) {
	pos := parser.currentToken.Pos

	err := parser.consume(BREAK)
	if err != nil {
		return nil, err
	}

	if parser.inLoop == 0 {
		return nil, NewSnowError(
			BREAK_OUTSIDE_OF_LOOP_ERROR,
			"break statement found outside of loop",
			"Break statements can only be used inside of loops",
			pos,
		)
	}

	if !(parser.inBlock != 0 && parser.currentToken.TType == RCURLYBRACKET) {
		err = parser.consume(NEWLINE)
		if err != nil {
			return nil, err
		}
	}

	return NewBreakStmt(pos), nil
}

func (parser *Parser) continueStmt() (Stmt, error) {
	pos := parser.currentToken.Pos

	err := parser.consume(CONTINUE)
	if err != nil {
		return nil, err
	}

	if parser.inLoop == 0 {
		return nil, NewSnowError(
			CONTINUE_OUTSIDE_OF_LOOP_ERROR,
			"continue statement found outside of loop",
			"Continue statements can only be used inside of loops",
			pos,
		)
	}

	if !(parser.inBlock != 0 && parser.currentToken.TType == RCURLYBRACKET) {
		err = parser.consume(NEWLINE)
		if err != nil {
			return nil, err
		}
	}

	return NewContinueStmt(pos), nil
}

func (parser *Parser) expression() (Expr, error) {
	assignment, err := parser.assignment()
	if err != nil {
		return nil, err
	}

	return assignment, nil
}

func (parser *Parser) assignment() (Expr, error) {
	logicOr, err := parser.logicOr()
	if err != nil {
		return nil, err
	}

	if parser.currentToken.TType == SINGLE_EQUALS {
		switch logicOr := logicOr.(type) {
		case *DotExpr:
			parser.advance()

			val, err := parser.expression()
			if err != nil {
				return nil, err
			}

			return NewVarAssignmentExpr(
				logicOr.Left,
				logicOr.Right.Value,
				val,
				*logicOr.GetPosition().Start.CreateSEPos(val.GetPosition().End, logicOr.GetPosition().File),
			), nil
		case *VarAccessExpr:
			parser.advance()

			val, err := parser.expression()
			if err != nil {
				return nil, err
			}

			return NewVarAssignmentExpr(
				nil,
				logicOr.Value,
				val,
				*logicOr.GetPosition().Start.CreateSEPos(val.GetPosition().End, logicOr.GetPosition().File),
			), nil
		default:
			return logicOr, nil
		}
	}

	return logicOr, nil
}

func (parser *Parser) logicOr() (Expr, error) {
	logicAnd, err := parser.logicAnd()
	if err != nil {
		return nil, err
	}

	return logicAnd, nil
}

func (parser *Parser) logicAnd() (Expr, error) {
	equality, err := parser.equality()
	if err != nil {
		return nil, err
	}

	return equality, nil
}

func (parser *Parser) equality() (Expr, error) {
	binary, err := parser.binary(
		EQUALS,
		NOT_EQUALS,
		PLACEHOLDER,
		PLACEHOLDER,
		parser.comparison,
	)
	if err != nil {
		return nil, err
	}

	return binary, nil
}

func (parser *Parser) comparison() (Expr, error) {
	binary, err := parser.binary(
		GREATER_THAN,
		GREATER_THAN_EQUALS,
		LESS_THAN,
		LESS_THAN_EQUALS,
		parser.term,
	)
	if err != nil {
		return nil, err
	}

	return binary, nil
}

func (parser *Parser) term() (Expr, error) {
	binary, err := parser.binary(
		PLUS,
		DASH,
		PLACEHOLDER,
		PLACEHOLDER,
		parser.factor,
	)
	if err != nil {
		return nil, err
	}

	return binary, nil
}

func (parser *Parser) factor() (Expr, error) {
	binary, err := parser.binary(
		STAR,
		SLASH,
		PLACEHOLDER,
		PLACEHOLDER,
		parser.unary,
	)
	if err != nil {
		return nil, err
	}

	return binary, nil
}

func (parser *Parser) unary() (Expr, error) {
	if parser.currentToken.TType == DASH || parser.currentToken.TType == NOT {
		token := parser.currentToken

		parser.advance()

		right, err := parser.unary()
		if err != nil {
			return nil, err
		}

		return NewUnaryExpr(
			token,
			right,
			*token.Pos.Start.CreateSEPos(right.GetPosition().End, parser.currentToken.Pos.File),
		), nil
	}

	call, err := parser.call()
	if err != nil {
		return nil, err
	}

	return call, nil
}

func (parser *Parser) call() (Expr, error) {
	primary, err := parser.primary()
	if err != nil {
		return nil, err
	}

	for parser.currentToken.TType == DOT || parser.currentToken.TType == LPAREN {
		if parser.currentToken.TType == DOT {
			parser.advance()

			if parser.currentToken.TType != IDENTIFIER {
				return nil, NewUnexpectedTokenError(IDENTIFIER, parser.currentToken)
			}

			right := parser.currentToken

			parser.advance()

			primary = NewDotExpr(primary, right, *primary.GetPosition().Start.CreateSEPos(right.Pos.End, right.Pos.File))
		} else {
			parser.advance()

			arguments := make([]Expr, 0)

			fmt.Println(parser.currentToken.TType)
			for parser.currentToken.TType != RPAREN && parser.currentToken.TType != EOF {
				argument, err := parser.expression()
				if err != nil {
					return nil, err
				}

				arguments = append(arguments, argument)
				fmt.Println("a", arguments)

				if parser.currentToken.TType != RPAREN {
					err = parser.consume(COMMA)
					if err != nil {
						return nil, err
					}
				}
			}

			fmt.Printf("arguments: %v\n", arguments)

			endPos := parser.currentToken.Pos.End

			err := parser.consume(RPAREN)
			if err != nil {
				return nil, err
			}

			primary = NewCallExpr(primary, arguments, *primary.GetPosition().Start.CreateSEPos(endPos, primary.GetPosition().File))
		}
	}

	return primary, err
}

func (parser *Parser) primary() (Expr, error) {
	startToken := parser.currentToken

	switch parser.currentToken.TType {
	case INT:
		parser.advance()

		intValue, err := strconv.Atoi(startToken.Value)
		if err != nil {
			return nil, NewSnowError(
				TOO_BIG_VALUE_ERROR,
				fmt.Sprintf("the value of number of type %s is too big", startToken.TType),
				"",
				startToken.Pos,
			)
		}

		return NewIntLiteralExpr(intValue, startToken.Pos), nil
	case FLOAT:
		parser.advance()

		floatValue, err := strconv.ParseFloat(startToken.Value, 64)
		if err != nil {
			return nil, NewSnowError(
				TOO_BIG_VALUE_ERROR,
				fmt.Sprintf("the value of number of type %s is too big", startToken.TType),
				"",
				startToken.Pos,
			)
		}

		return NewFloatLiteralExpr(floatValue, startToken.Pos), nil
	case TRUE:
		parser.advance()

		return NewBoolLiteralExpr(true, startToken.Pos), nil
	case FALSE:
		parser.advance()

		return NewBoolLiteralExpr(false, startToken.Pos), nil
	case LPAREN:
		parser.advance()

		expr, err := parser.expression()
		if err != nil {
			return nil, err
		}

		endPos := parser.currentToken.Pos.End

		err = parser.consume(RPAREN)
		if err != nil {
			return nil, err
		}

		pos := startToken.Pos.Start.CreateSEPos(endPos, startToken.Pos.File)

		return NewGroupingExpr(expr, *pos), nil
	case IDENTIFIER:
		parser.advance()

		return NewVarAccessExpr(startToken.Value, startToken.Pos), nil
	default:
		err := NewSnowError(
			INVALID_TOKEN_TYPE_ERROR,
			fmt.Sprintf("token of type '%s' is invalid", startToken.TType),
			"",
			parser.currentToken.Pos,
		)
		return nil, err
	}
}
