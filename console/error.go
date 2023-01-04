package console

import (
	"fmt"

	"github.com/deven96/welt/diagnostic"
	"github.com/fatih/color"
)

func printError(line string, err diagnostic.Diagnostic) {
	fmt.Println()
	color.Set(color.FgRed, color.Bold)
	fmt.Println(err.String())
	color.Unset()
	prefixChunk := line[0:err.Span.Start]
	errorChunk := line[err.Span.Start:err.Span.End()]
	suffixChunk := line[err.Span.End():]

	fmt.Print("    ")
	fmt.Print(prefixChunk)

	color.Set(color.FgRed, color.Bold)
	fmt.Print(errorChunk)
	color.Unset()

	fmt.Print(suffixChunk)
	fmt.Println()
}
