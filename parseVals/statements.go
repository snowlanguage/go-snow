package parsevals

import (
	"fmt"

	"github.com/snowlanguage/go-snow/runtimevalues"
)

type Stmt interface {
	Accept(visitor StmtVisitor) (runtimevalues.RTValue, error)
	ToString() string
}

type StmtVisitor interface {
	VisitExpressionStmt(stmt ExpressionStmt) (runtimevalues.RTValue, error)
}

type ExpressionStmt struct {
	Expression Expr
}

func NewExpressionStmt(expression Expr) *ExpressionStmt {
	return &ExpressionStmt{
		Expression: expression,
	}
}

func (expressionStmt ExpressionStmt) Accept(visitor StmtVisitor) (runtimevalues.RTValue, error) {
	return visitor.VisitExpressionStmt(expressionStmt)
}

func (expressionStmt ExpressionStmt) ToString() string {
	return fmt.Sprintf("(EXPRESSION: %s)", expressionStmt.Expression.ToString())
}
