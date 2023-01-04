package binding

import (
	"reflect"
	"strings"

	"github.com/deven96/welt/syntax"
)

// BoundStringExpression : represents the type of a quoted value.... "value" is whatever is between unescaped quotes
type BoundStringExpression struct {
	Value string
}

func (expression BoundStringExpression) Kind() boundNodeKind {
	return StringExpression
}

func (expression BoundStringExpression) Type() reflect.Type {
	return reflect.TypeOf(expression.Value)
}

func (b *Binder) BindString(input syntax.QuotedExpressionSyntax) BoundExpression {
	// Remove leading and ending quotes
	val := strings.TrimSuffix(input.QuotedIdentifier.Text, `"`)
	val = strings.TrimPrefix(val, `"`)
	return BoundStringExpression{
		Value: val,
	}
}
