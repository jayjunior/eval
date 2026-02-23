package internal

import (
	"testing"
)

func TestTokenizeSingleDigit(t *testing.T) {
	tokens, err := Tokenize("5")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(tokens) != 1 {
		t.Fatalf("expected 1 token, got %d", len(tokens))
	}
	if tokens[0].token_type != Number {
		t.Errorf("expected Number token, got %v", tokens[0].token_type)
	}
}

func TestTokenizeMultiDigitNumber(t *testing.T) {
	tokens, err := Tokenize("123")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(tokens) != 1 {
		t.Fatalf("expected 1 token, got %d", len(tokens))
	}
	if tokens[0].token != "123" {
		t.Errorf("expected '123', got '%s'", tokens[0].token)
	}
	if tokens[0].token_type != Number {
		t.Errorf("expected Number token, got %v", tokens[0].token_type)
	}
}

func TestTokenizeOperators(t *testing.T) {
	tests := []struct {
		input    string
		expected TokenType
	}{
		{"+", Plus},
		{"-", Minus},
		{"*", Multiplication},
		{"/", Division},
	}

	for _, tc := range tests {
		tokens, err := Tokenize(tc.input)
		if err != nil {
			t.Fatalf("unexpected error for '%s': %v", tc.input, err)
		}
		if len(tokens) != 1 {
			t.Fatalf("expected 1 token for '%s', got %d", tc.input, len(tokens))
		}
		if tokens[0].token_type != tc.expected {
			t.Errorf("for '%s': expected %v, got %v", tc.input, tc.expected, tokens[0].token_type)
		}
	}
}

func TestTokenizeParentheses(t *testing.T) {
	tokens, err := Tokenize("()")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(tokens) != 2 {
		t.Fatalf("expected 2 tokens, got %d", len(tokens))
	}
	if tokens[0].token_type != Open_Parentheses {
		t.Errorf("expected Open_Parentheses, got %v", tokens[0].token_type)
	}
	if tokens[1].token_type != Close_Parentheses {
		t.Errorf("expected Close_Parentheses, got %v", tokens[1].token_type)
	}
}

func TestTokenizeSimpleExpression(t *testing.T) {
	tokens, err := Tokenize("1+2")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(tokens) != 3 {
		t.Fatalf("expected 3 tokens, got %d", len(tokens))
	}
	if tokens[0].token_type != Number || tokens[0].token != "1" {
		t.Errorf("expected Number '1', got %v '%s'", tokens[0].token_type, tokens[0].token)
	}
	if tokens[1].token_type != Plus {
		t.Errorf("expected Plus, got %v", tokens[1].token_type)
	}
	if tokens[2].token_type != Number || tokens[2].token != "2" {
		t.Errorf("expected Number '2', got %v '%s'", tokens[2].token_type, tokens[2].token)
	}
}

func TestTokenizeComplexExpression(t *testing.T) {
	tokens, err := Tokenize("(5+3)*2")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := []struct {
		tokenType TokenType
		value     string
	}{
		{Open_Parentheses, "("},
		{Number, "5"},
		{Plus, "+"},
		{Number, "3"},
		{Close_Parentheses, ")"},
		{Multiplication, "*"},
		{Number, "2"},
	}
	if len(tokens) != len(expected) {
		t.Fatalf("expected %d tokens, got %d", len(expected), len(tokens))
	}
	for i, exp := range expected {
		if tokens[i].token_type != exp.tokenType {
			t.Errorf("token %d: expected type %v, got %v", i, exp.tokenType, tokens[i].token_type)
		}
	}
}

func TestTokenizeWithSpaces(t *testing.T) {
	tokens, err := Tokenize("1 + 2")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(tokens) != 3 {
		t.Fatalf("expected 3 tokens, got %d", len(tokens))
	}
	if tokens[0].token_type != Number {
		t.Errorf("expected Number, got %v", tokens[0].token_type)
	}
	if tokens[1].token_type != Plus {
		t.Errorf("expected Plus, got %v", tokens[1].token_type)
	}
	if tokens[2].token_type != Number {
		t.Errorf("expected Number, got %v", tokens[2].token_type)
	}
}

func TestTokenizeWithTabs(t *testing.T) {
	tokens, err := Tokenize("1\t+\t2")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(tokens) != 3 {
		t.Fatalf("expected 3 tokens, got %d", len(tokens))
	}
}

