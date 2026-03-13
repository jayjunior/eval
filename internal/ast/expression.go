package ast

type Expression interface {
	accept(visitor Visitor) (string, error)
}
