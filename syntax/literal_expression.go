package syntax

type LiteralExpressionSyntax struct {
	LiteralToken SyntaxToken
	Value        interface{}
}

func (es LiteralExpressionSyntax) Kind() SyntaxKind {
	return LiteralExpression
}

func (es LiteralExpressionSyntax) Children() []SyntaxNode {
	// SyntaxToken fulfills the SyntaxNode interface
	return []SyntaxNode{es.LiteralToken}
}
