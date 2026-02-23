package internal

type TokenType string
type Token struct {
	token      string
	token_type TokenType
}

const (
	Plus              TokenType = "+"
	Minus             TokenType = "-"
	Multiplication    TokenType = "*"
	Division          TokenType = "/"
	Open_Parentheses  TokenType = "("
	Close_Parentheses TokenType = ")"
	Number            TokenType = "\\d"
)
