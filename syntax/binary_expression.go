package syntax

type BinaryExpressionSyntax struct {
	Left     ExpressionSyntax
	Operator SyntaxToken
	Right    ExpressionSyntax
}

func (es BinaryExpressionSyntax) Kind() SyntaxKind {
	return BinaryExpression
}

func (es BinaryExpressionSyntax) Children() []SyntaxNode {
	return []SyntaxNode{es.Left, es.Operator, es.Right}
}
