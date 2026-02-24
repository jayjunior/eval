package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/jayjunior/eval/internal"
	"github.com/jayjunior/eval/internal/visitors"
)

func main() {
	if len(os.Args) >= 2 {
		evaluateExpression(os.Args[1], true)
		return
	}

	fmt.Fprintln(os.Stdout, "Welcome to the eval repl")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		if strings.ToLower(line) == "exit" {
			fmt.Println("Bye :)")
			os.Exit(0)
		}
		evaluateExpression(line, false)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
	}
}

func evaluateExpression(expression string, exitOnError bool) {
	tokens, err := internal.Tokenize(expression)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Lexer error: %v\n", err)
		if exitOnError {
			os.Exit(1)
		}
		return
	}

	exprAst, err := internal.Parse(tokens)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Parser error: %v\n", err)
		if exitOnError {
			os.Exit(1)
		}
		return
	}

	evaluator := visitors.Evaluator{}
	res, err := evaluator.Evaluate(exprAst)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error evaluating the expression: %v\n", err)
		if exitOnError {
			os.Exit(1)
		}
		return
	}
	fmt.Println(res)
}
