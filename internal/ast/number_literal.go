package ast

type NumberLiteral struct {
	TokenLiteral Token
}

func (this *NumberLiteral) accept() interface{} {
	return nil
}
