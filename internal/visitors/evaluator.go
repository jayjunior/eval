package visitors

import (
	"fmt"
	"strconv"

	"github.com/jayjunior/eval/internal/ast"
)

type Evaluator struct {
	identifiers map[string]string
}

func (this *Evaluator) visit(exp ast.Expression) (string, error) {
	switch e := exp.(type) {
	case *ast.VarDeclaration:
		operand := e.Operand.TokenLiteral.Literal
		if _, exist := this.identifiers[operand]; exist {
			return "", fmt.Errorf("double declaration of %s", operand)
		}
		this.identifiers[operand] = "0"
		return "0", nil
	case *ast.Assignement:
		operand := e.LHS.TokenLiteral.Literal
		rhs, err := this.visit(e.Rhs)
		if err != nil {
			return "", fmt.Errorf("error for assignement: %v", err)
		}
		this.identifiers[operand] = rhs
		return rhs, nil
	case *ast.BinaryExpression:
		lhs, err := this.visit(e.Lhs)
		if err != nil {
			return "", err
		}
		rhs, err := this.visit(e.Rhs)
		if err != nil {
			return "", err
		}
		res, err := this.evaluateBinaryExpression(lhs, e.Operator, rhs)
		if err != nil {
			return "", err
		}
		return res, nil
	case *ast.UnaryExpression:
		res, err := this.visit(e.Operand)
		if err != nil {
			return "", err
		}
		return "-" + res, nil
	case *ast.CONSTANT:
		return e.TokenLiteral.Literal, nil
	case *ast.Identifier:
		operand := e.TokenLiteral.Literal
		if _, exist := this.identifiers[operand]; !exist {
			return "", fmt.Errorf("undeclared identifier %s", operand)
		}
		return this.identifiers[operand], nil
	}
	return "", nil
}

func (this *Evaluator) evaluateBinaryExpression(lhs string, operator ast.Token, rhs string) (string, error) { //TODO
	lhs_casted, err := strconv.ParseFloat(lhs, 64)
	if err != nil {
		return "", fmt.Errorf("couldn't convert %s\n to float", lhs)
	}
	rhs_casted, err := strconv.ParseFloat(rhs, 64)
	if err != nil {
		return "", fmt.Errorf("couldn't convert %s\n to float", rhs)
	}
	switch operator.Token {
	case ast.Plus:
		return strconv.FormatFloat(lhs_casted+rhs_casted, 'f', -1, 64), nil
	case ast.Minus:
		return strconv.FormatFloat(lhs_casted-rhs_casted, 'f', -1, 64), nil
	case ast.Multiplication:
		return strconv.FormatFloat(lhs_casted*rhs_casted, 'f', -1, 64), nil
	case ast.Division:
		return strconv.FormatFloat(lhs_casted/rhs_casted, 'f', -1, 64), nil
	default:
		return "", fmt.Errorf("unexpected token %s\n", operator.Literal)
	}
}

func (this *Evaluator) Evaluate(exp ast.Expression) (string, error) {
	if this.identifiers == nil {
		this.identifiers = make(map[string]string)
	}
	return this.visit(exp)
}
