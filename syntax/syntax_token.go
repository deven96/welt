package syntax

import "github.com/deven96/welt/diagnostic"

type SyntaxToken struct {
	Kind_    SyntaxKind
	Text     string
	Value    interface{}
	position int
}

func (st SyntaxToken) Kind() SyntaxKind {
	return st.Kind_
}

func (st SyntaxToken) Children() []SyntaxNode {
	return []SyntaxNode{}
}

func (st SyntaxToken) Span() diagnostic.TextSpan {
	return diagnostic.TextSpan{Start: st.position, Length: len(st.Text)}
}
