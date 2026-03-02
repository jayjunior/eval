package ast

type Number struct {
	TokenLiteral Token
}

func (this *Number) accept() interface{} {
	return nil
}
