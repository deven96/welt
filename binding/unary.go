package binding

import (
	"reflect"

	"github.com/deven96/welt/syntax"
)

// BoundUnaryExpression : represents the type a unary expression
type BoundUnaryExpression struct {
	Operator boundUnaryOperator
	Operand  BoundExpression
}

func (expression BoundUnaryExpression) Kind() boundNodeKind {
	return UnaryExpression
}

func (expression BoundUnaryExpression) Type() reflect.Type {
	return expression.Operator.OperandType
}

func (b *Binder) BindUnary(input syntax.UnaryExpressionSyntax) BoundExpression {
	boundOperand := b.BindExpression(input.Operand)
	boundOperator, err := GetBoundUnaryOperator(input.Operator.Kind(), boundOperand.Type())
	if err == nil {
		return BoundUnaryExpression{
			Operator: *boundOperator,
			Operand:  boundOperand,
		}
	}
	b.diagnostics.ReportUndefinedUnaryOperator(input.Operator.Span(), input.Operator.Text, boundOperand.Type())
	return boundOperand
}
