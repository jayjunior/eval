package ast

import (
	"fmt"
	"os"
	"strconv"
)

type Evaluator struct {
	identifiers map[string]string
}

func (this *Evaluator) visit(exp Expression) (string, error) {
	switch e := exp.(type) {
	case *VarDeclaration:
		operand := e.Operand.TokenLiteral.Literal
		if _, exist := this.identifiers[operand]; exist {
			return "", fmt.Errorf("double declaration of %s", operand)
		}
		this.identifiers[operand] = "0"
		return "0", nil
	case *Assignement:
		operand := e.LHS.TokenLiteral.Literal
		rhs, err := e.Rhs.accept(this)
		if err != nil {
			return "", fmt.Errorf("error for assignement: %v", err)
		}
		this.identifiers[operand] = rhs
		return rhs, nil
	case *BinaryExpression:
		lhs, err := e.Lhs.accept(this)
		if err != nil {
			return "", err
		}
		rhs, err := e.Rhs.accept(this)
		if err != nil {
			return "", err
		}
		res, err := this.evaluateBinaryExpression(lhs, e.Operator, rhs)
		if err != nil {
			return "", err
		}
		return res, nil
	case *UnaryExpression:
		res, err := e.Operand.accept(this)
		if err != nil {
			return "", err
		}
		return "-" + res, nil
	case *CONSTANT:
		return e.TokenLiteral.Literal, nil
	case *Identifier:
		operand := e.TokenLiteral.Literal

		if _, exist := this.identifiers[operand]; exist {
			return this.identifiers[operand] , nil
		}
		if value , exist := os.LookupEnv(operand) ; exist {
			return value , nil;
		}
		return "", fmt.Errorf("undeclared identifier %s", operand)
	}
	return "", nil
}

func (this *Evaluator) evaluateBinaryExpression(lhs string, operator Token, rhs string) (string, error) { //TODO
	lhs_casted, err := strconv.ParseFloat(lhs, 64)
	if err != nil {
		return "", fmt.Errorf("couldn't convert %s\n to float", lhs)
	}
	rhs_casted, err := strconv.ParseFloat(rhs, 64)
	if err != nil {
		return "", fmt.Errorf("couldn't convert %s\n to float", rhs)
	}
	switch operator.Token {
	case Plus:
		return strconv.FormatFloat(lhs_casted+rhs_casted, 'f', -1, 64), nil
	case Minus:
		return strconv.FormatFloat(lhs_casted-rhs_casted, 'f', -1, 64), nil
	case Multiplication:
		return strconv.FormatFloat(lhs_casted*rhs_casted, 'f', -1, 64), nil
	case Division:
		return strconv.FormatFloat(lhs_casted/rhs_casted, 'f', -1, 64), nil
	default:
		return "", fmt.Errorf("unexpected token %s\n", operator.Literal)
	}
}

func (this *Evaluator) Evaluate(exp Expression) (string, error) {
	if this.identifiers == nil {
		this.identifiers = make(map[string]string)
	}
	return exp.accept(this)
}
