package ast

type VarDeclaration struct {
	Operand IdentifierLiteral
}

func (this *VarDeclaration) accept() interface{} {
	return nil
}
