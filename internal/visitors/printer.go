package visitors;

import (
	"fmt"
	"strings"
	"github.com/jayjunior/eval/internal/ast"
)

type Printer struct {
	builder strings.Builder
}

func (this *Printer) visit(exp ast.Expression, prefix string, isLast bool) {
	connector := "├── "
	if isLast {
		connector = "└── "
	}

	childPrefix := prefix
	if isLast {
		childPrefix += "    "
	} else {
		childPrefix += "│   "
	}

	switch e := exp.(type) {
	case *ast.BinaryExpression:
		this.builder.WriteString(prefix + connector + "BinaryExpr (" + e.Operator.Literal + ")\n")
		this.visit(e.Lhs, childPrefix, false)
		this.visit(e.Rhs, childPrefix, true)
	case *ast.UnaryExpression:
		this.builder.WriteString(prefix + connector + "UnaryExpr (" + e.Operator.Literal + ")\n")
		this.visit(e.Operand, childPrefix, true)
	case *ast.NumberLiteral:
		this.builder.WriteString(prefix + connector + "Number: " + e.Literal + "\n")
	}
}

func (this *Printer) print(exp ast.Expression) string {
	this.builder.Reset()
	this.builder.WriteString("Expression\n")
	this.visit(exp, "", true)
	return this.builder.String()
}

func (this *Printer) PrintAST(exp ast.Expression) {
	fmt.Print(this.print(exp))
}

func CreatePrinter() *Printer {
	return &Printer{};
}
