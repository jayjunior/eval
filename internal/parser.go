package internal

import (
	"fmt"

	"github.com/jayjunior/eval/internal/ast"
)

var current = 0
var Tokens []ast.Token = nil
var parseError error = nil

func Parse(tokens []ast.Token) (ast.Expression, error) {
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
        return nil, fmt.Errorf("unexpected token '%s' at position %d", Tokens[current].Literal, current)
    }

    return exp, nil
}

func expression() ast.Expression {
    return term()
}

func term() ast.Expression {
    exp := factor()
    if parseError != nil {
        return nil
    }
    for !isAtEnd() && (match(ast.Minus) || match(ast.Plus)) {
        operator := consume()
        rhs := factor()
        if parseError != nil {
            return nil
        }
        exp = &ast.BinaryExpression{Lhs: exp, Operator: operator, Rhs: rhs}
    }
    return exp
}

func factor() ast.Expression {
    exp := unary()
    if parseError != nil {
        return nil
    }
    for !isAtEnd() && (match(ast.Division) || match(ast.Multiplication)) {
        operator := consume()
        rhs := unary()
        if parseError != nil {
            return nil
        }
        exp = &ast.BinaryExpression{Lhs: exp,Operator: operator,Rhs: rhs}
    }
    return exp
}

func unary() ast.Expression {
    if parseError != nil {
        return nil
    }
    if isAtEnd() {
        parseError = fmt.Errorf("unexpected end of input: expected number or expression")
        return nil
    }
    if match(ast.Number) || match(ast.Open_Parentheses) {
        return primary()
    }
    if match(ast.Minus) {
        op := consume()
        operand := unary()
        if parseError != nil {
            return nil
        }
        return &ast.UnaryExpression{Operator: op, Operand: operand}
    }
    parseError = fmt.Errorf("unexpected token '%s' at position %d: expected number, '(' or '-'", Tokens[current].Literal, current)
    return nil
}

func primary() ast.Expression {
    if parseError != nil {
        return nil
    }
    if isAtEnd() {
        parseError = fmt.Errorf("unexpected end of input: expected number or '('")
        return nil
    }
    if match(ast.Open_Parentheses) {
        consume()
        exp := expression()
        if parseError != nil {
            return nil
        }
        if isAtEnd() {
            parseError = fmt.Errorf("unexpected end of input: expected ')'")
            return nil
        }
        if !match(ast.Close_Parentheses) {
            parseError = fmt.Errorf("expected ')' at position %d, got '%s'", current, Tokens[current].Literal)
            return nil
        }
        consume()
        return exp
    }
    if match(ast.Number) {
        token := consume()
        return &ast.NumberLiteral{Literal: token.Literal}
    }
    parseError = fmt.Errorf("unexpected token '%s' at position %d: expected number or '('", Tokens[current].Literal, current)
    return nil
}

func isAtEnd() bool {
    return current >= len(Tokens)
}

func match(tokenType ast.TokenType) bool {
    if isAtEnd() {
        return false
    }
    return Tokens[current].Token == tokenType
}

func consume() ast.Token {
    res := Tokens[current]
    current++
    return res
}