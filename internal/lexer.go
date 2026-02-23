package internal

import (
	"fmt"
	"unicode"
)

func Tokenize(input string) ([]Token, error) {
	res := make([]Token, 0)

	for i := 0; i < len(input); {
		token := input[i]
		switch token {
		case '+':
			res = append(res, Token{string(token), Plus})
			i++
			continue
		case '-':
			res = append(res, Token{string(token), Minus})
			i++
			continue

		case '*':
			res = append(res, Token{string(token), Multiplication})
			i++
			continue

		case '/':
			res = append(res, Token{string(token), Division})
			i++
			continue

		case '(':
			res = append(res, Token{string(token), Open_Parentheses})
			i++
			continue

		case ')':
			res = append(res, Token{string(token), Close_Parentheses})
			i++
			continue

		case '\t':
			i++
			continue

		case ' ':
			i++
			continue

		}
		if unicode.IsDigit(rune(token)) {
			digit := ""
			for i < len(input) && unicode.IsDigit(rune(input[i])) {
				digit += string(input[i])
				i++
			}
			res = append(res, Token{digit, Number})
		} else {
			return nil, fmt.Errorf("Unrecognized character at position %d", i)
		}
	}

	return res, nil

}
