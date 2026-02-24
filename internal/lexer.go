package internal

import (
	"fmt"
	"unicode"

	"github.com/jayjunior/eval/internal/ast"
)

func Tokenize(input string) ([]ast.Token, error) {
	res := make([]ast.Token, 0)

	for i := 0; i < len(input); {
		token := input[i]
		switch token {
		case '+':
			res = append(res, ast.Token{Literal: string(token), Token: ast.Plus})
			i++
			continue
		case '-':
			res = append(res, ast.Token{Literal: string(token), Token: ast.Minus})
			i++
			continue

		case '*':
			res = append(res, ast.Token{Literal: string(token), Token: ast.Multiplication})
			i++
			continue

		case '/':
			res = append(res, ast.Token{Literal: string(token), Token: ast.Division})
			i++
			continue

		case '(':
			res = append(res, ast.Token{Literal: string(token), Token: ast.Open_Parentheses})
			i++
			continue

		case ')':
			res = append(res, ast.Token{Literal: string(token), Token: ast.Close_Parentheses})
			i++
			continue

		case '\t', ' ':
			i++
			continue
		}
		if unicode.IsDigit(rune(token)) {
			digit := ""
			for i < len(input) && unicode.IsDigit(rune(input[i])) {
				digit += string(input[i])
				i++
			}
			res = append(res, ast.Token{Literal: digit, Token: ast.Number})
		} else {
			return nil, fmt.Errorf("Unrecognized character at position %d", i)
		}
	}

	return res, nil
}
