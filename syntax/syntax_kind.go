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

	// String Token
	StringToken

	// Arithmetric Token
	PlusToken
	MinusToken
	StarToken
	ForwardSlashToken
	ModuloToken

	// Assignment Tokens
	EqualsToken

	// Boolean Tokens
	DoubleAmpersandToken
	DoublePipeToken
	DoubleEqualToken
	BangToken
	BangEqualToken

	// Parenthesis Tokens
	OpenParenthesisToken
	CloseParenthesisToken

	// Identifier Token
	IdentifierToken
	QuotedIdentifierToken

	// Expressions
	LiteralExpression
	BinaryExpression
	ParenthesisedExpression
	QuotedExpression
	UnaryExpression
	NameExpression
	AssignmentExpression

	// Keywords
	FalseKeyWord
	TrueKeyWord
)

func (i SyntaxKind) String() string {
	switch i {
	case NumberToken:
		return "Number"
	case StringToken:
		return "String"
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
	case ModuloToken:
		return "Modulo"
	case OpenParenthesisToken:
		return "OpenParenthesis"
	case CloseParenthesisToken:
		return "CloseParenthesis"
	case IdentifierToken:
		return "Identifier"
	case QuotedIdentifierToken:
		return "QuotedIdentifier"
	case EqualsToken:
		return "Equals"
	case DoubleAmpersandToken:
		return "DoubleAmpersand"
	case DoublePipeToken:
		return "DoublePipe"
	case DoubleEqualToken:
		return "DoubleEqual"
	case BangToken:
		return "Bang"
	case NameExpression:
		return "NameExpression"
	case AssignmentExpression:
		return "AssignmentExpression"
	case BinaryExpression:
		return "BinaryExpression"
	case LiteralExpression:
		return "LiteralExpression"
	case UnaryExpression:
		return "UnaryExpression"
	case ParenthesisedExpression:
		return "ParenthesisedExpression"
	case QuotedExpression:
		return "QuotedExpression"
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
