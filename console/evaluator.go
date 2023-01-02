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

func (e Evaluator) Evaluate() interface{} {
	return e.evaluateExpression(e.root)
}

func (e Evaluator) evaluateExpression(node binding.BoundExpression) interface{} {
	nroot, isLiteralExpression := node.(binding.BoundLiteralExpression)
	if isLiteralExpression {
		return nroot.Value
	}
	uroot, isUnaryExpression := node.(binding.BoundUnaryExpression)
	if isUnaryExpression {
		operand := e.evaluateExpression(uroot.Operand)

		switch uroot.Operator.Kind {
		case binding.Identity:
			return operand.(int)
		case binding.Negation:
			return -operand.(int)
		case binding.LogicalNegation:
			return !operand.(bool)
		default:
			panic(fmt.Sprintf("Unexpected unary operator %s", uroot.Operator.Kind))
		}
	}

	broot, isBinaryExpression := node.(binding.BoundBinaryExpression)
	if isBinaryExpression {
		left := e.evaluateExpression(broot.Left)
		right := e.evaluateExpression(broot.Right)

		operatorKind := broot.Operator.Kind

		switch operatorKind {
		case binding.Addition:
			return left.(int) + right.(int)
		case binding.Subtraction:
			return left.(int) - right.(int)
		case binding.Multiplication:
			return left.(int) * right.(int)
		case binding.Division:
			return left.(int) / right.(int)
		case binding.LogicalAnd:
			return left.(bool) && right.(bool)
		case binding.LogicalOr:
			return left.(bool) || right.(bool)
		case binding.LogicalEquals:
			return left == right
		case binding.LogicalNotEquals:
			return left != right
		default:
			panic(fmt.Sprintf("Unexpected binary expression %s", operatorKind))
		}
	}

	panic(fmt.Sprintf("Unexpected node %s", node.Kind()))
}
