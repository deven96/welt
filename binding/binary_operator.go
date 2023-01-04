package binding

import (
	"errors"
	"reflect"

	"github.com/deven96/welt/syntax"
)

type boundBinaryOperatorKind int64

const (
	Addition boundBinaryOperatorKind = iota
	Subtraction
	Multiplication
	Division
	Modulus

	// Logical operators
	LogicalAnd
	LogicalOr
	LogicalEquals
	LogicalNotEquals
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
	case Modulus:
		return "Modulus"
	case LogicalAnd:
		return "LogicalAnd"
	case LogicalOr:
		return "LogicalOr"
	case LogicalEquals:
		return "LogicalEquals"
	case LogicalNotEquals:
		return "LogicalNotEquals"
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
	var intType int
	var boolType bool
	var stringType string
	operators := []boundBinaryOperator{
		{syntax.PlusToken, Addition, reflect.TypeOf(intType), reflect.TypeOf(intType), reflect.TypeOf(intType)},
		{syntax.PlusToken, Addition, reflect.TypeOf(stringType), reflect.TypeOf(stringType), reflect.TypeOf(stringType)},
		{syntax.MinusToken, Subtraction, reflect.TypeOf(intType), reflect.TypeOf(intType), reflect.TypeOf(intType)},
		{syntax.StarToken, Multiplication, reflect.TypeOf(intType), reflect.TypeOf(intType), reflect.TypeOf(intType)},
		{syntax.StarToken, Multiplication, reflect.TypeOf(stringType), reflect.TypeOf(intType), reflect.TypeOf(stringType)},
		{syntax.ForwardSlashToken, Division, reflect.TypeOf(intType), reflect.TypeOf(intType), reflect.TypeOf(intType)},
		{syntax.ModuloToken, Modulus, reflect.TypeOf(intType), reflect.TypeOf(intType), reflect.TypeOf(intType)},
		{syntax.DoubleEqualToken, LogicalEquals, reflect.TypeOf(intType), reflect.TypeOf(intType), reflect.TypeOf(boolType)},
		{syntax.DoubleEqualToken, LogicalEquals, reflect.TypeOf(stringType), reflect.TypeOf(stringType), reflect.TypeOf(boolType)},
		{syntax.DoubleEqualToken, LogicalEquals, reflect.TypeOf(boolType), reflect.TypeOf(boolType), reflect.TypeOf(boolType)},
		{syntax.BangEqualToken, LogicalNotEquals, reflect.TypeOf(intType), reflect.TypeOf(intType), reflect.TypeOf(boolType)},
		{syntax.BangEqualToken, LogicalNotEquals, reflect.TypeOf(stringType), reflect.TypeOf(stringType), reflect.TypeOf(boolType)},
		{syntax.BangEqualToken, LogicalNotEquals, reflect.TypeOf(boolType), reflect.TypeOf(boolType), reflect.TypeOf(boolType)},
		{syntax.DoubleAmpersandToken, LogicalAnd, reflect.TypeOf(boolType), reflect.TypeOf(boolType), reflect.TypeOf(boolType)},
		{syntax.DoublePipeToken, LogicalOr, reflect.TypeOf(boolType), reflect.TypeOf(boolType), reflect.TypeOf(boolType)},
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
	return nil, errors.New("")
}
