package parser

import (
	"fmt"
	"log"
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
	if parser.currentToken.TType != tType {
		return snowerror.NewSnowError(
			snowerror.EXPECTED_TOKEN_ERROR,
			fmt.Sprintf("expected token of type '%s', not token of type '%s'", tType, parser.currentToken.TType),
			"",
			parser.currentToken.Pos,
		)
	}

	return nil
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
	statement, err := parser.statement()
	if err != nil {
		return nil, err
	}

	return statement, nil
}

func (parser *Parser) statement() (parsevals.Stmt, error) {
	statement, err := parser.expressionStmt()
	if err != nil {
		return nil, err
	}

	return statement, nil
}

func (parser *Parser) expressionStmt() (parsevals.Stmt, error) {
	expression, err := parser.expression()
	if err != nil {
		return nil, err
	}

	parser.advance()

	if parser.currentToken.TType != token.EOF {
		err = parser.consume(token.NEWLINE)
		if err != nil {
			return nil, err
		}
	}

	return parsevals.NewExpressionStmt(expression), nil
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
	comparison, err := parser.comparison()
	if err != nil {
		return nil, err
	}

	return comparison, nil
}

func (parser *Parser) comparison() (parsevals.Expr, error) {
	term, err := parser.term()
	if err != nil {
		return nil, err
	}

	return term, nil
}

func (parser *Parser) term() (parsevals.Expr, error) {
	factor, err := parser.factor()
	if err != nil {
		return nil, err
	}

	return factor, nil
}

func (parser *Parser) factor() (parsevals.Expr, error) {
	unary, err := parser.unary()
	if err != nil {
		return nil, err
	}

	return unary, nil
}

func (parser *Parser) unary() (parsevals.Expr, error) {
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
	switch parser.currentToken.TType {
	case token.INT:
		intValue, err := strconv.Atoi(parser.currentToken.Value)
		if err != nil {
			log.Panic(err)
		}

		return parsevals.NewIntLiteralExpr(intValue, parser.currentToken.Pos), nil
	default:
		err := snowerror.NewSnowError(
			snowerror.INVALID_TOKEN_TYPE,
			fmt.Sprintf("token of type '%s' is invalid", parser.currentToken.TType),
			"",
			parser.currentToken.Pos,
		)
		return nil, err
	}
}