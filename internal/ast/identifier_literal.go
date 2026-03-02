package ast

type Identifier struct {
	TokenLiteral Token
}

func (this *Identifier) accept() interface{} {
	return nil
}
