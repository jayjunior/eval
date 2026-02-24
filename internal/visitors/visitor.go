package visitors

import (
	"github.com/jayjunior/eval/internal/ast"
)
type Visitor interface {
	visit(expression ast.Expression) interface{}
}
