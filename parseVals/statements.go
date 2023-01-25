package parsevals

import (
	"fmt"

	"github.com/snowlanguage/go-snow/position"
	"github.com/snowlanguage/go-snow/runtimevalues"
	"github.com/snowlanguage/go-snow/token"
)

type Stmt interface {
	Accept(visitor StmtVisitor, env *runtimevalues.Environment) (runtimevalues.RTValue, error)
	ToString() string
}

type StmtVisitor interface {
	VisitExpressionStmt(stmt ExpressionStmt, env *runtimevalues.Environment) (runtimevalues.RTValue, error)
	VisitVarDeclStmt(stmt VarDeclStmt, env *runtimevalues.Environment) (runtimevalues.RTValue, error)
}

type ExpressionStmt struct {
	Expression Expr
	Pos        position.SEPos
}

func NewExpressionStmt(expression Expr, pos position.SEPos) *ExpressionStmt {
	return &ExpressionStmt{
		Expression: expression,
		Pos:        pos,
	}
}

func (expressionStmt ExpressionStmt) Accept(visitor StmtVisitor, env *runtimevalues.Environment) (runtimevalues.RTValue, error) {
	return visitor.VisitExpressionStmt(expressionStmt, env)
}

func (expressionStmt ExpressionStmt) ToString() string {
	return fmt.Sprintf("(EXPRESSION: %s)", expressionStmt.Expression.ToString())
}

type VarDeclStmt struct {
	VarType    token.Token
	Identifier token.Token
	Expression Expr
	Pos        position.SEPos
}

func NewVarDeclStmt(varType token.Token, identifier token.Token, expression Expr, pos position.SEPos) *VarDeclStmt {
	return &VarDeclStmt{
		VarType:    varType,
		Identifier: identifier,
		Expression: expression,
		Pos:        pos,
	}
}

func (varDeclStmt VarDeclStmt) Accept(visitor StmtVisitor, env *runtimevalues.Environment) (runtimevalues.RTValue, error) {
	return visitor.VisitVarDeclStmt(varDeclStmt, env)
}

func (varDeclStmt VarDeclStmt) ToString() string {
	return fmt.Sprintf("(VAR_DECL_STMT: %s %s = %s)", varDeclStmt.VarType.Value, varDeclStmt.Identifier.ToString(), varDeclStmt.Expression.ToString())
}
