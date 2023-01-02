package binding

import (
	"fmt"
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
	b.diagnostics = append(b.diagnostics, fmt.Sprintf("Unary operator %s is not defined for %s: %s", input.Operator.Text, boundOperand.Type(), err))
	return boundOperand
}
