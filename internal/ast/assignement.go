package ast

type Assignement struct {
	LHS Identifier
	Rhs Expression
}

func (this *Assignement) accept(visitor Visitor) (string, error) {
	return visitor.visit(this)
}
