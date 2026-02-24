package internal_test

import (
	"strings"
	"testing"

	"github.com/jayjunior/eval/internal"
	"github.com/jayjunior/eval/internal/ast"
)

// Helper to create tokens for testing
func tokens(input string) []ast.Token {
	t, _ := internal.Tokenize(input)
	return t
}

// Valid expression tests

func TestParseNumber(t *testing.T) {
	exp, err := internal.Parse(tokens("5"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	num, ok := exp.(*ast.NumberLiteral)
	if !ok {
		t.Fatalf("expected NumberLiteral, got %T", exp)
	}
	if num.Literal != "5" {
		t.Errorf("expected '5', got '%s'", num.Literal)
	}
}

func TestParseMultiDigitNumber(t *testing.T) {
	exp, err := internal.Parse(tokens("123"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	num, ok := exp.(*ast.NumberLiteral)
	if !ok {
		t.Fatalf("expected NumberLiteral, got %T", exp)
	}
	if num.Literal != "123" {
		t.Errorf("expected '123', got '%s'", num.Literal)
	}
}

func TestParseAddition(t *testing.T) {
	exp, err := internal.Parse(tokens("1+2"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	bin, ok := exp.(*ast.BinaryExpression)
	if !ok {
		t.Fatalf("expected BinaryExpression, got %T", exp)
	}
	if bin.Operator.Token != ast.Plus {
		t.Errorf("expected Plus operator, got %v", bin.Operator.Token)
	}
}

func TestParseSubtraction(t *testing.T) {
	exp, err := internal.Parse(tokens("5-3"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	bin, ok := exp.(*ast.BinaryExpression)
	if !ok {
		t.Fatalf("expected BinaryExpression, got %T", exp)
	}
	if bin.Operator.Token != ast.Minus {
		t.Errorf("expected Minus operator, got %v", bin.Operator.Token)
	}
}

func TestParseMultiplication(t *testing.T) {
	exp, err := internal.Parse(tokens("2*3"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	bin, ok := exp.(*ast.BinaryExpression)
	if !ok {
		t.Fatalf("expected BinaryExpression, got %T", exp)
	}
	if bin.Operator.Token != ast.Multiplication {
		t.Errorf("expected Multiplication operator, got %v", bin.Operator.Token)
	}
}

func TestParseDivision(t *testing.T) {
	exp, err := internal.Parse(tokens("6/2"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	bin, ok := exp.(*ast.BinaryExpression)
	if !ok {
		t.Fatalf("expected BinaryExpression, got %T", exp)
	}
	if bin.Operator.Token != ast.Division {
		t.Errorf("expected Division operator, got %v", bin.Operator.Token)
	}
}

func TestParseUnaryMinus(t *testing.T) {
	exp, err := internal.Parse(tokens("-5"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	unary, ok := exp.(*ast.UnaryExpression)
	if !ok {
		t.Fatalf("expected UnaryExpression, got %T", exp)
	}
	if unary.Operator.Token != ast.Minus {
		t.Errorf("expected Minus operator, got %v", unary.Operator.Token)
	}
}

func TestParseDoubleUnaryMinus(t *testing.T) {
	exp, err := internal.Parse(tokens("--5"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	outer, ok := exp.(*ast.UnaryExpression)
	if !ok {
		t.Fatalf("expected UnaryExpression, got %T", exp)
	}
	inner, ok := outer.Operand.(*ast.UnaryExpression)
	if !ok {
		t.Fatalf("expected nested UnaryExpression, got %T", outer.Operand)
	}
	_, ok = inner.Operand.(*ast.NumberLiteral)
	if !ok {
		t.Fatalf("expected NumberLiteral, got %T", inner.Operand)
	}
}

func TestParseParentheses(t *testing.T) {
	exp, err := internal.Parse(tokens("(5)"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	num, ok := exp.(*ast.NumberLiteral)
	if !ok {
		t.Fatalf("expected NumberLiteral, got %T", exp)
	}
	if num.Literal != "5" {
		t.Errorf("expected '5', got '%s'", num.Literal)
	}
}

func TestParseNestedParentheses(t *testing.T) {
	exp, err := internal.Parse(tokens("((5))"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	num, ok := exp.(*ast.NumberLiteral)
	if !ok {
		t.Fatalf("expected NumberLiteral, got %T", exp)
	}
	if num.Literal != "5" {
		t.Errorf("expected '5', got '%s'", num.Literal)
	}
}

func TestParseParenthesizedExpression(t *testing.T) {
	exp, err := internal.Parse(tokens("(1+2)"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	bin, ok := exp.(*ast.BinaryExpression)
	if !ok {
		t.Fatalf("expected BinaryExpression, got %T", exp)
	}
	if bin.Operator.Token != ast.Plus {
		t.Errorf("expected Plus operator, got %v", bin.Operator.Token)
	}
}

// Operator precedence tests

func TestParsePrecedenceMultiplicationBeforeAddition(t *testing.T) {
	// 1+2*3 should be parsed as 1+(2*3)
	exp, err := internal.Parse(tokens("1+2*3"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	bin, ok := exp.(*ast.BinaryExpression)
	if !ok {
		t.Fatalf("expected BinaryExpression, got %T", exp)
	}
	if bin.Operator.Token != ast.Plus {
		t.Errorf("expected Plus as root operator, got %v", bin.Operator.Token)
	}
	// RHS should be multiplication
	rhs, ok := bin.Rhs.(*ast.BinaryExpression)
	if !ok {
		t.Fatalf("expected BinaryExpression on rhs, got %T", bin.Rhs)
	}
	if rhs.Operator.Token != ast.Multiplication {
		t.Errorf("expected Multiplication on rhs, got %v", rhs.Operator.Token)
	}
}

func TestParsePrecedenceDivisionBeforeSubtraction(t *testing.T) {
	// 6-4/2 should be parsed as 6-(4/2)
	exp, err := internal.Parse(tokens("6-4/2"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	bin, ok := exp.(*ast.BinaryExpression)
	if !ok {
		t.Fatalf("expected BinaryExpression, got %T", exp)
	}
	if bin.Operator.Token != ast.Minus {
		t.Errorf("expected Minus as root operator, got %v", bin.Operator.Token)
	}
	rhs, ok := bin.Rhs.(*ast.BinaryExpression)
	if !ok {
		t.Fatalf("expected BinaryExpression on rhs, got %T", bin.Rhs)
	}
	if rhs.Operator.Token != ast.Division {
		t.Errorf("expected Division on rhs, got %v", rhs.Operator.Token)
	}
}

func TestParsePrecedenceParenthesesOverride(t *testing.T) {
	// (1+2)*3 should be parsed with addition first
	exp, err := internal.Parse(tokens("(1+2)*3"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	bin, ok := exp.(*ast.BinaryExpression)
	if !ok {
		t.Fatalf("expected BinaryExpression, got %T", exp)
	}
	if bin.Operator.Token != ast.Multiplication {
		t.Errorf("expected Multiplication as root operator, got %v", bin.Operator.Token)
	}
	// LHS should be addition
	lhs, ok := bin.Lhs.(*ast.BinaryExpression)
	if !ok {
		t.Fatalf("expected BinaryExpression on lhs, got %T", bin.Lhs)
	}
	if lhs.Operator.Token != ast.Plus {
		t.Errorf("expected Plus on lhs, got %v", lhs.Operator.Token)
	}
}

func TestParseComplexExpression(t *testing.T) {
	// 1+2*3-4/2
	exp, err := internal.Parse(tokens("1+2*3-4/2"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if exp == nil {
		t.Fatal("expected expression, got nil")
	}
}

func TestParseUnaryWithBinary(t *testing.T) {
	// 5+-3 (5 + (-3))
	exp, err := internal.Parse(tokens("5+-3"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	bin, ok := exp.(*ast.BinaryExpression)
	if !ok {
		t.Fatalf("expected BinaryExpression, got %T", exp)
	}
	if bin.Operator.Token != ast.Plus {
		t.Errorf("expected Plus operator, got %v", bin.Operator.Token)
	}
	unary, ok := bin.Rhs.(*ast.UnaryExpression)
	if !ok {
		t.Fatalf("expected UnaryExpression on rhs, got %T", bin.Rhs)
	}
	if unary.Operator.Token != ast.Minus {
		t.Errorf("expected Minus operator on unary, got %v", unary.Operator.Token)
	}
}

func TestParseUnaryInParentheses(t *testing.T) {
	// -(-5)
	exp, err := internal.Parse(tokens("-(-5)"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	outer, ok := exp.(*ast.UnaryExpression)
	if !ok {
		t.Fatalf("expected UnaryExpression, got %T", exp)
	}
	inner, ok := outer.Operand.(*ast.UnaryExpression)
	if !ok {
		t.Fatalf("expected UnaryExpression inside parentheses, got %T", outer.Operand)
	}
	if inner.Operator.Token != ast.Minus {
		t.Errorf("expected Minus on inner unary, got %v", inner.Operator.Token)
	}
}

// Error handling tests

func TestParseEmptyInput(t *testing.T) {
	_, err := internal.Parse([]ast.Token{})
	if err == nil {
		t.Error("expected error for empty input, got nil")
	}
	if !strings.Contains(err.Error(), "empty input") {
		t.Errorf("expected 'empty input' in error message, got: %v", err)
	}
}

func TestParseMissingClosingParenthesis(t *testing.T) {
	_, err := internal.Parse(tokens("(1+2"))
	if err == nil {
		t.Error("expected error for missing closing parenthesis, got nil")
	}
	if !strings.Contains(err.Error(), "')'") {
		t.Errorf("expected mention of ')' in error message, got: %v", err)
	}
}

func TestParseMissingOpeningParenthesis(t *testing.T) {
	_, err := internal.Parse(tokens("1+2)"))
	if err == nil {
		t.Error("expected error for extra closing parenthesis, got nil")
	}
}

func TestParseUnmatchedParentheses(t *testing.T) {
	_, err := internal.Parse(tokens("((1+2)"))
	if err == nil {
		t.Error("expected error for unmatched parentheses, got nil")
	}
}

func TestParseTrailingOperator(t *testing.T) {
	_, err := internal.Parse(tokens("1+"))
	if err == nil {
		t.Error("expected error for trailing operator, got nil")
	}
	if !strings.Contains(err.Error(), "end of input") {
		t.Errorf("expected 'end of input' in error message, got: %v", err)
	}
}

func TestParseLeadingOperatorPlus(t *testing.T) {
	_, err := internal.Parse(tokens("+5"))
	if err == nil {
		t.Error("expected error for leading + operator, got nil")
	}
}

func TestParseConsecutiveBinaryOperators(t *testing.T) {
	_, err := internal.Parse(tokens("1++2"))
	if err == nil {
		t.Error("expected error for consecutive binary operators, got nil")
	}
}

func TestParseOnlyOperator(t *testing.T) {
	_, err := internal.Parse(tokens("+"))
	if err == nil {
		t.Error("expected error for only operator, got nil")
	}
}

func TestParseOnlyParentheses(t *testing.T) {
	_, err := internal.Parse(tokens("()"))
	if err == nil {
		t.Error("expected error for empty parentheses, got nil")
	}
}

func TestParseMultipleBinaryInRow(t *testing.T) {
	_, err := internal.Parse(tokens("1*/2"))
	if err == nil {
		t.Error("expected error for multiple binary operators in a row, got nil")
	}
}

// Edge cases

func TestParseLeftAssociativity(t *testing.T) {
	// 1-2-3 should be parsed as (1-2)-3, not 1-(2-3)
	exp, err := internal.Parse(tokens("1-2-3"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	bin, ok := exp.(*ast.BinaryExpression)
	if !ok {
		t.Fatalf("expected BinaryExpression, got %T", exp)
	}
	// Root should be rightmost minus
	if bin.Operator.Token != ast.Minus {
		t.Errorf("expected Minus as root, got %v", bin.Operator.Token)
	}
	// LHS should be (1-2)
	lhs, ok := bin.Lhs.(*ast.BinaryExpression)
	if !ok {
		t.Fatalf("expected BinaryExpression on lhs (left associativity), got %T", bin.Lhs)
	}
	if lhs.Operator.Token != ast.Minus {
		t.Errorf("expected Minus on lhs, got %v", lhs.Operator.Token)
	}
}

func TestParseDeeplyNested(t *testing.T) {
	_, err := internal.Parse(tokens("(((((1)))))"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestParseZero(t *testing.T) {
	exp, err := internal.Parse(tokens("0"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	num, ok := exp.(*ast.NumberLiteral)
	if !ok {
		t.Fatalf("expected NumberLiteral, got %T", exp)
	}
	if num.Literal != "0" {
		t.Errorf("expected '0', got '%s'", num.Literal)
	}
}

func TestParseLargeNumber(t *testing.T) {
	exp, err := internal.Parse(tokens("123456789"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	num, ok := exp.(*ast.NumberLiteral)
	if !ok {
		t.Fatalf("expected NumberLiteral, got %T", exp)
	}
	if num.Literal != "123456789" {
		t.Errorf("expected '123456789', got '%s'", num.Literal)
	}
}

func TestParseMultipleParses(t *testing.T) {
	// Test that parser state resets correctly between parses
	_, err1 := internal.Parse(tokens("1+2"))
	if err1 != nil {
		t.Fatalf("first parse failed: %v", err1)
	}

	_, err2 := internal.Parse(tokens("3*4"))
	if err2 != nil {
		t.Fatalf("second parse failed: %v", err2)
	}

	// Error parse should not affect subsequent parses
	internal.Parse(tokens("("))

	_, err3 := internal.Parse(tokens("5-6"))
	if err3 != nil {
		t.Fatalf("third parse failed after error: %v", err3)
	}
}
