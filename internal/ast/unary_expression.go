package ast

type UnaryExpression struct {
	Operator Token
	Operand  Expression
}

func (this *UnaryExpression) accept(visitor Visitor) (string, error) {
	return visitor.visit(this)
}
