package parsevals

import "fmt"

type Stmt interface {
	Accept(visitor StmtVisitor) interface{}
	ToString() string
}

type StmtVisitor interface {
	visitExpressionStmt(stmt *ExpressionStmt) interface{}
}

type ExpressionStmt struct {
	Expression Expr
}

func NewExpressionStmt(expression Expr) *ExpressionStmt {
	return &ExpressionStmt{
		Expression: expression,
	}
}

func (expressionStmt *ExpressionStmt) Accept(visitor StmtVisitor) interface{} {
	return visitor.visitExpressionStmt(expressionStmt)
}

func (expressionStmt *ExpressionStmt) ToString() string {
	return fmt.Sprintf("(EXPRESSION: %s)", expressionStmt.Expression.ToString())
}
