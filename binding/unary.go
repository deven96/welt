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
	LogicalNegation
)

func (kind boundUnaryOperatorKind) String() string {
	switch kind {
	case Identity:
		return "Identity"
	case Negation:
		return "Negation"
	case LogicalNegation:
		return "LogicalNegation"
	default:
		return "Unknown"
	}
}

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
	if typ == reflect.TypeOf(a) {
		switch kind {
		case syntax.PlusToken:
			a := Identity
			return &a, nil
		case syntax.MinusToken:
			a := Negation
			return &a, nil
		default:
			return nil, errors.New(fmt.Sprintf("Unexpected unary operator %s", kind))
		}
	}
	var b bool
	if typ == reflect.TypeOf(b) {
		switch kind {
		case syntax.BangToken:
			b := LogicalNegation
			return &b, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("Unary operand %s, not of type int, bool", typ))
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
	b.diagnostics = append(b.diagnostics, fmt.Sprintf("Unary operator %s is not defined for %s: %s", input.Operator.Text, boundOperand.Type(), err))
	return boundOperand
}
