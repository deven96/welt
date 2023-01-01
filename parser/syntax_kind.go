package parser

type SyntaxKind int64

const (
	NumberToken SyntaxKind = iota
	WhitespaceToken
	PlusToken
	MinusToken
	StarToken
	ForwardSlashToken
	OpenParenthesisToken
	CloseParenthesisToken
	NewLineToken
	EndOfFileToken
	BadToken
	NumberExpression
	BinaryExpression
	ParenthesisedExpression
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
	case NumberExpression:
		return "NumberExpression"
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

func (kind SyntaxKind) isBinaryTermOperator() bool {
	if kind == PlusToken || kind == MinusToken {
		return true
	}
	return false
}

func (kind SyntaxKind) isBinaryFactorOperator() bool {
	if kind == StarToken || kind == ForwardSlashToken {
		return true
	}
	return false
}
