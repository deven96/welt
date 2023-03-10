package compiler

import (
	"github.com/deven96/welt/binding"
	"github.com/deven96/welt/diagnostic"
	"github.com/deven96/welt/syntax"
	"github.com/deven96/welt/variables"
)

type Compilation struct {
	SyntaxTree syntax.SyntaxTree
}

func (comp Compilation) Evaluate(variables *variables.Variables) CompilationResult {
	binderObj := binding.NewBinder(variables)
	boundExpression := binderObj.BindExpression(comp.SyntaxTree.Root)
	diagnostics := append(comp.SyntaxTree.Diagnostics(), binderObj.Diagnostics()...)
	if len(diagnostics) >= 1 {
		return CompilationResult{diagnostics: diagnostics, result: nil}
	}
	evaluator := newEvaluator(boundExpression, variables)
	result := evaluator.evaluate()
	return CompilationResult{result: result}
}

type CompilationResult struct {
	diagnostics diagnostic.DiagnosticsBag
	result      interface{}
}

func (res CompilationResult) Diagnostics() diagnostic.DiagnosticsBag {
	return res.diagnostics
}

func (res CompilationResult) Result() interface{} {
	return res.result
}
