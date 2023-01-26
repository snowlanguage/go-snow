package parsevals

import (
	"fmt"

	"github.com/snowlanguage/go-snow/position"
	"github.com/snowlanguage/go-snow/runtimevalues"
	"github.com/snowlanguage/go-snow/token"
)

type Expr interface {
	Accept(visitor ExprVisitor, env *runtimevalues.Environment) (runtimevalues.RTValue, error)
	ToString() string
	GetPosition() position.SEPos
}

type ExprVisitor interface {
	VisitBinaryExpr(expr BinaryExpr, env *runtimevalues.Environment) (runtimevalues.RTValue, error)
	VisitUnaryExpr(expr UnaryExpr, env *runtimevalues.Environment) (runtimevalues.RTValue, error)
	VisitGroupingExpr(expr GroupingExpr, env *runtimevalues.Environment) (runtimevalues.RTValue, error)
	VisitIntLiteralExpr(expr IntLiteralExpr, env *runtimevalues.Environment) (runtimevalues.RTValue, error)
	VisitFloatLiteralExpr(expr FloatLiteralExpr, env *runtimevalues.Environment) (runtimevalues.RTValue, error)
	VisitBoolLiteralExpr(expr BoolLiteralExpr, env *runtimevalues.Environment) (runtimevalues.RTValue, error)
	VisitVarAccessExpr(expr VarAccessExpr, env *runtimevalues.Environment) (runtimevalues.RTValue, error)
}

type BinaryExpr struct {
	Left  Expr
	Right Expr
	Tok   token.Token
	Pos   position.SEPos
}

func NewBinaryExpr(left Expr, right Expr, tok token.Token, pos position.SEPos) *BinaryExpr {
	return &BinaryExpr{
		Left:  left,
		Right: right,
		Tok:   tok,
		Pos:   pos,
	}
}

func (binaryExpr BinaryExpr) Accept(visitor ExprVisitor, env *runtimevalues.Environment) (runtimevalues.RTValue, error) {
	return visitor.VisitBinaryExpr(binaryExpr, env)
}

func (binaryExpr BinaryExpr) ToString() string {
	return fmt.Sprintf("(%s %s %s)", binaryExpr.Left.ToString(), binaryExpr.Tok.ToString(), binaryExpr.Right.ToString())
}

func (binaryExpr BinaryExpr) GetPosition() position.SEPos {
	return binaryExpr.Pos
}

type UnaryExpr struct {
	Tok   token.Token
	Right Expr
	Pos   position.SEPos
}

func NewUnaryExpr(tok token.Token, right Expr, pos position.SEPos) *UnaryExpr {
	return &UnaryExpr{
		Right: right,
		Tok:   tok,
		Pos:   pos,
	}
}

func (unaryExpr UnaryExpr) Accept(visitor ExprVisitor, env *runtimevalues.Environment) (runtimevalues.RTValue, error) {
	return visitor.VisitUnaryExpr(unaryExpr, env)
}

func (unaryExpr UnaryExpr) ToString() string {
	return fmt.Sprintf("(%s %s)", unaryExpr.Tok.ToString(), unaryExpr.Right.ToString())
}

func (unaryExpr UnaryExpr) GetPosition() position.SEPos {
	return unaryExpr.Pos
}

type GroupingExpr struct {
	Expression Expr
	Pos        position.SEPos
}

func NewGroupingExpr(expression Expr, pos position.SEPos) *GroupingExpr {
	return &GroupingExpr{
		Expression: expression,
		Pos:        pos,
	}
}

func (groupingExpr GroupingExpr) Accept(visitor ExprVisitor, env *runtimevalues.Environment) (runtimevalues.RTValue, error) {
	return visitor.VisitGroupingExpr(groupingExpr, env)
}

func (groupingExpr GroupingExpr) ToString() string {
	return fmt.Sprintf("(%s)", groupingExpr.Expression.ToString())
}

func (groupingExpr GroupingExpr) GetPosition() position.SEPos {
	return groupingExpr.Pos
}

type IntLiteralExpr struct {
	Value int
	Pos   position.SEPos
}

func NewIntLiteralExpr(value int, pos position.SEPos) *IntLiteralExpr {
	return &IntLiteralExpr{
		Value: value,
		Pos:   pos,
	}
}

func (intLiteralExpr IntLiteralExpr) Accept(visitor ExprVisitor, env *runtimevalues.Environment) (runtimevalues.RTValue, error) {
	return visitor.VisitIntLiteralExpr(intLiteralExpr, env)
}

func (intLiteralExpr IntLiteralExpr) ToString() string {
	return fmt.Sprintf("(INT: %d)", intLiteralExpr.Value)
}

func (intLiteralExpr IntLiteralExpr) GetPosition() position.SEPos {
	return intLiteralExpr.Pos
}

type FloatLiteralExpr struct {
	Value float64
	Pos   position.SEPos
}

func NewFloatLiteralExpr(value float64, pos position.SEPos) *FloatLiteralExpr {
	return &FloatLiteralExpr{
		Value: value,
		Pos:   pos,
	}
}

func (floatLiteralExpr FloatLiteralExpr) Accept(visitor ExprVisitor, env *runtimevalues.Environment) (runtimevalues.RTValue, error) {
	return visitor.VisitFloatLiteralExpr(floatLiteralExpr, env)
}

func (floatLiteralExpr FloatLiteralExpr) ToString() string {
	return fmt.Sprintf("(FLOAT: %f)", floatLiteralExpr.Value)
}

func (floatLiteralExpr FloatLiteralExpr) GetPosition() position.SEPos {
	return floatLiteralExpr.Pos
}

type BoolLiteralExpr struct {
	Value bool
	Pos   position.SEPos
}

func NewBoolLiteralExpr(value bool, pos position.SEPos) *BoolLiteralExpr {
	return &BoolLiteralExpr{
		Value: value,
		Pos:   pos,
	}
}

func (boolLiteralExpr BoolLiteralExpr) Accept(visitor ExprVisitor, env *runtimevalues.Environment) (runtimevalues.RTValue, error) {
	return visitor.VisitBoolLiteralExpr(boolLiteralExpr, env)
}

func (boolLiteralExpr BoolLiteralExpr) ToString() string {
	return fmt.Sprintf("(BOOL: %t)", boolLiteralExpr.Value)
}

func (boolLiteralExpr BoolLiteralExpr) GetPosition() position.SEPos {
	return boolLiteralExpr.Pos
}

// TODO: ADD VAR ACCESS EXPR

type VarAccessExpr struct {
	Value string
	Pos   position.SEPos
}

func NewVarAccessExpr(value string, pos position.SEPos) *VarAccessExpr {
	return &VarAccessExpr{
		Value: value,
		Pos:   pos,
	}
}

func (varAccessExpr VarAccessExpr) Accept(visitor ExprVisitor, env *runtimevalues.Environment) (runtimevalues.RTValue, error) {
	return visitor.VisitVarAccessExpr(varAccessExpr, env)
}

func (varAccessExpr VarAccessExpr) ToString() string {
	return fmt.Sprintf("(VAR_ACCESS: %s)", varAccessExpr.Value)
}

func (varAccessExpr VarAccessExpr) GetPosition() position.SEPos {
	return varAccessExpr.Pos
}
