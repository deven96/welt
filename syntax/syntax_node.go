package syntax

type SyntaxNode interface {
	Kind() SyntaxKind
	Children() []SyntaxNode
}
