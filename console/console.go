package console

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/deven96/welt/compiler"
	"github.com/deven96/welt/syntax"
	"github.com/deven96/welt/variables"
	"github.com/fatih/color"
)

func clearScreen() {
	fmt.Printf("\x1bc")
}

func Console() {
	clearScreen()
	showTree := false
	variables := variables.Variables{}
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
		compilation := compiler.Compilation{SyntaxTree: syntaxTree}
		compilationResult := compilation.Evaluate(&variables)
		if showTree {
			color.Set(color.FgWhite, color.Bold)
			prettyPrint(PrettyPrint{node: syntaxTree.Root, isLast: true})
			color.Unset()
		}

		if len(compilationResult.Diagnostics()) > 0 {
			for _, err := range compilationResult.Diagnostics() {
				printError(line, err)
			}
		} else {
			fmt.Println(compilationResult.Result())
		}
	}
}