func TestTokenizeUnaryMinus(t *testing.T) {
	tokens, err := Tokenize("-5")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(tokens) != 2 {
		t.Fatalf("expected 2 tokens, got %d", len(tokens))
	}
	if tokens[0].token_type != Minus {
		t.Errorf("expected Minus, got %v", tokens[0].token_type)
	}
	if tokens[1].token_type != Number {
		t.Errorf("expected Number, got %v", tokens[1].token_type)
	}
}

func TestTokenizeDoubleUnaryMinus(t *testing.T) {
	tokens, err := Tokenize("--5")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(tokens) != 3 {
		t.Fatalf("expected 3 tokens, got %d", len(tokens))
	}
	if tokens[0].token_type != Minus {
		t.Errorf("expected Minus, got %v", tokens[0].token_type)
	}
	if tokens[1].token_type != Minus {
		t.Errorf("expected Minus, got %v", tokens[1].token_type)
	}
	if tokens[2].token_type != Number {
		t.Errorf("expected Number, got %v", tokens[2].token_type)
	}
}

func TestTokenizeNestedParentheses(t *testing.T) {
	tokens, err := Tokenize("((1+2))")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := []TokenType{
		Open_Parentheses,
		Open_Parentheses,
		Number,
		Plus,
		Number,
		Close_Parentheses,
		Close_Parentheses,
	}
	if len(tokens) != len(expected) {
		t.Fatalf("expected %d tokens, got %d", len(expected), len(tokens))
	}
	for i, exp := range expected {
		if tokens[i].token_type != exp {
			t.Errorf("token %d: expected %v, got %v", i, exp, tokens[i].token_type)
		}
	}
}

func TestTokenizeDivision(t *testing.T) {
	tokens, err := Tokenize("10/2")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(tokens) != 3 {
		t.Fatalf("expected 3 tokens, got %d", len(tokens))
	}
	if tokens[1].token_type != Division {
		t.Errorf("expected Division, got %v", tokens[1].token_type)
	}
}

func TestTokenizeAllOperators(t *testing.T) {
	tokens, err := Tokenize("1+2-3*4/5")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expectedTypes := []TokenType{
		Number, Plus, Number, Minus, Number, Multiplication, Number, Division, Number,
	}
	if len(tokens) != len(expectedTypes) {
		t.Fatalf("expected %d tokens, got %d", len(expectedTypes), len(tokens))
	}
	for i, exp := range expectedTypes {
		if tokens[i].token_type != exp {
			t.Errorf("token %d: expected %v, got %v", i, exp, tokens[i].token_type)
		}
	}
}

// Edge Cases

func TestTokenizeEmptyInput(t *testing.T) {
	tokens, err := Tokenize("")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(tokens) != 0 {
		t.Errorf("expected 0 tokens, got %d", len(tokens))
	}
}

func TestTokenizeOnlySpaces(t *testing.T) {
	tokens, err := Tokenize("   ")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(tokens) != 0 {
		t.Errorf("expected 0 tokens, got %d", len(tokens))
	}
}

func TestTokenizeLargeNumber(t *testing.T) {
	tokens, err := Tokenize("123456789")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(tokens) != 1 {
		t.Fatalf("expected 1 token, got %d", len(tokens))
	}
	if tokens[0].token != "123456789" {
		t.Errorf("expected '123456789', got '%s'", tokens[0].token)
	}
}

func TestTokenizeExpressionWithLeadingSpaces(t *testing.T) {
	tokens, err := Tokenize("  1+2")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(tokens) != 3 {
		t.Fatalf("expected 3 tokens, got %d", len(tokens))
	}
}

func TestTokenizeExpressionWithTrailingSpaces(t *testing.T) {
	tokens, err := Tokenize("1+2  ")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(tokens) != 3 {
		t.Fatalf("expected 3 tokens, got %d", len(tokens))
	}
}

func TestTokenizeZero(t *testing.T) {
	tokens, err := Tokenize("0")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(tokens) != 1 {
		t.Fatalf("expected 1 token, got %d", len(tokens))
	}
	if tokens[0].token != "0" {
		t.Errorf("expected '0', got '%s'", tokens[0].token)
	}
}

