package ast

type Visitor interface {
	visit(expression Expression) (string, error)
}
