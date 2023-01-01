package console

import (
	"fmt"

	"github.com/deven96/welt/parser"
)

type PrettyPrint struct {
	node   parser.SyntaxNode
	indent string
	isLast bool
}

func prettyPrint(printer PrettyPrint) {
	var (
		marker string
	)
	if printer.isLast {
		marker = `└──`
	} else {
		marker = `├──`
	}

	fmt.Print(printer.indent)
	fmt.Print(marker)
	fmt.Print(printer.node.Kind())
	syntaxToken, isSyntaxToken := printer.node.(parser.SyntaxToken)
	if isSyntaxToken && syntaxToken.Value != nil {
		fmt.Print(" ")
		fmt.Print(syntaxToken.Value)
	}
	fmt.Println()

	if printer.isLast {
		printer.indent += `    `
	} else {
		printer.indent += `|   `
	}

	children := printer.node.Children()
	var lastChild parser.SyntaxNode
	if len(children) >= 1 {
		lastChild = children[len(children)-1]
	} else {
		lastChild = nil
	}

	for _, child := range children {
		prettyPrint(PrettyPrint{
			node:   child,
			indent: printer.indent,
			isLast: child == lastChild,
		})
	}
}
