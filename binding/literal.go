package binding

import (
	"reflect"

	"github.com/deven96/welt/syntax"
)

// BoundLiteralExpression : represents the type a value expression
type BoundLiteralExpression struct {
	Value interface{}
}

func (expression BoundLiteralExpression) Kind() boundNodeKind {
	return LiteralExpression
}

func (expression BoundLiteralExpression) Type() reflect.Type {
	return reflect.TypeOf(expression.Value)
}

func (b *Binder) BindLiteral(input syntax.LiteralExpressionSyntax) BoundExpression {
	var value interface{}
	if input.Value != nil {
		value = input.Value
	} else {
		value = 0
	}
	return BoundLiteralExpression{
		Value: value,
	}
}

// BoundParenthesisedLiteralExpression: represents a literal expression in parenthesis
type BoundParenthesisedLiteralExpression struct {
	Expression BoundExpression
}

func (expression BoundParenthesisedLiteralExpression) Kind() boundNodeKind {
	return LiteralExpression
}

func (expression BoundParenthesisedLiteralExpression) Type() reflect.Type {
	return expression.Expression.Type()
}

func (b *Binder) BindParenthesisedLiteral(input syntax.ParenthesisedExpressionSyntax) BoundExpression {
	expression := b.BindExpression(input.Expression)
	return BoundParenthesisedLiteralExpression{
		Expression: expression,
	}
}
