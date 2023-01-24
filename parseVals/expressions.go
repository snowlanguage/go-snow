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
}

type ExprVisitor interface {
	VisitBinaryExpr(expr BinaryExpr, env *runtimevalues.Environment) (runtimevalues.RTValue, error)
	VisitUnaryExpr(expr UnaryExpr, env *runtimevalues.Environment) (runtimevalues.RTValue, error)
	VisitIntLiteralExpr(expr IntLiteralExpr, env *runtimevalues.Environment) (runtimevalues.RTValue, error)
	VisitFloatLiteralExpr(expr FloatLiteralExpr, env *runtimevalues.Environment) (runtimevalues.RTValue, error)
	VisitBoolLiteralExpr(expr BoolLiteralExpr, env *runtimevalues.Environment) (runtimevalues.RTValue, error)
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
