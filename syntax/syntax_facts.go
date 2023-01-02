package syntax

func KeyWordRecognition(input string) SyntaxKind {
	switch input {
	case "true":
		return TrueKeyWord
	case "false":
		return FalseKeyWord
	default:
		return IdentifierToken
	}
}

func (kind SyntaxKind) getUnaryOperatorPrecedence() int {
	switch kind {
	case PlusToken, MinusToken, BangToken:
		return 6
	default:
		return 0
	}
}

func (kind SyntaxKind) getBinaryOperatorPrecedence() int {
	switch kind {
	case StarToken, ForwardSlashToken:
		return 5
	case PlusToken, MinusToken:
		return 4
	case DoubleEqualToken, BangEqualToken:
		return 3
	case DoubleAmpersandToken:
		return 2
	case DoublePipeToken:
		return 1
	default:
		return 0
	}
}
