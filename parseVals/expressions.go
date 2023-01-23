package parsevals

import (
	"fmt"

	"github.com/snowlanguage/go-snow/position"
	"github.com/snowlanguage/go-snow/runtimevalues"
	"github.com/snowlanguage/go-snow/token"
)

type Expr interface {
	Accept(visitor ExprVisitor) (runtimevalues.RTValue, error)
	ToString() string
}

type ExprVisitor interface {
	VisitBinaryExpr(expr BinaryExpr) (runtimevalues.RTValue, error)
	VisitIntLiteralExpr(expr IntLiteralExpr) (runtimevalues.RTValue, error)
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

func (binaryExpr BinaryExpr) Accept(visitor ExprVisitor) (runtimevalues.RTValue, error) {
	return visitor.VisitBinaryExpr(binaryExpr)
}

func (binaryExpr BinaryExpr) ToString() string {
	return fmt.Sprintf("(%s %s %s)", binaryExpr.Left.ToString(), binaryExpr.Tok.ToString(), binaryExpr.Right.ToString())
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

func (intLiteralExpr IntLiteralExpr) Accept(visitor ExprVisitor) (runtimevalues.RTValue, error) {
	return visitor.VisitIntLiteralExpr(intLiteralExpr)
}

func (intLiteralExpr IntLiteralExpr) ToString() string {
	return fmt.Sprintf("(INT: %d)", intLiteralExpr.Value)
}
