package syntax

type UnaryExpressionSyntax struct {
	Operator SyntaxToken
	Operand  ExpressionSyntax
}

func (es UnaryExpressionSyntax) Kind() SyntaxKind {
	return UnaryExpression
}

func (es UnaryExpressionSyntax) Children() []SyntaxNode {
	return []SyntaxNode{es.Operator, es.Operand}
}
