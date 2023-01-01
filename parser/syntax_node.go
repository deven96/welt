package parser

type SyntaxNode interface {
	Kind() SyntaxKind
	Children() []SyntaxNode
}

type ExpressionSyntax interface {
	SyntaxNode
}

type NumberExpressionSyntax struct {
	Token SyntaxToken
}

func (es NumberExpressionSyntax) Kind() SyntaxKind {
	return NumberExpression
}

func (es NumberExpressionSyntax) Children() []SyntaxNode {
	// SyntaxToken fulfills the SyntaxNode interface
	return []SyntaxNode{es.Token}
}

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
