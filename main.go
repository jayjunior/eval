package main

import (
	"fmt"
	"os"
	"github.com/jayjunior/eval/internal/visitors"
	"github.com/jayjunior/eval/internal"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: eval <expression>")
		os.Exit(1)
	}

	expression := os.Args[1]

	tokens, err := internal.Tokenize(expression)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Lexer error: %v\n", err)
		os.Exit(1)
	}

	ast, err := internal.Parse(tokens)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Parser error: %v\n", err)
		os.Exit(1)
	}

	printer := visitors.CreatePrinter()
	printer.PrintAST(ast)

	evaluator := visitors.Evaluator{}
	res, err := evaluator.Evaluate(ast)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error evaluating the expression: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(res)
}
