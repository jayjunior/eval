package ast

type UnaryExpression struct {
    Operator Token
    Operand  Expression
}

func (this *UnaryExpression) accept() interface{} {
    return nil
}
