package visitors

import (
	"fmt"
	"strconv"

	"github.com/jayjunior/eval/internal/ast"
)

type Evaluator struct {
	identifiers map[string]int //TODO for now just integers
}

func (this *Evaluator) visit(exp ast.Expression) (int, error) {
	switch e := exp.(type) {
	case *ast.VarDeclaration:
		operand := e.Operand.TokenLiteral.Literal
		if _, exist := this.identifiers[operand]; exist {
			return 0, fmt.Errorf("double declaration of %s", operand)
		}
		this.identifiers[operand] = 0
		return 0, nil
	case *ast.Assignement:
		operand := e.LHS.TokenLiteral.Literal
		rhs, err := this.visit(e.Rhs)
		if err != nil {
			return 0, fmt.Errorf("error for assignement: %v", err)
		}
		this.identifiers[operand] = rhs
		return rhs, nil
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
	case *ast.Number:
		res, err := strconv.Atoi(e.TokenLiteral.Literal)
		if err != nil {
			return 0, fmt.Errorf("error converting %s to int: %v", e.TokenLiteral.Literal, err)
		}
		return res, nil
	case *ast.Identifier:
		operand := e.TokenLiteral.Literal
		if _, exist := this.identifiers[operand]; !exist {
			return 0, fmt.Errorf("undeclared identifier %s", operand)
		}
		return this.identifiers[operand], nil
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
	if this.identifiers == nil {
		this.identifiers = make(map[string]int)
	}
	return this.visit(exp)
}
