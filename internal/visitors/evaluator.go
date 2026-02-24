package visitors

import (
	"fmt"
	"strconv"

	"github.com/jayjunior/eval/internal/ast"
)

type Evaluator struct{}

func (this *Evaluator) visit(exp ast.Expression) (int, error) {
	switch e := exp.(type) {
	case *ast.BinaryExpression:
		lhs, err := this.visit(e.Lhs)
		if err != nil {
			return 0, err
		}
		rhs, err := this.visit(e.Rhs)
		if err != nil {
			return 0, err
		}
		return this.evaluateBinaryExpression(lhs, e.Operator, rhs)
	case *ast.UnaryExpression:
		res, err := this.visit(e.Operand)
		if err != nil {
			return 0, err
		}
		return -res, nil
	case *ast.NumberLiteral:
		res, err := strconv.Atoi(e.Literal)
		if err != nil {
			return 0, fmt.Errorf("error converting %s to int: %v", e.Literal, err)
		}
		return res, nil
	}
	return 0, nil
}

func (this *Evaluator) evaluateBinaryExpression(lhs int, operator ast.Token, rhs int) (int, error) {
	switch operator.Token {
	case ast.Plus:
		return lhs + rhs, nil
	case ast.Minus:
		return lhs - rhs, nil
	case ast.Multiplication:
		return lhs * rhs, nil
	case ast.Division:
		return lhs / rhs, nil
	default:
		return 0, fmt.Errorf("unexpected token %s\n", operator.Literal)
	}
}

func (this *Evaluator) Evaluate(exp ast.Expression) (int, error) {
	return this.visit(exp)
}
