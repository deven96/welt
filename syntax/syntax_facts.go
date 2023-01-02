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
