package binding

import (
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
	return expression.Operator.ResultType
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
	b.diagnostics.ReportUndefinedBinaryOperator(input.Operator.Span(), input.Operator.Text, boundLeft.Type(), boundRight.Type())
	return boundLeft
}
