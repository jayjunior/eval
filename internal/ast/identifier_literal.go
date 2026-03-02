package ast

type IdentifierLiteral struct {
	TokenLiteral Token
}

func (this *IdentifierLiteral) accept() interface{} {
	return nil
}
