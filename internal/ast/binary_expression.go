package ast

type BinaryExpression struct {
    Lhs      Expression
    Operator Token
    Rhs      Expression
}

func (this *BinaryExpression) accept() interface{} {
    return nil
}