func TestTokenizeMultipleZeros(t *testing.T) {
	tokens, err := Tokenize("000")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(tokens) != 1 {
		t.Fatalf("expected 1 token, got %d", len(tokens))
	}
	if tokens[0].token != "000" {
		t.Errorf("expected '000', got '%s'", tokens[0].token)
	}
}

func TestTokenizeConsecutiveOperators(t *testing.T) {
	// This is syntactically valid for tokenization: 5 + -3
	tokens, err := Tokenize("5+-3")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(tokens) != 4 {
		t.Fatalf("expected 4 tokens, got %d", len(tokens))
	}
	expectedTypes := []TokenType{Number, Plus, Minus, Number}
	for i, exp := range expectedTypes {
		if tokens[i].token_type != exp {
			t.Errorf("token %d: expected %v, got %v", i, exp, tokens[i].token_type)
		}
	}
}

// Error Cases

func TestTokenizeUnrecognizedCharacterLetter(t *testing.T) {
	_, err := Tokenize("abc")
	if err == nil {
		t.Error("expected error for unrecognized character, got nil")
	}
}

func TestTokenizeUnrecognizedCharacterInMiddle(t *testing.T) {
	_, err := Tokenize("1+a")
	if err == nil {
		t.Error("expected error for unrecognized character, got nil")
	}
}

func TestTokenizeUnrecognizedSpecialChar(t *testing.T) {
	_, err := Tokenize("1@2")
	if err == nil {
		t.Error("expected error for unrecognized character '@', got nil")
	}
}

func TestTokenizeUnrecognizedPercent(t *testing.T) {
	_, err := Tokenize("100%")
	if err == nil {
		t.Error("expected error for unrecognized character '%', got nil")
	}
}

func TestTokenizeUnrecognizedDot(t *testing.T) {
	// Decimal numbers are not supported by the grammar
	_, err := Tokenize("3.14")
	if err == nil {
		t.Error("expected error for decimal point, got nil")
	}
}

func TestTokenizeUnrecognizedBrackets(t *testing.T) {
	_, err := Tokenize("[1+2]")
	if err == nil {
		t.Error("expected error for unrecognized character '[', got nil")
	}
}

func TestTokenizeUnrecognizedCurlyBraces(t *testing.T) {
	_, err := Tokenize("{1+2}")
	if err == nil {
		t.Error("expected error for unrecognized character '{', got nil")
	}
}

func TestTokenizeUnrecognizedCaret(t *testing.T) {
	// Exponentiation is not in the grammar
	_, err := Tokenize("2^3")
	if err == nil {
		t.Error("expected error for unrecognized character '^', got nil")
	}
}

func TestTokenizeUnrecognizedEquals(t *testing.T) {
	_, err := Tokenize("x=5")
	if err == nil {
		t.Error("expected error for unrecognized character, got nil")
	}
}

// Grammar-based edge cases

func TestTokenizeExpressionFromGrammar(t *testing.T) {
	// Test expression matching the grammar rule: term ( ( "-" | "+" ) term )*
	tokens, err := Tokenize("1+2+3-4")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(tokens) != 7 {
		t.Fatalf("expected 7 tokens, got %d", len(tokens))
	}
}

func TestTokenizeFactorFromGrammar(t *testing.T) {
	// Test factor matching the grammar rule: unary ( ( "/" | "*" ) unary )*
	tokens, err := Tokenize("2*3/4*5")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(tokens) != 7 {
		t.Fatalf("expected 7 tokens, got %d", len(tokens))
	}
}

func TestTokenizeMixedPrecedence(t *testing.T) {
	// Test mixed addition/subtraction and multiplication/division
	tokens, err := Tokenize("1+2*3-4/5")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(tokens) != 9 {
		t.Fatalf("expected 9 tokens, got %d", len(tokens))
	}
}

func TestTokenizeParenthesizedUnary(t *testing.T) {
	// Test: -(-5)
	tokens, err := Tokenize("-(-5)")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expectedTypes := []TokenType{Minus, Open_Parentheses, Minus, Number, Close_Parentheses}
	if len(tokens) != len(expectedTypes) {
		t.Fatalf("expected %d tokens, got %d", len(expectedTypes), len(tokens))
	}
	for i, exp := range expectedTypes {
		if tokens[i].token_type != exp {
			t.Errorf("token %d: expected %v, got %v", i, exp, tokens[i].token_type)
		}
	}
}
