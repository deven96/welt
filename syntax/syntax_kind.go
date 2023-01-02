package syntax

type SyntaxKind int64

const (
	// Special Tokens
	NewLineToken SyntaxKind = iota
	EndOfFileToken
	WhitespaceToken
	BadToken

	// Number Token
	NumberToken

	// Arithmetric Token
	PlusToken
	MinusToken
	StarToken
	ForwardSlashToken

	// Parenthesis Tokens
	OpenParenthesisToken
	CloseParenthesisToken

	// Expressions
	LiteralExpression
	BinaryExpression
	ParenthesisedExpression
	UnaryExpression
)

func (i SyntaxKind) String() string {
	switch i {
	case NumberToken:
		return "Number"
	case WhitespaceToken:
		return "Whitespace"
	case MinusToken:
		return "Minus"
	case PlusToken:
		return "Plus"
	case StarToken:
		return "Star"
	case ForwardSlashToken:
		return "ForwardSlash"
	case OpenParenthesisToken:
		return "OpenParenthesis"
	case CloseParenthesisToken:
		return "CloseParenthesis"
	case BinaryExpression:
		return "BinaryExpression"
	case LiteralExpression:
		return "LiteralExpression"
	case UnaryExpression:
		return "UnaryExpression"
	case ParenthesisedExpression:
		return "ParenthesisedExpression"
	case NewLineToken:
		return "NewLine"
	case EndOfFileToken:
		return "EOF"
	case BadToken:
		return "Bad"
	}
	return "Unknown"
}

func (kind SyntaxKind) getBinaryOperatorPrecedence() int {
	switch kind {
	case StarToken, ForwardSlashToken:
		return 2
	case PlusToken, MinusToken:
		return 1
	default:
		return 0
	}
}

func (kind SyntaxKind) getUnaryOperatorPrecedence() int {
	switch kind {
	case PlusToken, MinusToken:
		return 3
	default:
		return 0
	}
}
