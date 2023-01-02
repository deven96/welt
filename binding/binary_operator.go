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

type boundBinaryOperator struct {
	SyntaxKind       syntax.SyntaxKind
	Kind             boundBinaryOperatorKind
	LeftOperandType  reflect.Type
	RightOperandType reflect.Type
	ResultType       reflect.Type
}

func binaryOperations() []boundBinaryOperator {
	var b bool
	var a int
	operators := []boundBinaryOperator{
		{syntax.PlusToken, Addition, reflect.TypeOf(a), reflect.TypeOf(a), reflect.TypeOf(a)},
		{syntax.MinusToken, Addition, reflect.TypeOf(a), reflect.TypeOf(a), reflect.TypeOf(a)},
		{syntax.StarToken, Multiplication, reflect.TypeOf(a), reflect.TypeOf(a), reflect.TypeOf(a)},
		{syntax.ForwardSlashToken, Addition, reflect.TypeOf(a), reflect.TypeOf(a), reflect.TypeOf(a)},
		{syntax.DoubleAmpersandToken, LogicalAnd, reflect.TypeOf(b), reflect.TypeOf(b), reflect.TypeOf(b)},
		{syntax.DoublePipeToken, LogicalOr, reflect.TypeOf(b), reflect.TypeOf(b), reflect.TypeOf(b)},
	}
	return operators
}

func GetBoundBinaryOperator(kind syntax.SyntaxKind, leftType, rightType reflect.Type) (*boundBinaryOperator, error) {
	operators := binaryOperations()
	for _, operator := range operators {
		if operator.SyntaxKind == kind && operator.LeftOperandType == leftType && operator.RightOperandType == rightType {
			return &operator, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("Binary operator %s not implemented over %s & %s", kind, leftType, rightType))
}
