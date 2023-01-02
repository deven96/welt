package parser

type SyntaxNode interface {
	Kind() SyntaxKind
	Children() []SyntaxNode
}

type ExpressionSyntax interface {
	SyntaxNode
}

type LiteralExpressionSyntax struct {
	LiteralToken SyntaxToken
}

func (es LiteralExpressionSyntax) Kind() SyntaxKind {
	return LiteralExpression
}

func (es LiteralExpressionSyntax) Children() []SyntaxNode {
	// SyntaxToken fulfills the SyntaxNode interface
	return []SyntaxNode{es.LiteralToken}
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
