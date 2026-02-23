package internal

import (
	"fmt"
	"strings"
)

type Printer struct {
	builder strings.Builder
}

func (this *Printer) visit(exp Expression, prefix string, isLast bool) {
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
	case *BinaryExpression:
		this.builder.WriteString(prefix + connector + "BinaryExpr (" + e.operator.token + ")\n")
		this.visit(e.lhs, childPrefix, false)
		this.visit(e.rhs, childPrefix, true)
	case *UnaryExpresson:
		this.builder.WriteString(prefix + connector + "UnaryExpr (" + e.operator.token + ")\n")
		this.visit(e.operand, childPrefix, true)
	case *NumberLiteral:
		this.builder.WriteString(prefix + connector + "Number: " + e.literal + "\n")
	}
}

func (this *Printer) print(exp Expression) string {
	this.builder.Reset()
	this.builder.WriteString("Expression\n")
	this.visit(exp, "", true)
	return this.builder.String()
}

func (this *Printer) PrintAST(exp Expression) {
	fmt.Print(this.print(exp))
}

func CreatePrinter() *Printer {
	return &Printer{};
}
