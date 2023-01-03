package binding

import (
	"reflect"

	"github.com/deven96/welt/syntax"
)

type BoundAssignmentExpression struct {
	Name  string
	Right BoundExpression
}

func (expression BoundAssignmentExpression) Kind() boundNodeKind {
	return AssignmentExpression
}

func (expression BoundAssignmentExpression) Type() reflect.Type {
	return expression.Right.Type()
}

func (b *Binder) BindAssignment(input syntax.AssignmentExpressionSyntax) BoundExpression {
	right := b.BindExpression(input.Expression)
	return BoundAssignmentExpression{
		Name:  input.Identifier.Text,
		Right: right,
	}
}
