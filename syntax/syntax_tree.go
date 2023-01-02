package syntax

type SyntaxTree struct {
	diagnostics    []string
	Root           ExpressionSyntax
	endOfFileToken SyntaxToken
}

func (st SyntaxTree) Diagnostics() []string {
	return st.diagnostics
}

func SyntaxTreeParse(text string) SyntaxTree {
	parser := NewParser(text)
	return parser.Parse()
}
