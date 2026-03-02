package ast

type TokenType string

type Token struct {
	Literal string
	Token   TokenType
}

const (
	Plus               TokenType = "+"
	Minus              TokenType = "-"
	Multiplication     TokenType = "*"
	Division           TokenType = "/"
	Open_Parentheses   TokenType = "("
	Close_Parentheses  TokenType = ")"
	NUMBER_LITERAL     TokenType = "\\d*"
	IDENTIFIER_LITERAL TokenType = "_[a-zA-Z]"
	EQUAL              TokenType = "="
	VAR                TokenType = "var"
)
