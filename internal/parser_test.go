package internal

import (
	"strings"
	"testing"
)

// Helper to create tokens for testing
func tokens(input string) []Token {
	t, _ := Tokenize(input)
	return t
}

// Valid expression tests

func TestParseNumber(t *testing.T) {
	exp, err := Parse(tokens("5"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	num, ok := exp.(*NumberLiteral)
	if !ok {
		t.Fatalf("expected NumberLiteral, got %T", exp)
	}
	if num.literal != "5" {
		t.Errorf("expected '5', got '%s'", num.literal)
	}
}

func TestParseMultiDigitNumber(t *testing.T) {
	exp, err := Parse(tokens("123"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	num, ok := exp.(*NumberLiteral)
	if !ok {
		t.Fatalf("expected NumberLiteral, got %T", exp)
	}
	if num.literal != "123" {
		t.Errorf("expected '123', got '%s'", num.literal)
	}
}

func TestParseAddition(t *testing.T) {
	exp, err := Parse(tokens("1+2"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	bin, ok := exp.(*BinaryExpression)
	if !ok {
		t.Fatalf("expected BinaryExpression, got %T", exp)
	}
	if bin.operator.token_type != Plus {
		t.Errorf("expected Plus operator, got %v", bin.operator.token_type)
	}
}

func TestParseSubtraction(t *testing.T) {
	exp, err := Parse(tokens("5-3"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	bin, ok := exp.(*BinaryExpression)
	if !ok {
		t.Fatalf("expected BinaryExpression, got %T", exp)
	}
	if bin.operator.token_type != Minus {
		t.Errorf("expected Minus operator, got %v", bin.operator.token_type)
	}
}

func TestParseMultiplication(t *testing.T) {
	exp, err := Parse(tokens("2*3"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	bin, ok := exp.(*BinaryExpression)
	if !ok {
		t.Fatalf("expected BinaryExpression, got %T", exp)
	}
	if bin.operator.token_type != Multiplication {
		t.Errorf("expected Multiplication operator, got %v", bin.operator.token_type)
	}
}

func TestParseDivision(t *testing.T) {
	exp, err := Parse(tokens("6/2"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	bin, ok := exp.(*BinaryExpression)
	if !ok {
		t.Fatalf("expected BinaryExpression, got %T", exp)
	}
	if bin.operator.token_type != Division {
		t.Errorf("expected Division operator, got %v", bin.operator.token_type)
	}
}

func TestParseUnaryMinus(t *testing.T) {
	exp, err := Parse(tokens("-5"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	unary, ok := exp.(*UnaryExpresson)
	if !ok {
		t.Fatalf("expected UnaryExpresson, got %T", exp)
	}
	if unary.operator.token_type != Minus {
		t.Errorf("expected Minus operator, got %v", unary.operator.token_type)
	}
}

func TestParseDoubleUnaryMinus(t *testing.T) {
	exp, err := Parse(tokens("--5"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	outer, ok := exp.(*UnaryExpresson)
	if !ok {
		t.Fatalf("expected UnaryExpresson, got %T", exp)
	}
	inner, ok := outer.operand.(*UnaryExpresson)
	if !ok {
		t.Fatalf("expected nested UnaryExpresson, got %T", outer.operand)
	}
	_, ok = inner.operand.(*NumberLiteral)
	if !ok {
		t.Fatalf("expected NumberLiteral, got %T", inner.operand)
	}
}

func TestParseParentheses(t *testing.T) {
	exp, err := Parse(tokens("(5)"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	num, ok := exp.(*NumberLiteral)
	if !ok {
		t.Fatalf("expected NumberLiteral, got %T", exp)
	}
	if num.literal != "5" {
		t.Errorf("expected '5', got '%s'", num.literal)
	}
}

func TestParseNestedParentheses(t *testing.T) {
	exp, err := Parse(tokens("((5))"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	num, ok := exp.(*NumberLiteral)
	if !ok {
		t.Fatalf("expected NumberLiteral, got %T", exp)
	}
	if num.literal != "5" {
		t.Errorf("expected '5', got '%s'", num.literal)
	}
}

func TestParseParenthesizedExpression(t *testing.T) {
	exp, err := Parse(tokens("(1+2)"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	bin, ok := exp.(*BinaryExpression)
	if !ok {
		t.Fatalf("expected BinaryExpression, got %T", exp)
	}
	if bin.operator.token_type != Plus {
		t.Errorf("expected Plus operator, got %v", bin.operator.token_type)
	}
}

// Operator precedence tests

func TestParsePrecedenceMultiplicationBeforeAddition(t *testing.T) {
	// 1+2*3 should be parsed as 1+(2*3)
	exp, err := Parse(tokens("1+2*3"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	bin, ok := exp.(*BinaryExpression)
	if !ok {
		t.Fatalf("expected BinaryExpression, got %T", exp)
	}
	if bin.operator.token_type != Plus {
		t.Errorf("expected Plus as root operator, got %v", bin.operator.token_type)
	}
	// RHS should be multiplication
	rhs, ok := bin.rhs.(*BinaryExpression)
	if !ok {
		t.Fatalf("expected BinaryExpression on rhs, got %T", bin.rhs)
	}
	if rhs.operator.token_type != Multiplication {
		t.Errorf("expected Multiplication on rhs, got %v", rhs.operator.token_type)
	}
}

func TestParsePrecedenceDivisionBeforeSubtraction(t *testing.T) {
	// 6-4/2 should be parsed as 6-(4/2)
	exp, err := Parse(tokens("6-4/2"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	bin, ok := exp.(*BinaryExpression)
	if !ok {
		t.Fatalf("expected BinaryExpression, got %T", exp)
	}
	if bin.operator.token_type != Minus {
		t.Errorf("expected Minus as root operator, got %v", bin.operator.token_type)
	}
	rhs, ok := bin.rhs.(*BinaryExpression)
	if !ok {
		t.Fatalf("expected BinaryExpression on rhs, got %T", bin.rhs)
	}
	if rhs.operator.token_type != Division {
		t.Errorf("expected Division on rhs, got %v", rhs.operator.token_type)
	}
}

func TestParsePrecedenceParenthesesOverride(t *testing.T) {
	// (1+2)*3 should be parsed with addition first
	exp, err := Parse(tokens("(1+2)*3"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	bin, ok := exp.(*BinaryExpression)
	if !ok {
		t.Fatalf("expected BinaryExpression, got %T", exp)
	}
	if bin.operator.token_type != Multiplication {
		t.Errorf("expected Multiplication as root operator, got %v", bin.operator.token_type)
	}
	// LHS should be addition
	lhs, ok := bin.lhs.(*BinaryExpression)
	if !ok {
		t.Fatalf("expected BinaryExpression on lhs, got %T", bin.lhs)
	}
	if lhs.operator.token_type != Plus {
		t.Errorf("expected Plus on lhs, got %v", lhs.operator.token_type)
	}
}

func TestParseComplexExpression(t *testing.T) {
	// 1+2*3-4/2
	exp, err := Parse(tokens("1+2*3-4/2"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if exp == nil {
		t.Fatal("expected expression, got nil")
	}
}

func TestParseUnaryWithBinary(t *testing.T) {
	// 5+-3 (5 + (-3))
	exp, err := Parse(tokens("5+-3"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	bin, ok := exp.(*BinaryExpression)
	if !ok {
		t.Fatalf("expected BinaryExpression, got %T", exp)
	}
	if bin.operator.token_type != Plus {
		t.Errorf("expected Plus operator, got %v", bin.operator.token_type)
	}
	unary, ok := bin.rhs.(*UnaryExpresson)
	if !ok {
		t.Fatalf("expected UnaryExpresson on rhs, got %T", bin.rhs)
	}
	if unary.operator.token_type != Minus {
		t.Errorf("expected Minus operator on unary, got %v", unary.operator.token_type)
	}
}

func TestParseUnaryInParentheses(t *testing.T) {
	// -(-5)
	exp, err := Parse(tokens("-(-5)"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	outer, ok := exp.(*UnaryExpresson)
	if !ok {
		t.Fatalf("expected UnaryExpresson, got %T", exp)
	}
	inner, ok := outer.operand.(*UnaryExpresson)
	if !ok {
		t.Fatalf("expected UnaryExpresson inside parentheses, got %T", outer.operand)
	}
	if inner.operator.token_type != Minus {
		t.Errorf("expected Minus on inner unary, got %v", inner.operator.token_type)
	}
}

// Error handling tests

func TestParseEmptyInput(t *testing.T) {
	_, err := Parse([]Token{})
	if err == nil {
		t.Error("expected error for empty input, got nil")
	}
	if !strings.Contains(err.Error(), "empty input") {
		t.Errorf("expected 'empty input' in error message, got: %v", err)
	}
}

func TestParseMissingClosingParenthesis(t *testing.T) {
	_, err := Parse(tokens("(1+2"))
	if err == nil {
		t.Error("expected error for missing closing parenthesis, got nil")
	}
	if !strings.Contains(err.Error(), "')'") {
		t.Errorf("expected mention of ')' in error message, got: %v", err)
	}
}

func TestParseMissingOpeningParenthesis(t *testing.T) {
	_, err := Parse(tokens("1+2)"))
	if err == nil {
		t.Error("expected error for extra closing parenthesis, got nil")
	}
}

func TestParseUnmatchedParentheses(t *testing.T) {
	_, err := Parse(tokens("((1+2)"))
	if err == nil {
		t.Error("expected error for unmatched parentheses, got nil")
	}
}

func TestParseTrailingOperator(t *testing.T) {
	_, err := Parse(tokens("1+"))
	if err == nil {
		t.Error("expected error for trailing operator, got nil")
	}
	if !strings.Contains(err.Error(), "end of input") {
		t.Errorf("expected 'end of input' in error message, got: %v", err)
	}
}

func TestParseLeadingOperatorPlus(t *testing.T) {
	_, err := Parse(tokens("+5"))
	if err == nil {
		t.Error("expected error for leading + operator, got nil")
	}
}

func TestParseConsecutiveBinaryOperators(t *testing.T) {
	_, err := Parse(tokens("1++2"))
	if err == nil {
		t.Error("expected error for consecutive binary operators, got nil")
	}
}

func TestParseOnlyOperator(t *testing.T) {
	_, err := Parse(tokens("+"))
	if err == nil {
		t.Error("expected error for only operator, got nil")
	}
}

func TestParseOnlyParentheses(t *testing.T) {
	_, err := Parse(tokens("()"))
	if err == nil {
		t.Error("expected error for empty parentheses, got nil")
	}
}

func TestParseMultipleBinaryInRow(t *testing.T) {
	_, err := Parse(tokens("1*/2"))
	if err == nil {
		t.Error("expected error for multiple binary operators in a row, got nil")
	}
}

// Edge cases

func TestParseLeftAssociativity(t *testing.T) {
	// 1-2-3 should be parsed as (1-2)-3, not 1-(2-3)
	exp, err := Parse(tokens("1-2-3"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	bin, ok := exp.(*BinaryExpression)
	if !ok {
		t.Fatalf("expected BinaryExpression, got %T", exp)
	}
	// Root should be rightmost minus
	if bin.operator.token_type != Minus {
		t.Errorf("expected Minus as root, got %v", bin.operator.token_type)
	}
	// LHS should be (1-2)
	lhs, ok := bin.lhs.(*BinaryExpression)
	if !ok {
		t.Fatalf("expected BinaryExpression on lhs (left associativity), got %T", bin.lhs)
	}
	if lhs.operator.token_type != Minus {
		t.Errorf("expected Minus on lhs, got %v", lhs.operator.token_type)
	}
}

func TestParseDeeplyNested(t *testing.T) {
	_, err := Parse(tokens("(((((1)))))"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestParseZero(t *testing.T) {
	exp, err := Parse(tokens("0"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	num, ok := exp.(*NumberLiteral)
	if !ok {
		t.Fatalf("expected NumberLiteral, got %T", exp)
	}
	if num.literal != "0" {
		t.Errorf("expected '0', got '%s'", num.literal)
	}
}

func TestParseLargeNumber(t *testing.T) {
	exp, err := Parse(tokens("123456789"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	num, ok := exp.(*NumberLiteral)
	if !ok {
		t.Fatalf("expected NumberLiteral, got %T", exp)
	}
	if num.literal != "123456789" {
		t.Errorf("expected '123456789', got '%s'", num.literal)
	}
}

func TestParseMultipleParses(t *testing.T) {
	// Test that parser state resets correctly between parses
	_, err1 := Parse(tokens("1+2"))
	if err1 != nil {
		t.Fatalf("first parse failed: %v", err1)
	}

	_, err2 := Parse(tokens("3*4"))
	if err2 != nil {
		t.Fatalf("second parse failed: %v", err2)
	}

	// Error parse should not affect subsequent parses
	Parse(tokens("("))

	_, err3 := Parse(tokens("5-6"))
	if err3 != nil {
		t.Fatalf("third parse failed after error: %v", err3)
	}
}
