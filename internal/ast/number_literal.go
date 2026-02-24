package ast

type NumberLiteral struct {
	Literal string
}

func (this *NumberLiteral) accept() interface{} {
	return nil
}
