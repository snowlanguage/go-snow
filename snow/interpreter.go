package snow

import (
	"fmt"
)

type Interpreter struct {
	statements   []Stmt
	file         *File
	currentStmt  Stmt
	index        int
	end          bool
	environment  *Environment
	inLoop       int
	inFunc       int
	continueLoop bool
	breakLoop    bool
	returnVal    RTValue
}

func NewInterpreter(statements []Stmt, file *File, env *Environment) *Interpreter {
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

func (interpreter *Interpreter) Interpret() ([]RTValue, error) {
	values := make([]RTValue, 0)

	interpreter.advance()

	for _, stmt := range interpreter.statements {
		value, err := interpreter.execute(stmt, interpreter.environment)
		if err != nil {
			return nil, err
		}

		if value != nil {
			values = append(values, value)
		}
	}

	return values, nil
}

func (interpreter *Interpreter) execute(statement Stmt, env *Environment) (RTValue, error) {
	return statement.Accept(interpreter, env)
}

func (interpreter *Interpreter) evaluate(expression Expr, env *Environment) (RTValue, error) {
	return expression.Accept(interpreter, env)
}

func (interpreter *Interpreter) VisitExpressionStmt(stmt ExpressionStmt, env *Environment) (RTValue, error) {
	value, err := interpreter.evaluate(stmt.Expression, env)
	if err != nil {
		return nil, err
	}

	return value, nil
}

func (interpreter *Interpreter) VisitVarDeclStmt(stmt VarDeclStmt, env *Environment) (RTValue, error) {
	value, err := interpreter.evaluate(stmt.Expression, env)
	if err != nil {
		return nil, err
	}

	err = env.Declare(
		stmt.VarType.TType == CONST,
		stmt.Identifier.Value,
		value,
		stmt.Pos,
	)
	if err != nil {
		return nil, err
	}

	return value, nil
}

func (interpreter *Interpreter) VisitBlockStmt(stmt BlockStmt, env *Environment, newEnv bool) (RTValue, error) {
	blockEnv := NewEnvironment(env, stmt.Name, stmt.Pos.Start.Ln, stmt.Pos.File.Name, false)

	for _, statement := range stmt.Statements {
		var err error
		if newEnv {
			_, err = interpreter.execute(statement, blockEnv)
		} else {
			_, err = interpreter.execute(statement, blockEnv)
		}

		if interpreter.breakLoop || interpreter.continueLoop {
			return nil, nil
		}

		if err != nil {
			return nil, err
		}
	}

	return nil, nil
}

func (interpreter *Interpreter) VisitWhileStmt(stmt WhileStmt, env *Environment) (RTValue, error) {
	exprVisited, err := interpreter.evaluate(stmt.Expression, env)
	if err != nil {
		return nil, err
	}

	exprBool, err := exprVisited.ToBool(stmt.Expression.GetPosition())
	if err != nil {
		return nil, err
	}

	interpreter.inLoop += 1

	for exprBool.GetValue() == true {
		_, err := interpreter.execute(stmt.Statement, env)
		if err != nil {
			return nil, err
		}

		if interpreter.continueLoop {
			interpreter.continueLoop = false
			continue
		} else if interpreter.breakLoop {
			interpreter.breakLoop = false
			break
		}

		exprVisited, err = interpreter.evaluate(stmt.Expression, env)
		if err != nil {
			return nil, err
		}

		exprBool, err = exprVisited.ToBool(stmt.Expression.GetPosition())
		if err != nil {
			return nil, err
		}
	}

	interpreter.inLoop -= 1

	return nil, nil
}

func (interpreter *Interpreter) VisitFunctionDeclStmt(stmt FunctionDeclStmt, env *Environment) (RTValue, error) {
	rTFunc := NewRTFunction(stmt.Name, stmt.Parameters, stmt.Block, stmt.Pos, env)

	err := env.Declare(true, stmt.Name, rTFunc, rTFunc.Pos)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (interpreter *Interpreter) VisitBreakStmt(stmt BreakStmt, env *Environment) (RTValue, error) {
	interpreter.breakLoop = true
	return nil, nil
}

func (interpreter *Interpreter) VisitContinueStmt(stmt ContinueStmt, env *Environment) (RTValue, error) {
	interpreter.continueLoop = true
	return nil, nil
}

func (interpreter *Interpreter) VisitBinaryExpr(expr BinaryExpr, env *Environment) (RTValue, error) {
	left, err := interpreter.evaluate(expr.Left, env)
	if err != nil {
		return nil, err
	}

	right, err := interpreter.evaluate(expr.Right, env)
	if err != nil {
		return nil, err
	}

	switch expr.Tok.TType {
	case PLUS:
		return left.Add(right, expr.Pos)
	case DASH:
		return left.Subtract(right, expr.Pos)
	case STAR:
		return left.Multiply(right, expr.Pos)
	case SLASH:
		return left.Divide(right, expr.Pos)
	case EQUALS:
		return left.Equals(right, expr.Pos)
	case NOT_EQUALS:
		return left.NotEquals(right, expr.Pos)
	case GREATER_THAN:
		return left.GreaterThan(right, expr.Pos)
	case GREATER_THAN_EQUALS:
		return left.GreaterThanEquals(right, expr.Pos)
	case LESS_THAN:
		return left.LessThan(right, expr.Pos)
	case LESS_THAN_EQUALS:
		return left.LessThanEquals(right, expr.Pos)
	default:
		return nil, NewSnowError(
			INVALID_OP_TOKEN_ERROR,
			fmt.Sprintf("the op token '%s' is not valid", expr.Tok.TType),
			"",
			expr.Pos,
		)
	}
}

func (interpreter *Interpreter) VisitUnaryExpr(expr UnaryExpr, env *Environment) (RTValue, error) {
	right, err := interpreter.evaluate(expr.Right, env)
	if err != nil {
		return nil, err
	}

	if expr.Tok.TType == DASH {
		val, err := right.Multiply(NewRTInt(expr.Pos, -1, env), expr.Pos)
		if err != nil {
			return nil, err
		}

		return val, nil
	} else if expr.Tok.TType == NOT {
		val, err := right.Not(expr.Pos)
		if err != nil {
			return nil, err
		}

		return val, nil
	}

	return nil, NewRuntimeError(
		INVALID_OP_TOKEN_ERROR,
		fmt.Sprintf("the op token of type '%s' is invalid for unary expressions", expr.Tok.TType),
		"",
		expr.Pos,
		env,
	)
}

func (interpreter *Interpreter) VisitGroupingExpr(expr GroupingExpr, env *Environment) (RTValue, error) {
	val, err := interpreter.evaluate(expr.Expression, env)
	if err != nil {
		return nil, err
	}

	return val, nil
}

func (interpreter *Interpreter) VisitIntLiteralExpr(expr IntLiteralExpr, env *Environment) (RTValue, error) {
	return NewRTInt(expr.Pos, expr.Value, env), nil
}

func (interpreter *Interpreter) VisitFloatLiteralExpr(expr FloatLiteralExpr, env *Environment) (RTValue, error) {
	return NewRTFloat(expr.Pos, expr.Value, env), nil
}

func (interpreter *Interpreter) VisitBoolLiteralExpr(expr BoolLiteralExpr, env *Environment) (RTValue, error) {
	return NewRTBool(expr.Pos, expr.Value, env), nil
}

func (interpreter *Interpreter) VisitVarAccessExpr(expr VarAccessExpr, env *Environment) (RTValue, error) {
	val, err := env.Get(expr.Value, expr.Pos, env)
	if err != nil {
		return nil, err
	}

	return val, nil
}

func (interpreter *Interpreter) VisitVarAssignmentExpr(expr VarAssignmentExpr, env *Environment) (RTValue, error) {
	if expr.Object == nil {
		val, err := interpreter.evaluate(expr.Value, env)
		if err != nil {
			return nil, err
		}

		return env.Set(expr.Name, val, env, expr.Pos)
	} else {
		left, err := interpreter.evaluate(expr.Object, env)
		if err != nil {
			return nil, err
		}

		val, err := interpreter.evaluate(expr.Value, env)
		if err != nil {
			return nil, err
		}

		return left.SetAttribute(expr.Name, val, expr.Pos)
	}
}

func (interpreter *Interpreter) VisitDotExpr(expr DotExpr, env *Environment) (RTValue, error) {
	left, err := interpreter.evaluate(expr.Left, env)
	if err != nil {
		return nil, err
	}

	val, err := left.Dot(expr.Right, expr.Pos)
	if err != nil {
		return nil, err
	}

	return val, nil
}

func (interpreter *Interpreter) VisitCallExpr(expr CallExpr, env *Environment) (RTValue, error) {
	function, err := interpreter.evaluate(expr.Function, env)
	if err != nil {
		return nil, err
	}

	arguments := make([]RTValue, 0)
	for _, arg := range expr.Arguments {
		argVisited, err := interpreter.evaluate(arg, env)
		if err != nil {
			return nil, err
		}

		arguments = append(arguments, argVisited)
	}

	fmt.Printf("arguments: %v\n", arguments)

	interpreter.inFunc += 1

	val, err := function.Call(arguments, expr.Pos, interpreter)
	if err != nil {
		return nil, err
	}

	interpreter.inFunc -= 1

	return val, nil
}
