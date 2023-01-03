package syntax

type AssignmentExpressionSyntax struct {
	Identifier  SyntaxToken
	EqualsToken SyntaxToken
	Expression  ExpressionSyntax
}

func (es AssignmentExpressionSyntax) Kind() SyntaxKind {
	return AssignmentExpression
}

func (es AssignmentExpressionSyntax) Children() []SyntaxNode {
	return []SyntaxNode{es.Identifier, es.EqualsToken, es.Expression}
}
