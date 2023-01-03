package syntax

type ParenthesisedExpressionSyntax struct {
	OpenParenthesisToken  SyntaxToken
	Expression            ExpressionSyntax
	CloseParenthesisToken SyntaxToken
}

func (es ParenthesisedExpressionSyntax) Kind() SyntaxKind {
	return ParenthesisedExpression
}

func (es ParenthesisedExpressionSyntax) Children() []SyntaxNode {
	return []SyntaxNode{es.OpenParenthesisToken, es.Expression, es.CloseParenthesisToken}
}
