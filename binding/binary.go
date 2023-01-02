package binding

import (
	"fmt"
	"reflect"

	"github.com/deven96/welt/syntax"
)

// BoundBinaryExpression : represents the type a value expression
type BoundBinaryExpression struct {
	Left     BoundExpression
	Operator boundBinaryOperator
	Right    BoundExpression
}

func (expression BoundBinaryExpression) Kind() boundNodeKind {
	return BinaryExpression
}

func (expression BoundBinaryExpression) Type() reflect.Type {
	return expression.Left.Type()
}

func (b *Binder) BindBinary(input syntax.BinaryExpressionSyntax) BoundExpression {
	boundLeft := b.BindExpression(input.Left)
	boundRight := b.BindExpression(input.Right)
	boundOperator, err := GetBoundBinaryOperator(input.Operator.Kind(), boundLeft.Type(), boundRight.Type())
	if err == nil {
		return BoundBinaryExpression{
			Left:     boundLeft,
			Operator: *boundOperator,
			Right:    boundRight,
		}
	}
	b.diagnostics = append(b.diagnostics, fmt.Sprintf("Binary operator %s is not defined for types (%s, %s)", input.Operator.Text, boundLeft.Type(), boundRight.Type()))
	return boundLeft
}
