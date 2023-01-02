package binding

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/deven96/welt/syntax"
)

type boundBinaryOperatorKind int64

const (
	Addition boundBinaryOperatorKind = iota
	Subtraction
	Multiplication
	Division
)

// BoundBinaryExpression : represents the type a value expression
type BoundBinaryExpression struct {
	Left         BoundExpression
	OperatorKind boundBinaryOperatorKind
	Right        BoundExpression
}

func (expression BoundBinaryExpression) Kind() boundNodeKind {
	return BinaryExpression
}

func (expression BoundBinaryExpression) Type() reflect.Type {
	return expression.Left.Type()
}

func BindBinaryOperatorKind(kind syntax.SyntaxKind, leftTyp reflect.Type, rightTyp reflect.Type) (*boundBinaryOperatorKind, error) {
	var a int
	if leftTyp != reflect.TypeOf(a) || rightTyp != reflect.TypeOf(a) {
		return nil, errors.New("BinaryOperator is not of type int")
	}

	var ret boundBinaryOperatorKind
	switch kind {
	case syntax.PlusToken:
		ret = Addition
	case syntax.MinusToken:
		ret = Subtraction
	case syntax.StarToken:
		ret = Multiplication
	case syntax.ForwardSlashToken:
		ret = Division
	default:
		panic(fmt.Sprintf("Unexpected binary operator %s", kind))
	}
	return &ret, nil
}

func (b *Binder) BindBinary(input syntax.BinaryExpressionSyntax) BoundExpression {
	boundLeft := b.BindExpression(input.Left)
	boundRight := b.BindExpression(input.Right)
	boundOperatorKind, err := BindBinaryOperatorKind(input.Operator.Kind(), boundLeft.Type(), boundRight.Type())
	if err == nil {
		return BoundBinaryExpression{
			Left:         boundLeft,
			OperatorKind: *boundOperatorKind,
			Right:        boundRight,
		}
	}
	b.diagnostics = append(b.diagnostics, fmt.Sprintf("Binary operator %s is not defined for types (%s, %s(", input.Operator.Text, boundLeft.Type(), boundRight.Type()))
	return boundLeft
}
