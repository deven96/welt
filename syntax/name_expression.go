package syntax

type NameExpressionSyntax struct {
	Identifier SyntaxToken
}

func (es NameExpressionSyntax) Kind() SyntaxKind {
	return NameExpression
}

func (es NameExpressionSyntax) Children() []SyntaxNode {
	return []SyntaxNode{es.Identifier}
}
