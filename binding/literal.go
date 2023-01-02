package binding

import (
	"reflect"

	"github.com/deven96/welt/syntax"
)

// BoundLiteralExpression : represents the type a value expression
type BoundLiteralExpression struct {
	Value interface{}
}

func (expression BoundLiteralExpression) Kind() boundNodeKind {
	return LiteralExpression
}

func (expression BoundLiteralExpression) Type() reflect.Type {
	return reflect.TypeOf(expression.Value)
}

func (b *Binder) BindLiteral(input syntax.LiteralExpressionSyntax) BoundExpression {
	value, isInt := input.LiteralToken.Value.(int)
	if !isInt {
		value = 0
	}
	return BoundLiteralExpression{
		Value: value,
	}
}
