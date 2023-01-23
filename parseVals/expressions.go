package parsevals

import (
	"fmt"

	"github.com/snowlanguage/go-snow/position"
	"github.com/snowlanguage/go-snow/runtimevalues"
)

type Expr interface {
	Accept(visitor ExprVisitor) (runtimevalues.RTValue, error)
	ToString() string
}

type ExprVisitor interface {
	VisitIntLiteralExpr(expr IntLiteralExpr) (runtimevalues.RTValue, error)
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
