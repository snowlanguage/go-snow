package interpreter

import (
	"fmt"

	"github.com/snowlanguage/go-snow/file"
	parsevals "github.com/snowlanguage/go-snow/parseVals"
	"github.com/snowlanguage/go-snow/runtimevalues"
	snowerror "github.com/snowlanguage/go-snow/snowError"
	"github.com/snowlanguage/go-snow/token"
)

type Interpreter struct {
	statements  []parsevals.Stmt
	file        *file.File
	currentStmt parsevals.Stmt
	index       int
	end         bool
}

func NewInterpreter(statements []parsevals.Stmt, file *file.File) *Interpreter {
	return &Interpreter{
		statements: statements,
		file:       file,
		index:      -1,
	}
}

func (interpreter *Interpreter) advance() {
	if !interpreter.end {
		interpreter.index++
	}

	if len(interpreter.statements) > interpreter.index {
		interpreter.currentStmt = interpreter.statements[interpreter.index]
	}
}

func (interpreter *Interpreter) Interpret() ([]runtimevalues.RTValue, error) {
	values := make([]runtimevalues.RTValue, 0)

	interpreter.advance()

	for _, stmt := range interpreter.statements {
		fmt.Println("a")
		value, err := interpreter.execute(stmt)
		if err != nil {
			return nil, err
		}

		values = append(values, value)
	}

	return values, nil
}

func (interpreter *Interpreter) execute(statement parsevals.Stmt) (runtimevalues.RTValue, error) {
	return statement.Accept(interpreter)
}

func (interpreter *Interpreter) evaluate(expression parsevals.Expr) (runtimevalues.RTValue, error) {
	return expression.Accept(interpreter)
}

func (interpreter *Interpreter) VisitExpressionStmt(stmt parsevals.ExpressionStmt) (runtimevalues.RTValue, error) {
	value, err := interpreter.evaluate(stmt.Expression)
	if err != nil {
		return nil, err
	}

	return value, nil
}

func (interpreter *Interpreter) VisitBinaryExpr(expr parsevals.BinaryExpr) (runtimevalues.RTValue, error) {
	left, err := interpreter.evaluate(expr.Left)
	if err != nil {
		return nil, err
	}

	right, err := interpreter.evaluate(expr.Right)
	if err != nil {
		return nil, err
	}

	switch expr.Tok.TType {
	case token.PLUS:
		return left.Add(right, expr.Pos)
	default:
		return nil, snowerror.NewSnowError(
			snowerror.INVALID_OP_TOKEN_ERROR,
			fmt.Sprintf("the op token '%s' is not valid", expr.Tok.TType),
			"",
			expr.Pos,
		)
	}
}

func (interpreter *Interpreter) VisitIntLiteralExpr(expr parsevals.IntLiteralExpr) (runtimevalues.RTValue, error) {
	return runtimevalues.NewRTInt(expr.Pos, expr.Value), nil
}