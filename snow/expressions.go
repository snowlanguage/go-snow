package snow

import (
	"fmt"
)

type Expr interface {
	Accept(visitor ExprVisitor, env *Environment) (RTValue, error)
	ToString() string
	GetPosition() SEPos
}

type ExprVisitor interface {
	VisitBinaryExpr(expr BinaryExpr, env *Environment) (RTValue, error)
	VisitUnaryExpr(expr UnaryExpr, env *Environment) (RTValue, error)
	VisitGroupingExpr(expr GroupingExpr, env *Environment) (RTValue, error)
	VisitIntLiteralExpr(expr IntLiteralExpr, env *Environment) (RTValue, error)
	VisitFloatLiteralExpr(expr FloatLiteralExpr, env *Environment) (RTValue, error)
	VisitBoolLiteralExpr(expr BoolLiteralExpr, env *Environment) (RTValue, error)
	VisitVarAccessExpr(expr VarAccessExpr, env *Environment) (RTValue, error)
	VisitVarAssignmentExpr(expr VarAssignmentExpr, env *Environment) (RTValue, error)
	VisitDotExpr(expr DotExpr, env *Environment) (RTValue, error)
	VisitCallExpr(expr CallExpr, env *Environment) (RTValue, error)
}

type BinaryExpr struct {
	Left  Expr
	Right Expr
	Tok   Token
	Pos   SEPos
}

func NewBinaryExpr(left Expr, right Expr, tok Token, pos SEPos) *BinaryExpr {
	return &BinaryExpr{
		Left:  left,
		Right: right,
		Tok:   tok,
		Pos:   pos,
	}
}

func (binaryExpr BinaryExpr) Accept(visitor ExprVisitor, env *Environment) (RTValue, error) {
	return visitor.VisitBinaryExpr(binaryExpr, env)
}

func (binaryExpr BinaryExpr) ToString() string {
	return fmt.Sprintf("(%s %s %s)", binaryExpr.Left.ToString(), binaryExpr.Tok.ToString(), binaryExpr.Right.ToString())
}

func (binaryExpr BinaryExpr) GetPosition() SEPos {
	return binaryExpr.Pos
}

type UnaryExpr struct {
	Tok   Token
	Right Expr
	Pos   SEPos
}

func NewUnaryExpr(tok Token, right Expr, pos SEPos) *UnaryExpr {
	return &UnaryExpr{
		Right: right,
		Tok:   tok,
		Pos:   pos,
	}
}

func (unaryExpr UnaryExpr) Accept(visitor ExprVisitor, env *Environment) (RTValue, error) {
	return visitor.VisitUnaryExpr(unaryExpr, env)
}

func (unaryExpr UnaryExpr) ToString() string {
	return fmt.Sprintf("(%s %s)", unaryExpr.Tok.ToString(), unaryExpr.Right.ToString())
}

func (unaryExpr UnaryExpr) GetPosition() SEPos {
	return unaryExpr.Pos
}

type GroupingExpr struct {
	Expression Expr
	Pos        SEPos
}

func NewGroupingExpr(expression Expr, pos SEPos) *GroupingExpr {
	return &GroupingExpr{
		Expression: expression,
		Pos:        pos,
	}
}

func (groupingExpr GroupingExpr) Accept(visitor ExprVisitor, env *Environment) (RTValue, error) {
	return visitor.VisitGroupingExpr(groupingExpr, env)
}

func (groupingExpr GroupingExpr) ToString() string {
	return fmt.Sprintf("(%s)", groupingExpr.Expression.ToString())
}

func (groupingExpr GroupingExpr) GetPosition() SEPos {
	return groupingExpr.Pos
}

type IntLiteralExpr struct {
	Value int
	Pos   SEPos
}

func NewIntLiteralExpr(value int, pos SEPos) *IntLiteralExpr {
	return &IntLiteralExpr{
		Value: value,
		Pos:   pos,
	}
}

func (intLiteralExpr IntLiteralExpr) Accept(visitor ExprVisitor, env *Environment) (RTValue, error) {
	return visitor.VisitIntLiteralExpr(intLiteralExpr, env)
}

func (intLiteralExpr IntLiteralExpr) ToString() string {
	return fmt.Sprintf("(INT: %d)", intLiteralExpr.Value)
}

func (intLiteralExpr IntLiteralExpr) GetPosition() SEPos {
	return intLiteralExpr.Pos
}

type FloatLiteralExpr struct {
	Value float64
	Pos   SEPos
}

func NewFloatLiteralExpr(value float64, pos SEPos) *FloatLiteralExpr {
	return &FloatLiteralExpr{
		Value: value,
		Pos:   pos,
	}
}

