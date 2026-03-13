package ast

type CONSTANT struct {
	TokenLiteral Token
}

func (this *CONSTANT) accept(visitor Visitor) (string, error) {
	return visitor.visit(this)
}
