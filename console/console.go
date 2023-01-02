package console

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/deven96/welt/syntax"
	"github.com/fatih/color"
)

func clearScreen() {
	fmt.Printf("\x1bc")
}

func Console() {
	clearScreen()
	showTree := false
	for {
		fmt.Print(">>> ")
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		line := strings.TrimSuffix(text, "\n")
		if line == "#toggleTree" {
			showTree = !showTree
			if showTree {
				fmt.Println("Showing Parse trees")
			} else {
				fmt.Println("Hiding Parse trees")
			}
			continue
		} else if line == "#cls" {
			clearScreen()
			continue
		} else if line == "#exit" {
			break
		}

		syntaxTree := syntax.SyntaxTreeParse(line)
		if showTree {
			color.Set(color.FgWhite, color.Bold)
			prettyPrint(PrettyPrint{node: syntaxTree.Root, isLast: true})
			color.Unset()
		}

		diagnostics := syntaxTree.Diagnostics()

		if len(diagnostics) > 0 {
			color.Set(color.FgRed, color.Bold)
			for _, err := range diagnostics {
				fmt.Println(err)
			}
			color.Unset()
		} else {
			evaluator := NewEvaluator(syntaxTree.Root)
			result := evaluator.Evaluate()
			fmt.Println(result)
		}
	}
}
