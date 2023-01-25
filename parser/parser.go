package parser

import (
	"fmt"
	"strconv"

	"github.com/snowlanguage/go-snow/file"
	parsevals "github.com/snowlanguage/go-snow/parseVals"
	snowerror "github.com/snowlanguage/go-snow/snowError"
	"github.com/snowlanguage/go-snow/token"
)

type Parser struct {
	tokens       []token.Token
	file         *file.File
	currentToken token.Token
	index        int
}

func NewParser(tokens []token.Token, file *file.File) *Parser {
	return &Parser{
		tokens: tokens,
		file:   file,
		index:  -1,
	}
}

func (parser *Parser) advance() {
	if parser.currentToken.TType != token.EOF {
		parser.index++
		parser.currentToken = parser.tokens[parser.index]
	}
}

func (parser *Parser) consume(tType token.TokenType) error {
	if parser.currentToken.TType != tType && parser.currentToken.TType != token.EOF {
		pos := parser.currentToken.Pos
		tType2 := parser.currentToken.TType

		parser.advance()

		return snowerror.NewSnowError(
			snowerror.EXPECTED_TOKEN_ERROR,
			fmt.Sprintf("expected token of type '%s', not token of type '%s'", tType, tType2),
			"",
			pos,
		)
	}

	parser.advance()

	return nil
}

func (parser *Parser) binary(t1 token.TokenType, t2 token.TokenType, t3 token.TokenType, t4 token.TokenType, function func() (parsevals.Expr, error)) (parsevals.Expr, error) {
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

		left = parsevals.NewBinaryExpr(
			left,
			right,
			opToken,
			*startPos.CreateSEPos(right.GetPosition().End, parser.currentToken.Pos.File),
		)
	}

	return left, nil
}

func (parser *Parser) Parse() ([]parsevals.Stmt, error) {
	statements := make([]parsevals.Stmt, 0)

	parser.advance()

	for parser.currentToken.TType != token.EOF {
		stmt, err := parser.deceleration()
		if err != nil {
			return nil, err
		}

		statements = append(statements, stmt)
	}

	return statements, nil
}

func (parser *Parser) deceleration() (parsevals.Stmt, error) {
	if parser.currentToken.TType == token.VAR || parser.currentToken.TType == token.CONST {
		varDeclStmt, err := parser.varDeclStmt()
		if err != nil {
			return nil, err
		}

		return varDeclStmt, nil
	}

	statement, err := parser.statement()
	if err != nil {
		return nil, err
	}

	return statement, nil
}

func (parser *Parser) varDeclStmt() (parsevals.Stmt, error) {
	startTok := parser.currentToken

	parser.advance()

	if parser.currentToken.TType != token.IDENTIFIER {
		return nil, snowerror.NewUnexpectedTokenError(token.IDENTIFIER, parser.currentToken)
	}

	identifier := parser.currentToken

	parser.advance()

	err := parser.consume(token.SINGLE_EQUALS)
	if err != nil {
		return nil, err
	}

	endPos := parser.currentToken.Pos.End

	expr, err := parser.expression()
	if err != nil {
		return nil, err
	}

	return parsevals.NewVarDeclStmt(startTok, identifier, expr, *startTok.Pos.Start.CreateSEPos(endPos, startTok.Pos.File)), nil
}

func (parser *Parser) statement() (parsevals.Stmt, error) {
	statement, err := parser.expressionStmt()
	if err != nil {
		return nil, err
	}

	return statement, nil
}

func (parser *Parser) expressionStmt() (parsevals.Stmt, error) {
	pos := parser.currentToken.Pos

	expression, err := parser.expression()
	if err != nil {
		return nil, err
	}

	if parser.currentToken.TType != token.EOF {
		err = parser.consume(token.NEWLINE)
		if err != nil {
			return nil, err
		}
	}

	return parsevals.NewExpressionStmt(expression, pos), nil
}

func (parser *Parser) expression() (parsevals.Expr, error) {
	assignment, err := parser.assignment()
	if err != nil {
		return nil, err
	}

	return assignment, nil
}

func (parser *Parser) assignment() (parsevals.Expr, error) {
	logicOr, err := parser.logicOr()
	if err != nil {
		return nil, err
	}

	return logicOr, nil
}

