package console

import (
	"fmt"

	"github.com/deven96/welt/binding"
)

type Evaluator struct {
	root binding.BoundExpression
}

func NewEvaluator(expression binding.BoundExpression) Evaluator {
	return Evaluator{
		root: expression,
	}
}

func (e Evaluator) Evaluate() int {
	return e.evaluateExpression(e.root)
}

func (e Evaluator) evaluateExpression(node binding.BoundExpression) int {
	nroot, isLiteralExpression := node.(binding.BoundLiteralExpression)
	if isLiteralExpression {
		return nroot.Value.(int)
	}
	uroot, isUnaryExpression := node.(binding.BoundUnaryExpression)
	if isUnaryExpression {
		operand := e.evaluateExpression(uroot.Operand)

		switch uroot.OperatorKind {
		case binding.Identity:
			return operand
		case binding.Negation:
			return -operand
		default:
			panic(fmt.Sprintf("Unexpected unary operator %s", uroot.OperatorKind))
		}
	}

	broot, isBinaryExpression := node.(binding.BoundBinaryExpression)
	if isBinaryExpression {
		left := e.evaluateExpression(broot.Left)
		right := e.evaluateExpression(broot.Right)

		operatorKind := broot.OperatorKind

		switch operatorKind {
		case binding.Addition:
			return left + right
		case binding.Subtraction:
			return left - right
		case binding.Multiplication:
			return left * right
		case binding.Division:
			return left / right
		default:
			panic(fmt.Sprintf("Unexpected binary expression %s", operatorKind))
		}
	}

	panic(fmt.Sprintf("Unexpected node %s", node.Kind()))
}
