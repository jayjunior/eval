package ast

type Assignement struct {
	LHS IdentifierLiteral
	Rhs Expression
}

func (this *Assignement) accept() interface{} {
	return nil
}
