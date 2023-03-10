package compiler

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/deven96/welt/binding"
	"github.com/deven96/welt/variables"
)

type Evaluator struct {
	root      binding.BoundExpression
	variables *variables.Variables
}

func newEvaluator(expression binding.BoundExpression, variables *variables.Variables) Evaluator {
	return Evaluator{
		root:      expression,
		variables: variables,
	}
}

func (e Evaluator) evaluate() interface{} {
	return e.evaluateExpression(e.root)
}

func (e Evaluator) evaluateExpression(node binding.BoundExpression) interface{} {
	nroot, isLiteralExpression := node.(binding.BoundLiteralExpression)
	if isLiteralExpression {
		return nroot.Value
	}
	naroot, isNameExpression := node.(binding.BoundVariableExpression)
	if isNameExpression {
		val := (*e.variables)[naroot.Name]
		return val
	}
	aroot, isAssignmentExpression := node.(binding.BoundAssignmentExpression)
	if isAssignmentExpression {
		val := e.evaluateExpression(aroot.Right)
		(*e.variables)[aroot.Name] = val
		return val
	}
	proot, isParenthesisExpression := node.(binding.BoundParenthesisedLiteralExpression)
	if isParenthesisExpression {
		return e.evaluateExpression(proot.Expression)
	}
	sroot, isStringExpression := node.(binding.BoundStringExpression)
	if isStringExpression {
		return sroot.Value
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

		var (
			intType    int
			stringType string
		)

		switch operatorKind {
		case binding.Addition:
			if broot.Left.Type() == reflect.TypeOf(intType) && broot.Right.Type() == reflect.TypeOf(intType) {
				return left.(int) + right.(int)
			}
			if broot.Left.Type() == reflect.TypeOf(stringType) && broot.Right.Type() == reflect.TypeOf(stringType) {
				return left.(string) + right.(string)
			}
		case binding.Subtraction:
			return left.(int) - right.(int)
		case binding.Multiplication:
			if broot.Left.Type() == reflect.TypeOf(intType) && broot.Right.Type() == reflect.TypeOf(intType) {
				return left.(int) * right.(int)
			}
			if broot.Left.Type() == reflect.TypeOf(stringType) && broot.Right.Type() == reflect.TypeOf(intType) {
				return strings.Repeat(left.(string), right.(int))
			}
		case binding.Division:
			return left.(int) / right.(int)
		case binding.Modulus:
			return left.(int) % right.(int)
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

	panic(fmt.Sprintf("Unexpected node %s %v %v", node.Kind(), node, (*e.variables)))
}
