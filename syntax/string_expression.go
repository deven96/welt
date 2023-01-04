package syntax

type QuotedExpressionSyntax struct {
	QuotedIdentifier SyntaxToken
}

func (es QuotedExpressionSyntax) Kind() SyntaxKind {
	return QuotedExpression
}

func (es QuotedExpressionSyntax) Children() []SyntaxNode {
	return []SyntaxNode{es.QuotedIdentifier}
}
