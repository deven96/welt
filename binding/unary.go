package binding

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/deven96/welt/syntax"
)

type boundUnaryOperatorKind int64

const (
	Identity boundUnaryOperatorKind = iota
	Negation
)

// BoundUnaryExpression : represents the type a unary expression
type BoundUnaryExpression struct {
	OperatorKind boundUnaryOperatorKind
	Operand      BoundExpression
}

func (expression BoundUnaryExpression) Kind() boundNodeKind {
	return UnaryExpression
}

func (expression BoundUnaryExpression) Type() reflect.Type {
	return expression.Operand.Type()
}

func BindUnaryOperatorKind(kind syntax.SyntaxKind, typ reflect.Type) (*boundUnaryOperatorKind, error) {
	var a int
	if typ != reflect.TypeOf(a) {
		return nil, errors.New("Unary operator is not of type int")
	}
	switch kind {
	case syntax.PlusToken:
		a := Identity
		return &a, nil
	case syntax.MinusToken:
		a := Negation
		return &a, nil
	default:
		panic(fmt.Sprintf("Unexpected unary operator %s", kind))
	}
}

func (b *Binder) BindUnary(input syntax.UnaryExpressionSyntax) BoundExpression {
	boundOperand := b.BindExpression(input.Operand)
	boundOperatorKind, err := BindUnaryOperatorKind(input.Operator.Kind(), boundOperand.Type())
	if err == nil {
		return BoundUnaryExpression{
			OperatorKind: *boundOperatorKind,
			Operand:      boundOperand,
		}
	}
	b.diagnostics = append(b.diagnostics, fmt.Sprintf("Unary operator %s is not defined for %s", input.Operator.Text, boundOperand.Type()))
	return boundOperand
}
