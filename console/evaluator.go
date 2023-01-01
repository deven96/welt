package console

import (
	"fmt"

	"github.com/deven96/welt/parser"
)

type Evaluator struct {
	root parser.ExpressionSyntax
}

func NewEvaluator(expression parser.ExpressionSyntax) Evaluator {
	return Evaluator{
		root: expression,
	}
}

func (e Evaluator) Evaluate() int {
	return e.evaluateExpression(e.root)
}

func (e Evaluator) evaluateExpression(node parser.ExpressionSyntax) int {
	nroot, isNumberExpression := node.(parser.NumberExpressionSyntax)
	if isNumberExpression {
		return nroot.Token.Value.(int)
	}
	broot, isBinaryExpression := node.(parser.BinaryExpressionSyntax)
	if isBinaryExpression {
		left := e.evaluateExpression(broot.Left)
		right := e.evaluateExpression(broot.Right)

		operatorKind := broot.Operator.Kind()

		if operatorKind == parser.PlusToken {
			return left + right
		} else if operatorKind == parser.MinusToken {
			return left - right
		} else if operatorKind == parser.StarToken {
			return left * right
		} else if operatorKind == parser.ForwardSlashToken {
			return left / right
		} else {
			panic(fmt.Sprintf("Unexpected binary expression %s", operatorKind))
		}
	}

	proot, isParenthesisedExpression := node.(parser.ParenthesisedExpressionSyntax)
	if isParenthesisedExpression {
		return e.evaluateExpression(proot.Expression)
	}

	panic(fmt.Sprintf("Unexpected node %s", node.Kind()))
}
