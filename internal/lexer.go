package internal

import (
	"fmt"
	"unicode"

	"github.com/jayjunior/eval/internal/ast"
)

var input = ""
var current_index = 0
var res = make([]ast.Token, 0)

func Tokenize(input_string string) ([]ast.Token, error) {
	input = input_string
	current_index = 0
	res = make([]ast.Token, 0)
	for !isEnd() {
		token := peek_char()
		if token == '+' || token == '-' || token == '*' || token == '/' || token == '(' || token == ')' || token == '=' {
			operator()
		} else if unicode.IsDigit(rune(token)) {
			number()
		} else if unicode.IsLetter(rune(token)) || token == '_' {
			word()
		} else if token == '\t' || token == ' ' {
			consume_char()
			continue
		} else {
			return nil, fmt.Errorf("Unrecognized character at position %d", current_index)
		}
	}

	return res, nil
}

func operator() {
	token := consume_char()
	switch token {
	case '+':
		res = append(res, ast.Token{Literal: string(token), Token: ast.Plus})
	case '-':
		res = append(res, ast.Token{Literal: string(token), Token: ast.Minus})
	case '*':
		res = append(res, ast.Token{Literal: string(token), Token: ast.Multiplication})
	case '/':
		res = append(res, ast.Token{Literal: string(token), Token: ast.Division})
	case '(':
		res = append(res, ast.Token{Literal: string(token), Token: ast.Open_Parentheses})
	case ')':
		res = append(res, ast.Token{Literal: string(token), Token: ast.Close_Parentheses})
	case '=':
		res = append(res, ast.Token{Literal: string(token), Token: ast.EQUAL})
	}
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
	for !isEnd() && unicode.IsDigit(rune(peek_char())) {
		digit += string(consume_char())
	}
	res = append(res, ast.Token{Literal: digit, Token: ast.NUMBER})
}

func word() {
	result := ""
	for !isEnd() && (unicode.IsLetter(rune(peek_char())) || peek_char() == '_') { //TODO use a list of valid identifier characters
		result += string(consume_char())
	}
	if result == string(ast.VAR) { // TODO use a map for keywords
		res = append(res, ast.Token{Literal: result, Token: ast.VAR})
	} else {
		res = append(res, ast.Token{Literal: result, Token: ast.IDENTIFIER})
	}
}

func isEnd() bool {
	return current_index >= len(input)
}
