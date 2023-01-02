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

	// Logical operators
	LogicalAnd
	LogicalOr
)

func (kind boundBinaryOperatorKind) String() string {
	switch kind {
	case Addition:
		return "Addition"
	case Subtraction:
		return "Subtraction"
	case Multiplication:
		return "Multiplication"
	case Division:
		return "Division"
	case LogicalAnd:
		return "LogicalAnd"
	case LogicalOr:
		return "LogicalOr"
	default:
		return "Unknown"
	}
}

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
	var b bool
	var ret boundBinaryOperatorKind
	if leftTyp == reflect.TypeOf(a) && rightTyp == reflect.TypeOf(a) {
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
	} else if leftTyp == reflect.TypeOf(b) && rightTyp == reflect.TypeOf(b) {
		switch kind {
		case syntax.DoubleAmpersandToken:
			ret = LogicalAnd
		case syntax.DoublePipeToken:
			ret = LogicalOr
		default:
			panic(fmt.Sprintf("Unexpected binary bool operator %s", kind))
		}
		return &ret, nil
	} else {
		return nil, errors.New("Binary operand is not of type int, bool")
	}
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
	b.diagnostics = append(b.diagnostics, fmt.Sprintf("Binary operator %s is not defined for types (%s, %s)", input.Operator.Text, boundLeft.Type(), boundRight.Type()))
	return boundLeft
}
