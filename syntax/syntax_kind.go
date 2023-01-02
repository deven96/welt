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

	// Identifier Token
	IdentifierToken

	// Boolean Tokens
	DoubleAmpersandToken
	DoublePipeToken
	DoubleEqualToken
	BangToken
	BangEqualToken

	// Parenthesis Tokens
	OpenParenthesisToken
	CloseParenthesisToken

	// Expressions
	LiteralExpression
	BinaryExpression
	ParenthesisedExpression
	UnaryExpression

	// Keywords
	FalseKeyWord
	TrueKeyWord
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
	case IdentifierToken:
		return "Identifier"
	case DoubleAmpersandToken:
		return "DoubleAmpersand"
	case DoublePipeToken:
		return "DoublePipe"
	case DoubleEqualToken:
		return "DoubleEqual"
	case BangToken:
		return "Bang"
	case BinaryExpression:
		return "BinaryExpression"
	case LiteralExpression:
		return "LiteralExpression"
	case UnaryExpression:
		return "UnaryExpression"
	case ParenthesisedExpression:
		return "ParenthesisedExpression"
	case TrueKeyWord:
		return "TrueKeyword"
	case FalseKeyWord:
		return "FalseKeyword"
	case NewLineToken:
		return "NewLine"
	case EndOfFileToken:
		return "EOF"
	case BadToken:
		return "Bad"
	}
	return "Unknown"
}