func (floatLiteralExpr FloatLiteralExpr) Accept(visitor ExprVisitor, env *Environment) (RTValue, error) {
	return visitor.VisitFloatLiteralExpr(floatLiteralExpr, env)
}

func (floatLiteralExpr FloatLiteralExpr) ToString() string {
	return fmt.Sprintf("(FLOAT: %f)", floatLiteralExpr.Value)
}

func (floatLiteralExpr FloatLiteralExpr) GetPosition() SEPos {
	return floatLiteralExpr.Pos
}

type BoolLiteralExpr struct {
	Value bool
	Pos   SEPos
}

func NewBoolLiteralExpr(value bool, pos SEPos) *BoolLiteralExpr {
	return &BoolLiteralExpr{
		Value: value,
		Pos:   pos,
	}
}

func (boolLiteralExpr BoolLiteralExpr) Accept(visitor ExprVisitor, env *Environment) (RTValue, error) {
	return visitor.VisitBoolLiteralExpr(boolLiteralExpr, env)
}

func (boolLiteralExpr BoolLiteralExpr) ToString() string {
	return fmt.Sprintf("(BOOL: %t)", boolLiteralExpr.Value)
}

func (boolLiteralExpr BoolLiteralExpr) GetPosition() SEPos {
	return boolLiteralExpr.Pos
}

type VarAccessExpr struct {
	Value string
	Pos   SEPos
}

func NewVarAccessExpr(value string, pos SEPos) *VarAccessExpr {
	return &VarAccessExpr{
		Value: value,
		Pos:   pos,
	}
}

func (varAccessExpr VarAccessExpr) Accept(visitor ExprVisitor, env *Environment) (RTValue, error) {
	return visitor.VisitVarAccessExpr(varAccessExpr, env)
}

func (varAccessExpr VarAccessExpr) ToString() string {
	return fmt.Sprintf("(VAR_ACCESS: %s)", varAccessExpr.Value)
}

func (varAccessExpr VarAccessExpr) GetPosition() SEPos {
	return varAccessExpr.Pos
}

type VarAssignmentExpr struct {
	Object Expr
	Name   string
	Value  Expr
	Pos    SEPos
}

func NewVarAssignmentExpr(object Expr, name string, value Expr, pos SEPos) *VarAssignmentExpr {
	return &VarAssignmentExpr{
		Object: object,
		Name:   name,
		Value:  value,
		Pos:    pos,
	}
}

func (varAssignmentExpr VarAssignmentExpr) Accept(visitor ExprVisitor, env *Environment) (RTValue, error) {
	return visitor.VisitVarAssignmentExpr(varAssignmentExpr, env)
}

func (varAssignmentExpr VarAssignmentExpr) ToString() string {
	return fmt.Sprintf("(VAR_ASSIGNMENT_EXPR: %s . %s = %s)", varAssignmentExpr.Object.ToString(), varAssignmentExpr.Name, varAssignmentExpr.Value.ToString())
}

func (varAssignmentExpr VarAssignmentExpr) GetPosition() SEPos {
	return varAssignmentExpr.Pos
}

type DotExpr struct {
	Left  Expr
	Right Token
	Pos   SEPos
}

func NewDotExpr(left Expr, right Token, pos SEPos) *DotExpr {
	return &DotExpr{
		Left:  left,
		Right: right,
		Pos:   pos,
	}
}

func (dotExpr DotExpr) Accept(visitor ExprVisitor, env *Environment) (RTValue, error) {
	return visitor.VisitDotExpr(dotExpr, env)
}

func (dotExpr DotExpr) ToString() string {
	return fmt.Sprintf("(DOT_EXPR: %s . %s)", dotExpr.Left.ToString(), dotExpr.Right.ToString())
}

func (dotExpr DotExpr) GetPosition() SEPos {
	return dotExpr.Pos
}

type CallExpr struct {
	Function  Expr
	Arguments []Expr
	Pos       SEPos
}

func NewCallExpr(function Expr, arguments []Expr, pos SEPos) *CallExpr {
	return &CallExpr{
		Function:  function,
		Arguments: arguments,
		Pos:       pos,
	}
}

func (callExpr CallExpr) Accept(visitor ExprVisitor, env *Environment) (RTValue, error) {
	return visitor.VisitCallExpr(callExpr, env)
}

func (callExpr CallExpr) ToString() string {
	return fmt.Sprintf("(CALL_EXPR: %s)", callExpr.Function.ToString())
}

func (callExpr CallExpr) GetPosition() SEPos {
	return callExpr.Pos
}
