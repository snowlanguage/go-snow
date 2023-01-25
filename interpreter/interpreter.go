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
	environment *runtimevalues.Environment
}

func NewInterpreter(statements []parsevals.Stmt, file *file.File, env *runtimevalues.Environment) *Interpreter {
	return &Interpreter{
		statements:  statements,
		file:        file,
		index:       -1,
		environment: env,
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
		value, err := interpreter.execute(stmt, interpreter.environment)
		if err != nil {
			return nil, err
		}

		values = append(values, value)
	}

	return values, nil
}

func (interpreter *Interpreter) execute(statement parsevals.Stmt, env *runtimevalues.Environment) (runtimevalues.RTValue, error) {
	return statement.Accept(interpreter, env)
}

func (interpreter *Interpreter) evaluate(expression parsevals.Expr, env *runtimevalues.Environment) (runtimevalues.RTValue, error) {
	return expression.Accept(interpreter, env)
}

func (interpreter *Interpreter) VisitExpressionStmt(stmt parsevals.ExpressionStmt, env *runtimevalues.Environment) (runtimevalues.RTValue, error) {
	value, err := interpreter.evaluate(stmt.Expression, env)
	if err != nil {
		return nil, err
	}

	return value, nil
}

func (interpreter *Interpreter) VisitVarDeclStmt(stmt parsevals.VarDeclStmt, env *runtimevalues.Environment) (runtimevalues.RTValue, error) {
	value, err := interpreter.evaluate(stmt.Expression, env)
	if err != nil {
		return nil, err
	}

	err = env.Declare(
		stmt.VarType.TType == token.CONST,
		stmt.Identifier.Value,
		value,
		stmt.Pos,
	)
	if err != nil {
		return nil, err
	}

	return value, nil
}

func (interpreter *Interpreter) VisitBinaryExpr(expr parsevals.BinaryExpr, env *runtimevalues.Environment) (runtimevalues.RTValue, error) {
	left, err := interpreter.evaluate(expr.Left, env)
	if err != nil {
		return nil, err
	}

	right, err := interpreter.evaluate(expr.Right, env)
	if err != nil {
		return nil, err
	}

	switch expr.Tok.TType {
	case token.PLUS:
		return left.Add(right, expr.Pos)
	case token.DASH:
		return left.Subtract(right, expr.Pos)
	case token.STAR:
		return left.Multiply(right, expr.Pos)
	case token.SLASH:
		return left.Divide(right, expr.Pos)
	case token.EQUALS:
		return left.Equals(right, expr.Pos)
	case token.NOT_EQUALS:
		return left.NotEquals(right, expr.Pos)
	case token.GREATER_THAN:
		return left.GreaterThan(right, expr.Pos)
	case token.GREATER_THAN_EQUALS:
		return left.GreaterThanEquals(right, expr.Pos)
	case token.LESS_THAN:
		return left.LessThan(right, expr.Pos)
	case token.LESS_THAN_EQUALS:
		return left.LessThanEquals(right, expr.Pos)
	default:
		return nil, snowerror.NewSnowError(
			snowerror.INVALID_OP_TOKEN_ERROR,
			fmt.Sprintf("the op token '%s' is not valid", expr.Tok.TType),
			"",
			expr.Pos,
		)
	}
}

func (interpreter *Interpreter) VisitUnaryExpr(expr parsevals.UnaryExpr, env *runtimevalues.Environment) (runtimevalues.RTValue, error) {
	right, err := interpreter.evaluate(expr.Right, interpreter.environment)
	if err != nil {
		return nil, err
	}

	if expr.Tok.TType == token.DASH {
		val, err := right.Multiply(runtimevalues.NewRTInt(expr.Pos, -1, env), expr.Pos)
		if err != nil {
			return nil, err
		}

		return val, nil
	} else if expr.Tok.TType == token.NOT {
		val, err := right.Not(expr.Pos)
		if err != nil {
			return nil, err
		}

		return val, nil
	}

	return nil, runtimevalues.NewRuntimeError(
		snowerror.INVALID_OP_TOKEN_ERROR,
		fmt.Sprintf("the op token of type '%s' is invalid for unary expressions", expr.Tok.TType),
		"",
		expr.Pos,
		env,
	)
}

func (interpreter *Interpreter) VisitGroupingExpr(expr parsevals.GroupingExpr, env *runtimevalues.Environment) (runtimevalues.RTValue, error) {
	val, err := interpreter.evaluate(expr.Expression, env)
	if err != nil {
		return nil, err
	}

	return val, nil
}

func (interpreter *Interpreter) VisitIntLiteralExpr(expr parsevals.IntLiteralExpr, env *runtimevalues.Environment) (runtimevalues.RTValue, error) {
	return runtimevalues.NewRTInt(expr.Pos, expr.Value, env), nil
}

func (interpreter *Interpreter) VisitFloatLiteralExpr(expr parsevals.FloatLiteralExpr, env *runtimevalues.Environment) (runtimevalues.RTValue, error) {
	return runtimevalues.NewRTFloat(expr.Pos, expr.Value, env), nil
}

func (interpreter *Interpreter) VisitBoolLiteralExpr(expr parsevals.BoolLiteralExpr, env *runtimevalues.Environment) (runtimevalues.RTValue, error) {
	return runtimevalues.NewRTBool(expr.Pos, expr.Value, env), nil
}
