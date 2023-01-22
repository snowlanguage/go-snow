package parsevals

import (
	"fmt"

	"github.com/snowlanguage/go-snow/position"
)

type Expr interface {
	Accept(visitor ExprVisitor) interface{}
	ToString() string
}

type ExprVisitor interface {
	visitIntLiteralExpr(expr *IntLiteralExpr) interface{}
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

func (intLiteralExpr *IntLiteralExpr) Accept(visitor ExprVisitor) interface{} {
	return visitor.visitIntLiteralExpr(intLiteralExpr)
}

func (intLiteralExpr *IntLiteralExpr) ToString() string {
	return fmt.Sprintf("(INT: %d)", intLiteralExpr.Value)
}
