package binding

import (
	"fmt"
	"reflect"

	"github.com/deven96/welt/diagnostic"
	"github.com/deven96/welt/syntax"
	"github.com/deven96/welt/variables"
)

type boundNodeKind int64

const (
	// Expressions
	UnaryExpression boundNodeKind = iota
	LiteralExpression
	BinaryExpression
	VariableExpression
	AssignmentExpression
	StringExpression
)

func (kind boundNodeKind) String() string {
	switch kind {
	case UnaryExpression:
		return "BoundUnaryExpression"
	case LiteralExpression:
		return "BoundLiteralExpression"
	case BinaryExpression:
		return "BoundBinaryExpression"
	case VariableExpression:
		return "BoundVariableExpression"
	case AssignmentExpression:
		return "BoundAssignmentExpression"
	case StringExpression:
		return "BoundStringExpression"
	default:
		return "UnknownBoundKind"
	}
}

type BoundNode interface {
	Kind() boundNodeKind
}

type BoundExpression interface {
	BoundNode
	Type() reflect.Type
}

type Binder struct {
	diagnostics diagnostic.DiagnosticsBag
	variables   *variables.Variables
}

func (b Binder) Diagnostics() diagnostic.DiagnosticsBag {
	return b.diagnostics
}

func (b *Binder) BindExpression(syntaxExpression syntax.ExpressionSyntax) BoundExpression {
	kind := syntaxExpression.Kind()
	switch kind {
	case syntax.LiteralExpression:
		return b.BindLiteral(syntaxExpression.(syntax.LiteralExpressionSyntax))
	case syntax.QuotedExpression:
		return b.BindString(syntaxExpression.(syntax.QuotedExpressionSyntax))
	case syntax.UnaryExpression:
		return b.BindUnary(syntaxExpression.(syntax.UnaryExpressionSyntax))
	case syntax.BinaryExpression:
		return b.BindBinary(syntaxExpression.(syntax.BinaryExpressionSyntax))
	case syntax.ParenthesisedExpression:
		return b.BindParenthesisedLiteral(syntaxExpression.(syntax.ParenthesisedExpressionSyntax))
	case syntax.NameExpression:
		return b.BindName(syntaxExpression.(syntax.NameExpressionSyntax))
	case syntax.AssignmentExpression:
		return b.BindAssignment(syntaxExpression.(syntax.AssignmentExpressionSyntax))
	default:
		panic(fmt.Sprintf("Unexpected syntax %s", kind))
	}
}

func NewBinder(variables *variables.Variables) Binder {
	return Binder{
		diagnostics: []diagnostic.Diagnostic{},
		variables:   variables,
	}
}
