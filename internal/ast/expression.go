package ast

type Expression interface {
	accept() interface{}
}