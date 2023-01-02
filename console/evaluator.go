package console

import (
	"fmt"

	"github.com/deven96/welt/syntax"
)

type Evaluator struct {
	root syntax.ExpressionSyntax
}

func NewEvaluator(expression syntax.ExpressionSyntax) Evaluator {
	return Evaluator{
		root: expression,
	}
}

func (e Evaluator) Evaluate() int {
	return e.evaluateExpression(e.root)
}

func (e Evaluator) evaluateExpression(node syntax.ExpressionSyntax) int {
	nroot, isLiteralExpression := node.(syntax.LiteralExpressionSyntax)
	if isLiteralExpression {
		return nroot.LiteralToken.Value.(int)
	}
	uroot, isUnaryExpression := node.(syntax.UnaryExpressionSyntax)
	if isUnaryExpression {
		operand := e.evaluateExpression(uroot.Operand)

		switch uroot.Operator.Kind() {
		case syntax.PlusToken:
			return operand
		case syntax.MinusToken:
			return -operand
		default:
			panic(fmt.Sprintf("Unexpected unary operator %s", uroot.Operator.Kind()))
		}
	}

	broot, isBinaryExpression := node.(syntax.BinaryExpressionSyntax)
	if isBinaryExpression {
		left := e.evaluateExpression(broot.Left)
		right := e.evaluateExpression(broot.Right)

		operatorKind := broot.Operator.Kind()

		switch operatorKind {
		case syntax.PlusToken:
			return left + right
		case syntax.MinusToken:
			return left - right
		case syntax.StarToken:
			return left * right
		case syntax.ForwardSlashToken:
			return left / right
		default:
			panic(fmt.Sprintf("Unexpected binary expression %s", operatorKind))
		}
	}

	proot, isParenthesisedExpression := node.(syntax.ParenthesisedExpressionSyntax)
	if isParenthesisedExpression {
		return e.evaluateExpression(proot.Expression)
	}

	panic(fmt.Sprintf("Unexpected node %s", node.Kind()))
}
