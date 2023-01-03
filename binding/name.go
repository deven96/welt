package binding

import (
	"errors"
	"reflect"

	"github.com/deven96/welt/syntax"
	"github.com/deven96/welt/variables"
)

type BoundVariableExpression struct {
	Name string
	Typ  reflect.Type
}

func (expression BoundVariableExpression) Kind() boundNodeKind {
	return VariableExpression
}

func (expression BoundVariableExpression) Type() reflect.Type {
	return expression.Typ
}

func GetBoundName(es syntax.NameExpressionSyntax, variables *variables.Variables) (*BoundVariableExpression, error) {
	variable, ok := (*variables)[es.Identifier.Text]
	if ok {
		return &BoundVariableExpression{es.Identifier.Text, reflect.TypeOf(variable)}, nil
	} else {
		return nil, errors.New("")
	}

}

func (b *Binder) BindName(input syntax.NameExpressionSyntax) BoundExpression {
	name, err := GetBoundName(input, b.variables)
	if err == nil {
		return *name
	}
	b.diagnostics.ReportUndefinedName(input.Identifier.Span(), input.Identifier.Text)
	return BoundLiteralExpression{0}
}
