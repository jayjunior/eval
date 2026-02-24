package ast

type TokenType string

type Token struct {
	Literal string
	Token   TokenType
}

const (
	Plus              TokenType = "+"
	Minus             TokenType = "-"
	Multiplication    TokenType = "*"
	Division          TokenType = "/"
	Open_Parentheses  TokenType = "("
	Close_Parentheses TokenType = ")"
	Number            TokenType = "\\d+"
)
