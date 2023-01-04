package binding

import (
	"errors"
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

type boundUnaryOperator struct {
	SyntaxKind  syntax.SyntaxKind
	Kind        boundUnaryOperatorKind
	OperandType reflect.Type
}

func unaryOperations() []boundUnaryOperator {
	var intType int
	var boolType bool
	operators := []boundUnaryOperator{
		{syntax.BangToken, LogicalNegation, reflect.TypeOf(boolType)},
		{syntax.PlusToken, Identity, reflect.TypeOf(intType)},
		{syntax.MinusToken, Negation, reflect.TypeOf(intType)},
	}
	return operators
}

func GetBoundUnaryOperator(kind syntax.SyntaxKind, operandType reflect.Type) (*boundUnaryOperator, error) {
	operators := unaryOperations()
	for _, operator := range operators {
		if operator.SyntaxKind == kind && operator.OperandType == operandType {
			return &operator, nil
		}
	}
	return nil, errors.New("")
}
