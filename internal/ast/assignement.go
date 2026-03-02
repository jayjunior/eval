package ast

type Assignement struct {
	LHS Identifier
	Rhs Expression
}

func (this *Assignement) accept() interface{} {
	return nil
}
