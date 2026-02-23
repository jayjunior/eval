package internal

import (
	"fmt"
	"strconv"
)

type Evaluator struct{}

func (this *Evaluator) visit(exp Expression) (int, error) {
	switch e := exp.(type) {
	case *BinaryExpression:
		lhs, err := this.visit(e.lhs)
		if err != nil {
			return 0, err
		}
		rhs, err := this.visit(e.rhs)
		if err != nil {
			return 0, err
		}
		return this.evaluateBinaryExpression(lhs, e.operator, rhs)
	case *UnaryExpresson:
		res, err := this.visit(e.operand)
		if err != nil {
			return 0, nil
		}
		return -res, nil
	case *NumberLiteral:
		res, err := strconv.Atoi(e.literal)
		if err != nil {
			return 0, fmt.Errorf("error converting %s to string: %v\n", e.literal, err)
		}
		return res, nil
	}
	return 0, nil
}

func (this *Evaluator) evaluateBinaryExpression(lhs int, operator Token, rhs int) (int, error) {
	switch operator.token_type {
	case Plus:
		return lhs + rhs, nil
	case Minus:
		return lhs - rhs, nil
	case Multiplication:
		return lhs * rhs, nil
	case Division:
		return lhs / rhs, nil
	default:
		return 0, fmt.Errorf("unexpected token %s\n", operator.token)
	}
}

func (this *Evaluator) Evaluate(exp Expression) (int , error){
	return this.visit(exp);
}
