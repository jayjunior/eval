package ast

type VarDeclaration struct {
	Operand Identifier
}

func (this *VarDeclaration) accept() interface{} {
	return nil
}
