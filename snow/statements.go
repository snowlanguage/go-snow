package snow

import (
	"fmt"
)

type Stmt interface {
	Accept(visitor StmtVisitor, env *Environment) (RTValue, error)
	ToString() string
	GetPos() SEPos
}

type StmtVisitor interface {
	VisitExpressionStmt(stmt ExpressionStmt, env *Environment) (RTValue, error)
	VisitVarDeclStmt(stmt VarDeclStmt, env *Environment) (RTValue, error)
	VisitBlockStmt(stmt BlockStmt, env *Environment, newEnv bool) (RTValue, error)
	VisitWhileStmt(stmt WhileStmt, env *Environment) (RTValue, error)
	VisitBreakStmt(stmt BreakStmt, env *Environment) (RTValue, error)
	VisitContinueStmt(stmt ContinueStmt, env *Environment) (RTValue, error)
	VisitFunctionDeclStmt(stmt FunctionDeclStmt, env *Environment) (RTValue, error)
	VisitReturnStmt(stmt ReturnStmt, env *Environment) (RTValue, error)
	VisitIfStmt(stmt IfStmt, env *Environment) (RTValue, error)
	VisitIfStmtContainer(stmt IfStmtContainer, env *Environment) (RTValue, error)
}

type ExpressionStmt struct {
	Expression Expr
	Pos        SEPos
}

func NewExpressionStmt(expression Expr, pos SEPos) *ExpressionStmt {
	return &ExpressionStmt{
		Expression: expression,
		Pos:        pos,
	}
}

func (expressionStmt ExpressionStmt) Accept(visitor StmtVisitor, env *Environment) (RTValue, error) {
	return visitor.VisitExpressionStmt(expressionStmt, env)
}

func (expressionStmt ExpressionStmt) ToString() string {
	return fmt.Sprintf("(EXPRESSION: %s)", expressionStmt.Expression.ToString())
}

func (expressionStmt ExpressionStmt) GetPos() SEPos {
	return expressionStmt.Pos
}

type VarDeclStmt struct {
	VarType    Token
	Identifier Token
	Expression Expr
	Pos        SEPos
}

func NewVarDeclStmt(varType Token, identifier Token, expression Expr, pos SEPos) *VarDeclStmt {
	return &VarDeclStmt{
		VarType:    varType,
		Identifier: identifier,
		Expression: expression,
		Pos:        pos,
	}
}

func (varDeclStmt VarDeclStmt) Accept(visitor StmtVisitor, env *Environment) (RTValue, error) {
	return visitor.VisitVarDeclStmt(varDeclStmt, env)
}

func (varDeclStmt VarDeclStmt) ToString() string {
	return fmt.Sprintf("(VAR_DECL_STMT: %s %s = %s)", varDeclStmt.VarType.Value, varDeclStmt.Identifier.ToString(), varDeclStmt.Expression.ToString())
}

func (varDeclStmt VarDeclStmt) GetPos() SEPos {
	return varDeclStmt.Pos
}

type BlockStmt struct {
	Statements []Stmt
	Name       string
	Pos        SEPos
}

func NewBlockStmt(statements []Stmt, name string, pos SEPos) *BlockStmt {
	return &BlockStmt{
		Statements: statements,
		Name:       name,
		Pos:        pos,
	}
}

func (blockStmt BlockStmt) Accept(visitor StmtVisitor, env *Environment) (RTValue, error) {
	return visitor.VisitBlockStmt(blockStmt, env, true)
}

func (blockStmt BlockStmt) ToString() string {
	s := "["
	for _, v := range blockStmt.Statements {
		s += v.ToString() + " "
	}
	s += "]"

	return fmt.Sprintf("(BLOCK_STMT(%s): %s)", blockStmt.Name, s)
}

func (blockStmt BlockStmt) GetPos() SEPos {
	return blockStmt.Pos
}

type WhileStmt struct {
	Statement  Stmt
	Expression Expr
	Pos        SEPos
}

func NewWhileStmt(statement Stmt, expression Expr, pos SEPos) *WhileStmt {
	return &WhileStmt{
		Statement:  statement,
		Expression: expression,
		Pos:        pos,
	}
}

func (whileStmt WhileStmt) Accept(visitor StmtVisitor, env *Environment) (RTValue, error) {
	return visitor.VisitWhileStmt(whileStmt, env)
}

func (whileStmt WhileStmt) ToString() string {
	return fmt.Sprintf("(WHILE_STMT: %s %s)", whileStmt.Expression.ToString(), whileStmt.Statement.ToString())
}

