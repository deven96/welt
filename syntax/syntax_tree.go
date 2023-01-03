package syntax

import "github.com/deven96/welt/diagnostic"

type SyntaxTree struct {
	diagnostics    diagnostic.DiagnosticsBag
	Root           ExpressionSyntax
	endOfFileToken SyntaxToken
}

func (st SyntaxTree) Diagnostics() diagnostic.DiagnosticsBag {
	return st.diagnostics
}

func SyntaxTreeParse(text string) SyntaxTree {
	parser := NewParser(text)
	return parser.Parse()
}
