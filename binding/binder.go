package binding

import (
	"fmt"
	"reflect"

	"github.com/deven96/welt/syntax"
)

type boundNodeKind int64

const (
	// Expressions
	UnaryExpression boundNodeKind = iota
	LiteralExpression
	BinaryExpression
)

func (kind boundNodeKind) String() string {
	switch kind {
	case UnaryExpression:
		return "BoundUnaryExpression"
	case LiteralExpression:
		return "BoundLiteralExpression"
	case BinaryExpression:
		return "BoundBinaryExpression"
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
	diagnostics []string
}

func (b Binder) Diagnostics() []string {
	return b.diagnostics
}

func (b *Binder) BindExpression(syntaxExpression syntax.ExpressionSyntax) BoundExpression {
	kind := syntaxExpression.Kind()
	switch kind {
	case syntax.LiteralExpression:
		return b.BindLiteral(syntaxExpression.(syntax.LiteralExpressionSyntax))
	case syntax.UnaryExpression:
		return b.BindUnary(syntaxExpression.(syntax.UnaryExpressionSyntax))
	case syntax.BinaryExpression:
		return b.BindBinary(syntaxExpression.(syntax.BinaryExpressionSyntax))
	case syntax.ParenthesisedExpression:
		return b.BindParenthesisedLiteral(syntaxExpression.(syntax.ParenthesisedExpressionSyntax))
	default:
		panic(fmt.Sprintf("Unexpected syntax %s", kind))
	}
}