func (whileStmt WhileStmt) GetPos() SEPos {
	return whileStmt.Pos
}

type BreakStmt struct {
	Pos SEPos
}

func NewBreakStmt(pos SEPos) *BreakStmt {
	return &BreakStmt{
		Pos: pos,
	}
}

func (breakStmt BreakStmt) Accept(visitor StmtVisitor, env *Environment) (RTValue, error) {
	return visitor.VisitBreakStmt(breakStmt, env)
}

func (breakStmt BreakStmt) ToString() string {
	return "(BREAK_STMT)"
}

func (breakStmt BreakStmt) GetPos() SEPos {
	return breakStmt.Pos
}

type ContinueStmt struct {
	Pos SEPos
}

func NewContinueStmt(pos SEPos) *ContinueStmt {
	return &ContinueStmt{
		Pos: pos,
	}
}

func (continueStmt ContinueStmt) Accept(visitor StmtVisitor, env *Environment) (RTValue, error) {
	return visitor.VisitContinueStmt(continueStmt, env)
}

func (continueStmt ContinueStmt) ToString() string {
	return "(CONTINUE_STMT)"
}

func (continueStmt ContinueStmt) GetPos() SEPos {
	return continueStmt.Pos
}

type FunctionDeclStmt struct {
	Name       string
	Parameters []Token
	Block      *BlockStmt
	Pos        SEPos
}

func NewFunctionDeclStmt(name string, parameters []Token, block *BlockStmt, pos SEPos) *FunctionDeclStmt {
	return &FunctionDeclStmt{
		Name:       name,
		Parameters: parameters,
		Block:      block,
		Pos:        pos,
	}
}

func (functionDeclStmt FunctionDeclStmt) Accept(visitor StmtVisitor, env *Environment) (RTValue, error) {
	return visitor.VisitFunctionDeclStmt(functionDeclStmt, env)
}

func (functionDeclStmt FunctionDeclStmt) ToString() string {
	p := "["
	for _, param := range functionDeclStmt.Parameters {
		p += param.ToString() + " "
	}
	p += "]"

	return fmt.Sprintf("(FUNCTION_DECL_STMT: %s %s %s)", functionDeclStmt.Name, p, functionDeclStmt.Block.ToString())
}

func (functionDeclStmt FunctionDeclStmt) GetPos() SEPos {
	return functionDeclStmt.Pos
}

type ReturnStmt struct {
	Value Expr
	Pos   SEPos
}

func NewReturnStmt(value Expr, pos SEPos) *ReturnStmt {
	return &ReturnStmt{
		Value: value,
		Pos:   pos,
	}
}

func (returnStmt ReturnStmt) Accept(visitor StmtVisitor, env *Environment) (RTValue, error) {
	return visitor.VisitReturnStmt(returnStmt, env)
}

func (returnStmt ReturnStmt) ToString() string {
	return fmt.Sprintf("(RETURN_STMT: %s)", returnStmt.Value.ToString())
}

func (returnStmt ReturnStmt) GetPos() SEPos {
	return returnStmt.Pos
}

type IfStmt struct {
	Expression Expr
	Statement  Stmt
	Pos        SEPos
}

func NewIfStmt(expression Expr, statement Stmt, pos SEPos) *IfStmt {
	return &IfStmt{
		Expression: expression,
		Statement:  statement,
		Pos:        pos,
	}
}

func (ifStmt IfStmt) Accept(visitor StmtVisitor, env *Environment) (RTValue, error) {
	return visitor.VisitIfStmt(ifStmt, env)
}

func (ifStmt IfStmt) ToString() string {
	return fmt.Sprintf("(IF_STMT: %s %s)", ifStmt.Expression.ToString(), ifStmt.Statement.ToString())
}

func (ifStmt IfStmt) GetPos() SEPos {
	return ifStmt.Pos
}

type IfStmtContainer struct {
	IfStmts []IfStmt
	Pos     SEPos
}

func NewIfStmtContainer(ifStmts []IfStmt, pos SEPos) *IfStmtContainer {
	return &IfStmtContainer{
		IfStmts: ifStmts,
		Pos:     pos,
	}
}

func (ifStmtContainer IfStmtContainer) Accept(visitor StmtVisitor, env *Environment) (RTValue, error) {
	return visitor.VisitIfStmtContainer(ifStmtContainer, env)
}

func (ifStmtContainer IfStmtContainer) ToString() string {
	return fmt.Sprintf("(ifStmtContainer)")
}

func (ifStmtContainer IfStmtContainer) GetPos() SEPos {
	return ifStmtContainer.Pos
}
