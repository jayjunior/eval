package internal

import (
	"fmt"
)

type Expression interface {
	accept(visitor Visitor) interface{}
}

type Visitor interface {
	visit(expression Expression) interface{}
}

type NumberLiteral struct {
	literal string
}

type BinaryExpression struct {
	lhs      Expression
	operator Token
	rhs      Expression
}

type UnaryExpresson struct {
	operator Token
	operand  Expression
}

var current = 0
var Tokens []Token = nil
var parseError error = nil

func (this *BinaryExpression) accept(visitor Visitor) interface{} {
	return visitor.visit(this);
}

func (this *UnaryExpresson) accept(visitor Visitor) interface{} {
	return visitor.visit(this);
}

func (this *NumberLiteral) accept(visitor Visitor) interface{} {
	return visitor.visit(this);
}

func Parse(tokens []Token) (Expression, error) {
	current = 0
	parseError = nil
	Tokens = tokens

	if len(tokens) == 0 {
		return nil, fmt.Errorf("empty input: no tokens to parse")
	}

	exp := expression()
	if parseError != nil {
		return nil, parseError
	}

	if current < len(Tokens) {
		return nil, fmt.Errorf("unexpected token '%s' at position %d", Tokens[current].token, current)
	}

	return exp, nil
}

func expression() Expression {
	return term()
}

func term() Expression {
	exp := factor()
	if parseError != nil {
		return nil
	}
	for !isAtEnd() && (match(Minus) || match(Plus)) {
		operator := consume()
		rhs := factor()
		if parseError != nil {
			return nil
		}
		exp = &BinaryExpression{exp, operator, rhs}
	}
	return exp
}

func factor() Expression {
	exp := unary()
	if parseError != nil {
		return nil
	}
	for !isAtEnd() && (match(Division) || match(Multiplication)) {
		operator := consume()
		rhs := unary()
		if parseError != nil {
			return nil
		}
		exp = &BinaryExpression{exp, operator, rhs}
	}
	return exp
}

func unary() Expression {
	if parseError != nil {
		return nil
	}
	if isAtEnd() {
		parseError = fmt.Errorf("unexpected end of input: expected number or expression")
		return nil
	}
	if match(Number) || match(Open_Parentheses) {
		return primary()
	}
	if match(Minus) {
		op := consume()
		operand := unary()
		if parseError != nil {
			return nil
		}
		return &UnaryExpresson{op, operand}
	}
	parseError = fmt.Errorf("unexpected token '%s' at position %d: expected number, '(' or '-'", Tokens[current].token, current)
	return nil
}

func primary() Expression {
	if parseError != nil {
		return nil
	}
	if isAtEnd() {
		parseError = fmt.Errorf("unexpected end of input: expected number or '('")
		return nil
	}
	if match(Open_Parentheses) {
		consume()
		exp := expression()
		if parseError != nil {
			return nil
		}
		if isAtEnd() {
			parseError = fmt.Errorf("unexpected end of input: expected ')'")
			return nil
		}
		if !match(Close_Parentheses) {
			parseError = fmt.Errorf("expected ')' at position %d, got '%s'", current, Tokens[current].token)
			return nil
		}
		consume()
		return exp
	}
	if match(Number) {
		token := consume()
		return &NumberLiteral{token.token}
	}
	parseError = fmt.Errorf("unexpected token '%s' at position %d: expected number or '('", Tokens[current].token, current)
	return nil
}

func isAtEnd() bool {
	return current >= len(Tokens)
}

func match(token_type TokenType) bool {
	if isAtEnd() {
		return false
	}
	return Tokens[current].token_type == token_type
}

func consume() Token {
	res := Tokens[current]
	current++
	return res
}


