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
	GetPos() position.SEPos
}

type StmtVisitor interface {
	VisitExpressionStmt(stmt ExpressionStmt, env *runtimevalues.Environment) (runtimevalues.RTValue, error)
	VisitVarDeclStmt(stmt VarDeclStmt, env *runtimevalues.Environment) (runtimevalues.RTValue, error)
	VisitBlockStmt(stmt BlockStmt, env *runtimevalues.Environment) (runtimevalues.RTValue, error)
	VisitWhileStmt(stmt WhileStmt, env *runtimevalues.Environment) (runtimevalues.RTValue, error)
	VisitBreakStmt(stmt BreakStmt, env *runtimevalues.Environment) (runtimevalues.RTValue, error)
	VisitContinueStmt(stmt ContinueStmt, env *runtimevalues.Environment) (runtimevalues.RTValue, error)
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

func (expressionStmt ExpressionStmt) GetPos() position.SEPos {
	return expressionStmt.Pos
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

func (varDeclStmt VarDeclStmt) GetPos() position.SEPos {
	return varDeclStmt.Pos
}

type BlockStmt struct {
	Statements []Stmt
	Name       string
	Pos        position.SEPos
}

func NewBlockStmt(statements []Stmt, name string, pos position.SEPos) *BlockStmt {
	return &BlockStmt{
		Statements: statements,
		Name:       name,
		Pos:        pos,
	}
}

func (blockStmt BlockStmt) Accept(visitor StmtVisitor, env *runtimevalues.Environment) (runtimevalues.RTValue, error) {
	return visitor.VisitBlockStmt(blockStmt, env)
}

func (blockStmt BlockStmt) ToString() string {
	s := "["
	for _, v := range blockStmt.Statements {
		s += v.ToString() + " "
	}
	s += "]"

	return fmt.Sprintf("(BLOCK_STMT(%s): %s)", blockStmt.Name, s)
}

func (blockStmt BlockStmt) GetPos() position.SEPos {
	return blockStmt.Pos
}

type WhileStmt struct {
	Statement  Stmt
	Expression Expr
	Pos        position.SEPos
}

func NewWhileStmt(statement Stmt, expression Expr, pos position.SEPos) *WhileStmt {
	return &WhileStmt{
		Statement:  statement,
		Expression: expression,
		Pos:        pos,
	}
}

func (whileStmt WhileStmt) Accept(visitor StmtVisitor, env *runtimevalues.Environment) (runtimevalues.RTValue, error) {
	return visitor.VisitWhileStmt(whileStmt, env)
}

func (whileStmt WhileStmt) ToString() string {
	return fmt.Sprintf("(WHILE_STMT: %s %s)", whileStmt.Expression.ToString(), whileStmt.Statement.ToString())
}

func (whileStmt WhileStmt) GetPos() position.SEPos {
	return whileStmt.Pos
}

type BreakStmt struct {
	Pos position.SEPos
}

func NewBreakStmt(pos position.SEPos) *BreakStmt {
	return &BreakStmt{
		Pos: pos,
	}
}

func (breakStmt BreakStmt) Accept(visitor StmtVisitor, env *runtimevalues.Environment) (runtimevalues.RTValue, error) {
	return visitor.VisitBreakStmt(breakStmt, env)
}

func (breakStmt BreakStmt) ToString() string {
	return "(BREAK_STMT)"
}

func (breakStmt BreakStmt) GetPos() position.SEPos {
	return breakStmt.Pos
}

type ContinueStmt struct {
	Pos position.SEPos
}

func NewContinueStmt(pos position.SEPos) *ContinueStmt {
	return &ContinueStmt{
		Pos: pos,
	}
}

func (continueStmt ContinueStmt) Accept(visitor StmtVisitor, env *runtimevalues.Environment) (runtimevalues.RTValue, error) {
	return visitor.VisitContinueStmt(continueStmt, env)
}

func (continueStmt ContinueStmt) ToString() string {
	return "(CONTINUE_STMT)"
}

func (continueStmt ContinueStmt) GetPos() position.SEPos {
	return continueStmt.Pos
}