func (parser *Parser) logicOr() (parsevals.Expr, error) {
	logicAnd, err := parser.logicAnd()
	if err != nil {
		return nil, err
	}

	return logicAnd, nil
}

func (parser *Parser) logicAnd() (parsevals.Expr, error) {
	equality, err := parser.equality()
	if err != nil {
		return nil, err
	}

	return equality, nil
}

func (parser *Parser) equality() (parsevals.Expr, error) {
	binary, err := parser.binary(
		token.EQUALS,
		token.NOT_EQUALS,
		token.PLACEHOLDER,
		token.PLACEHOLDER,
		parser.comparison,
	)
	if err != nil {
		return nil, err
	}

	return binary, nil
}

func (parser *Parser) comparison() (parsevals.Expr, error) {
	binary, err := parser.binary(
		token.GREATER_THAN,
		token.GREATER_THAN_EQUALS,
		token.LESS_THAN,
		token.LESS_THAN_EQUALS,
		parser.term,
	)
	if err != nil {
		return nil, err
	}

	return binary, nil
}

func (parser *Parser) term() (parsevals.Expr, error) {
	binary, err := parser.binary(
		token.PLUS,
		token.DASH,
		token.PLACEHOLDER,
		token.PLACEHOLDER,
		parser.factor,
	)
	if err != nil {
		return nil, err
	}

	return binary, nil
}

func (parser *Parser) factor() (parsevals.Expr, error) {
	binary, err := parser.binary(
		token.STAR,
		token.SLASH,
		token.PLACEHOLDER,
		token.PLACEHOLDER,
		parser.unary,
	)
	if err != nil {
		return nil, err
	}

	return binary, nil
}

func (parser *Parser) unary() (parsevals.Expr, error) {
	if parser.currentToken.TType == token.DASH || parser.currentToken.TType == token.NOT {
		token := parser.currentToken

		parser.advance()

		right, err := parser.unary()
		if err != nil {
			return nil, err
		}

		return parsevals.NewUnaryExpr(
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

func (parser *Parser) call() (parsevals.Expr, error) {
	primary, err := parser.primary()
	if err != nil {
		return nil, err
	}

	return primary, err
}

func (parser *Parser) primary() (parsevals.Expr, error) {
	startToken := parser.currentToken

	switch parser.currentToken.TType {
	case token.INT:
		parser.advance()

		intValue, err := strconv.Atoi(startToken.Value)
		if err != nil {
			return nil, snowerror.NewSnowError(
				snowerror.TOO_BIG_VALUE_ERROR,
				fmt.Sprintf("the value of number of type %s is too big", startToken.TType),
				"",
				startToken.Pos,
			)
		}

		return parsevals.NewIntLiteralExpr(intValue, parser.currentToken.Pos), nil
	case token.FLOAT:
		parser.advance()

		floatValue, err := strconv.ParseFloat(startToken.Value, 64)
		if err != nil {
			return nil, snowerror.NewSnowError(
				snowerror.TOO_BIG_VALUE_ERROR,
				fmt.Sprintf("the value of number of type %s is too big", startToken.TType),
				"",
				startToken.Pos,
			)
		}

		return parsevals.NewFloatLiteralExpr(floatValue, parser.currentToken.Pos), nil
	case token.TRUE:
		parser.advance()

		return parsevals.NewBoolLiteralExpr(true, startToken.Pos), nil
	case token.FALSE:
		parser.advance()

		return parsevals.NewBoolLiteralExpr(false, startToken.Pos), nil
	case token.LPAREN:
		parser.advance()

		expr, err := parser.expression()
		if err != nil {
			return nil, err
		}

		endPos := parser.currentToken.Pos.End

		parser.consume(token.RPAREN)

		pos := startToken.Pos.Start.CreateSEPos(endPos, startToken.Pos.File)

		return parsevals.NewGroupingExpr(expr, *pos), nil
	default:
		err := snowerror.NewSnowError(
			snowerror.INVALID_TOKEN_TYPE_ERROR,
			fmt.Sprintf("token of type '%s' is invalid", parser.currentToken.TType),
			"",
			parser.currentToken.Pos,
		)
		return nil, err
	}
}
