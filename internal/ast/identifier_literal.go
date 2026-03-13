package ast

type Identifier struct {
	TokenLiteral Token
}

func (this *Identifier) accept(visitor Visitor) (string, error) {
	return visitor.visit(this)
}
