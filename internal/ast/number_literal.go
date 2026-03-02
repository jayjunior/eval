package ast

type CONSTANT struct {
	TokenLiteral Token
}

func (this *CONSTANT) accept() interface{} {
	return nil
}
