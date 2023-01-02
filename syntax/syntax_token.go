package syntax

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
