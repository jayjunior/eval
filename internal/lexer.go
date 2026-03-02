package internal

import (
	"fmt"

	"github.com/jayjunior/eval/internal/ast"
)

var input = ""
var current_index = 0
var res = make([]ast.Token, 0)
var operators = map[byte]ast.TokenType{
	'+': ast.Plus,
	'-': ast.Minus,
	'*': ast.Multiplication,
	'/': ast.Division,
	'(': ast.Open_Parentheses,
	')': ast.Close_Parentheses,
	'=': ast.EQUAL,
}
var keywords = map[string]ast.TokenType{
	"var":   ast.VAR,
	"true":  ast.TRUE,
	"false": ast.FALSE,
}

func Tokenize(input_string string) ([]ast.Token, error) {
	input = input_string
	current_index = 0
	res = make([]ast.Token, 0)
	for !isEnd() {
		token := peek_char()
		if tokenType, exist := operators[token]; exist {
			operator(tokenType)
		} else if isDigit(rune(token)) {
			number()
		} else if isLetter(rune(token)) || token == '_' {
			word()
		} else if token == '\t' || token == ' ' {
			consume_char()
		} else {
			return nil, fmt.Errorf("Unrecognized character at position %d", current_index)
		}
	}

	return res, nil
}

func operator(tokenType ast.TokenType) {
	token := consume_char()
	res = append(res, ast.Token{Literal: string(token), Token: tokenType})
}

func peek_char() byte {
	return input[current_index]
}

func consume_char() byte {
	res := input[current_index]
	current_index++
	return res
}

func number() {
	digit := ""
	isFloat := false
	for !isEnd() && (isDigit(rune(peek_char())) || peek_char() == '.') {
		if peek_char() == '.' {
			digit += string(consume_char())
			isFloat = true
			break
		}
		digit += string(consume_char())
	}

	for !isEnd() && isFloat && isDigit(rune(peek_char())) {
		digit += string(consume_char())
	}
	res = append(res, ast.Token{Literal: digit, Token: ast.NUMBER_LITERAL})
}

func word() {
	result := ""
	for !isEnd() && (isLetter(rune(peek_char())) || peek_char() == '_') {
		result += string(consume_char())
	}
	if tokenType, exists := keywords[result]; exists {
		res = append(res, ast.Token{Literal: result, Token: tokenType})
	} else {
		res = append(res, ast.Token{Literal: result, Token: ast.IDENTIFIER_LITERAL})
	}
}

func isEnd() bool {
	return current_index >= len(input)
}

func isDigit(digit rune) bool {
	return digit >= '0' && digit <= '9'
}

func isLetter(letter rune) bool {
	return (letter >= 'a' && letter <= 'z') || (letter >= 'A' && letter <= 'Z')
}
