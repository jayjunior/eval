package ast

type VarDeclaration struct {
	Operand Identifier
}

func (this *VarDeclaration) accept(visitor Visitor) (string, error) {
	return visitor.visit(this)
}
