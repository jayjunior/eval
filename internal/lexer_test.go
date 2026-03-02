package internal_test

import (
	"testing"

	"github.com/jayjunior/eval/internal"
	"github.com/jayjunior/eval/internal/ast"
)

func TestTokenizeSingleDigit(t *testing.T) {
	tokens, err := internal.Tokenize("5")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(tokens) != 1 {
		t.Fatalf("expected 1 token, got %d", len(tokens))
	}
	if tokens[0].Token != ast.NUMBER {
		t.Errorf("expected Number token, got %v", tokens[0].Token)
	}
}

func TestTokenizeMultiDigitNumber(t *testing.T) {
	tokens, err := internal.Tokenize("123")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(tokens) != 1 {
		t.Fatalf("expected 1 token, got %d", len(tokens))
	}
	if tokens[0].Literal != "123" {
		t.Errorf("expected '123', got '%s'", tokens[0].Literal)
	}
	if tokens[0].Token != ast.NUMBER {
		t.Errorf("expected Number token, got %v", tokens[0].Token)
	}
}

func TestTokenizeOperators(t *testing.T) {
	tests := []struct {
		input    string
		expected ast.TokenType
	}{
		{"+", ast.Plus},
		{"-", ast.Minus},
		{"*", ast.Multiplication},
		{"/", ast.Division},
	}

	for _, tc := range tests {
		tokens, err := internal.Tokenize(tc.input)
		if err != nil {
			t.Fatalf("unexpected error for '%s': %v", tc.input, err)
		}
		if len(tokens) != 1 {
			t.Fatalf("expected 1 token for '%s', got %d", tc.input, len(tokens))
		}
		if tokens[0].Token != tc.expected {
			t.Errorf("for '%s': expected %v, got %v", tc.input, tc.expected, tokens[0].Token)
		}
	}
}

func TestTokenizeParentheses(t *testing.T) {
	tokens, err := internal.Tokenize("()")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(tokens) != 2 {
		t.Fatalf("expected 2 tokens, got %d", len(tokens))
	}
	if tokens[0].Token != ast.Open_Parentheses {
		t.Errorf("expected Open_Parentheses, got %v", tokens[0].Token)
	}
	if tokens[1].Token != ast.Close_Parentheses {
		t.Errorf("expected Close_Parentheses, got %v", tokens[1].Token)
	}
}

func TestTokenizeSimpleExpression(t *testing.T) {
	tokens, err := internal.Tokenize("1+2")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(tokens) != 3 {
		t.Fatalf("expected 3 tokens, got %d", len(tokens))
	}
	if tokens[0].Token != ast.NUMBER || tokens[0].Literal != "1" {
		t.Errorf("expected Number '1', got %v '%s'", tokens[0].Token, tokens[0].Literal)
	}
	if tokens[1].Token != ast.Plus {
		t.Errorf("expected Plus, got %v", tokens[1].Token)
	}
	if tokens[2].Token != ast.NUMBER || tokens[2].Literal != "2" {
		t.Errorf("expected Number '2', got %v '%s'", tokens[2].Token, tokens[2].Literal)
	}
}

func TestTokenizeComplexExpression(t *testing.T) {
	tokens, err := internal.Tokenize("(5+3)*2")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := []struct {
		tokenType ast.TokenType
		value     string
	}{
		{ast.Open_Parentheses, "("},
		{ast.NUMBER, "5"},
		{ast.Plus, "+"},
		{ast.NUMBER, "3"},
		{ast.Close_Parentheses, ")"},
		{ast.Multiplication, "*"},
		{ast.NUMBER, "2"},
	}
	if len(tokens) != len(expected) {
		t.Fatalf("expected %d tokens, got %d", len(expected), len(tokens))
	}
	for i, exp := range expected {
		if tokens[i].Token != exp.tokenType {
			t.Errorf("token %d: expected type %v, got %v", i, exp.tokenType, tokens[i].Token)
		}
	}
}

func TestTokenizeWithSpaces(t *testing.T) {
	tokens, err := internal.Tokenize("1 + 2")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(tokens) != 3 {
		t.Fatalf("expected 3 tokens, got %d", len(tokens))
	}
	if tokens[0].Token != ast.NUMBER {
		t.Errorf("expected Number, got %v", tokens[0].Token)
	}
	if tokens[1].Token != ast.Plus {
		t.Errorf("expected Plus, got %v", tokens[1].Token)
	}
	if tokens[2].Token != ast.NUMBER {
		t.Errorf("expected Number, got %v", tokens[2].Token)
	}
}

func TestTokenizeWithTabs(t *testing.T) {
	tokens, err := internal.Tokenize("1\t+\t2")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(tokens) != 3 {
		t.Fatalf("expected 3 tokens, got %d", len(tokens))
	}
}

func TestTokenizeUnaryMinus(t *testing.T) {
	tokens, err := internal.Tokenize("-5")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(tokens) != 2 {
		t.Fatalf("expected 2 tokens, got %d", len(tokens))
	}
	if tokens[0].Token != ast.Minus {
		t.Errorf("expected Minus, got %v", tokens[0].Token)
	}
	if tokens[1].Token != ast.NUMBER {
		t.Errorf("expected Number, got %v", tokens[1].Token)
	}
}

func TestTokenizeDoubleUnaryMinus(t *testing.T) {
	tokens, err := internal.Tokenize("--5")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(tokens) != 3 {
		t.Fatalf("expected 3 tokens, got %d", len(tokens))
	}
	if tokens[0].Token != ast.Minus {
		t.Errorf("expected Minus, got %v", tokens[0].Token)
	}
	if tokens[1].Token != ast.Minus {
		t.Errorf("expected Minus, got %v", tokens[1].Token)
	}
	if tokens[2].Token != ast.NUMBER {
		t.Errorf("expected Number, got %v", tokens[2].Token)
	}
}

func TestTokenizeNestedParentheses(t *testing.T) {
	tokens, err := internal.Tokenize("((1+2))")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := []ast.TokenType{
		ast.Open_Parentheses,
		ast.Open_Parentheses,
		ast.NUMBER,
		ast.Plus,
		ast.NUMBER,
		ast.Close_Parentheses,
		ast.Close_Parentheses,
	}
	if len(tokens) != len(expected) {
		t.Fatalf("expected %d tokens, got %d", len(expected), len(tokens))
	}
	for i, exp := range expected {
		if tokens[i].Token != exp {
			t.Errorf("token %d: expected %v, got %v", i, exp, tokens[i].Token)
		}
	}
}

func TestTokenizeDivision(t *testing.T) {
	tokens, err := internal.Tokenize("10/2")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(tokens) != 3 {
		t.Fatalf("expected 3 tokens, got %d", len(tokens))
	}
	if tokens[1].Token != ast.Division {
		t.Errorf("expected Division, got %v", tokens[1].Token)
	}
}

func TestTokenizeAllOperators(t *testing.T) {
	tokens, err := internal.Tokenize("1+2-3*4/5")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expectedTypes := []ast.TokenType{
		ast.NUMBER, ast.Plus, ast.NUMBER, ast.Minus, ast.NUMBER, ast.Multiplication, ast.NUMBER, ast.Division, ast.NUMBER,
	}
	if len(tokens) != len(expectedTypes) {
		t.Fatalf("expected %d tokens, got %d", len(expectedTypes), len(tokens))
	}
	for i, exp := range expectedTypes {
		if tokens[i].Token != exp {
			t.Errorf("token %d: expected %v, got %v", i, exp, tokens[i].Token)
		}
	}
}

// Edge Cases

func TestTokenizeEmptyInput(t *testing.T) {
	tokens, err := internal.Tokenize("")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(tokens) != 0 {
		t.Errorf("expected 0 tokens, got %d", len(tokens))
	}
}

func TestTokenizeOnlySpaces(t *testing.T) {
	tokens, err := internal.Tokenize("   ")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(tokens) != 0 {
		t.Errorf("expected 0 tokens, got %d", len(tokens))
	}
}

func TestTokenizeLargeNumber(t *testing.T) {
	tokens, err := internal.Tokenize("123456789")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(tokens) != 1 {
		t.Fatalf("expected 1 token, got %d", len(tokens))
	}
	if tokens[0].Literal != "123456789" {
		t.Errorf("expected '123456789', got '%s'", tokens[0].Literal)
	}
}

func TestTokenizeExpressionWithLeadingSpaces(t *testing.T) {
	tokens, err := internal.Tokenize("  1+2")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(tokens) != 3 {
		t.Fatalf("expected 3 tokens, got %d", len(tokens))
	}
}

func TestTokenizeExpressionWithTrailingSpaces(t *testing.T) {
	tokens, err := internal.Tokenize("1+2  ")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(tokens) != 3 {
		t.Fatalf("expected 3 tokens, got %d", len(tokens))
	}
}

func TestTokenizeZero(t *testing.T) {
	tokens, err := internal.Tokenize("0")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(tokens) != 1 {
		t.Fatalf("expected 1 token, got %d", len(tokens))
	}
	if tokens[0].Literal != "0" {
		t.Errorf("expected '0', got '%s'", tokens[0].Literal)
	}
}

func TestTokenizeMultipleZeros(t *testing.T) {
	tokens, err := internal.Tokenize("000")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(tokens) != 1 {
		t.Fatalf("expected 1 token, got %d", len(tokens))
	}
	if tokens[0].Literal != "000" {
		t.Errorf("expected '000', got '%s'", tokens[0].Literal)
	}
}

func TestTokenizeConsecutiveOperators(t *testing.T) {
	// This is syntactically valid for tokenization: 5 + -3
	tokens, err := internal.Tokenize("5+-3")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(tokens) != 4 {
		t.Fatalf("expected 4 tokens, got %d", len(tokens))
	}
	expectedTypes := []ast.TokenType{ast.NUMBER, ast.Plus, ast.Minus, ast.NUMBER}
	for i, exp := range expectedTypes {
		if tokens[i].Token != exp {
			t.Errorf("token %d: expected %v, got %v", i, exp, tokens[i].Token)
		}
	}
}

// Error Cases




func TestTokenizeUnrecognizedSpecialChar(t *testing.T) {
	_, err := internal.Tokenize("1@2")
	if err == nil {
		t.Error("expected error for unrecognized character '@', got nil")
	}
}

func TestTokenizeUnrecognizedPercent(t *testing.T) {
	_, err := internal.Tokenize("100%")
	if err == nil {
		t.Error("expected error for unrecognized character '%', got nil")
	}
}

func TestTokenizeUnrecognizedDot(t *testing.T) {
	// Decimal numbers are not supported by the grammar
	_, err := internal.Tokenize("3.14")
	if err == nil {
		t.Error("expected error for decimal point, got nil")
	}
}

func TestTokenizeUnrecognizedBrackets(t *testing.T) {
	_, err := internal.Tokenize("[1+2]")
	if err == nil {
		t.Error("expected error for unrecognized character '[', got nil")
	}
}

func TestTokenizeUnrecognizedCurlyBraces(t *testing.T) {
	_, err := internal.Tokenize("{1+2}")
	if err == nil {
		t.Error("expected error for unrecognized character '{', got nil")
	}
}

func TestTokenizeUnrecognizedCaret(t *testing.T) {
	// Exponentiation is not in the grammar
	_, err := internal.Tokenize("2^3")
	if err == nil {
		t.Error("expected error for unrecognized character '^', got nil")
	}
}



// Grammar-based edge cases

func TestTokenizeExpressionFromGrammar(t *testing.T) {
	// Test expression matching the grammar rule: term ( ( "-" | "+" ) term )*
	tokens, err := internal.Tokenize("1+2+3-4")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(tokens) != 7 {
		t.Fatalf("expected 7 tokens, got %d", len(tokens))
	}
}

func TestTokenizeFactorFromGrammar(t *testing.T) {
	// Test factor matching the grammar rule: unary ( ( "/" | "*" ) unary )*
	tokens, err := internal.Tokenize("2*3/4*5")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(tokens) != 7 {
		t.Fatalf("expected 7 tokens, got %d", len(tokens))
	}
}

func TestTokenizeMixedPrecedence(t *testing.T) {
	// Test mixed addition/subtraction and multiplication/division
	tokens, err := internal.Tokenize("1+2*3-4/5")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(tokens) != 9 {
		t.Fatalf("expected 9 tokens, got %d", len(tokens))
	}
}

func TestTokenizeParenthesizedUnary(t *testing.T) {
	// Test: -(-5)
	tokens, err := internal.Tokenize("-(-5)")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expectedTypes := []ast.TokenType{ast.Minus, ast.Open_Parentheses, ast.Minus, ast.NUMBER, ast.Close_Parentheses}
	if len(tokens) != len(expectedTypes) {
		t.Fatalf("expected %d tokens, got %d", len(expectedTypes), len(tokens))
	}
	for i, exp := range expectedTypes {
		if tokens[i].Token != exp {
			t.Errorf("token %d: expected %v, got %v", i, exp, tokens[i].Token)
		}
	}
}

// Variable Declaration Tests

func TestTokenizeVarKeyword(t *testing.T) {
	tokens, err := internal.Tokenize("var")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(tokens) != 1 {
		t.Fatalf("expected 1 token, got %d", len(tokens))
	}
	if tokens[0].Token != ast.VAR {
		t.Errorf("expected VAR token, got %v", tokens[0].Token)
	}
	if tokens[0].Literal != "var" {
		t.Errorf("expected 'var', got '%s'", tokens[0].Literal)
	}
}

func TestTokenizeSimpleIdentifier(t *testing.T) {
	tokens, err := internal.Tokenize("x")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(tokens) != 1 {
		t.Fatalf("expected 1 token, got %d", len(tokens))
	}
	if tokens[0].Token != ast.IDENTIFIER {
		t.Errorf("expected IDENTIFIER token, got %v", tokens[0].Token)
	}
	if tokens[0].Literal != "x" {
		t.Errorf("expected 'x', got '%s'", tokens[0].Literal)
	}
}

func TestTokenizeMultiCharIdentifier(t *testing.T) {
	tokens, err := internal.Tokenize("myVar")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(tokens) != 1 {
		t.Fatalf("expected 1 token, got %d", len(tokens))
	}
	if tokens[0].Token != ast.IDENTIFIER {
		t.Errorf("expected IDENTIFIER token, got %v", tokens[0].Token)
	}
	if tokens[0].Literal != "myVar" {
		t.Errorf("expected 'myVar', got '%s'", tokens[0].Literal)
	}
}

func TestTokenizeEqualSign(t *testing.T) {
	tokens, err := internal.Tokenize("=")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(tokens) != 1 {
		t.Fatalf("expected 1 token, got %d", len(tokens))
	}
	if tokens[0].Token != ast.EQUAL {
		t.Errorf("expected EQUAL token, got %v", tokens[0].Token)
	}
}

func TestTokenizeVariableDeclaration(t *testing.T) {
	tokens, err := internal.Tokenize("var myVar")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(tokens) != 2 {
		t.Fatalf("expected 2 tokens, got %d", len(tokens))
	}
	if tokens[0].Token != ast.VAR {
		t.Errorf("token 0: expected VAR, got %v", tokens[0].Token)
	}
	if tokens[1].Token != ast.IDENTIFIER {
		t.Errorf("token 1: expected IDENTIFIER, got %v", tokens[1].Token)
	}
	if tokens[1].Literal != "myVar" {
		t.Errorf("token 1: expected 'myVar', got '%s'", tokens[1].Literal)
	}
}

func TestTokenizeSimpleAssignment(t *testing.T) {
	tokens, err := internal.Tokenize("x = 5")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(tokens) != 3 {
		t.Fatalf("expected 3 tokens, got %d", len(tokens))
	}
	expectedTypes := []ast.TokenType{ast.IDENTIFIER, ast.EQUAL, ast.NUMBER}
	for i, exp := range expectedTypes {
		if tokens[i].Token != exp {
			t.Errorf("token %d: expected %v, got %v", i, exp, tokens[i].Token)
		}
	}
}

func TestTokenizeAssignmentWithExpression(t *testing.T) {
	tokens, err := internal.Tokenize("x = 1 + 2")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(tokens) != 5 {
		t.Fatalf("expected 5 tokens, got %d", len(tokens))
	}
	expectedTypes := []ast.TokenType{ast.IDENTIFIER, ast.EQUAL, ast.NUMBER, ast.Plus, ast.NUMBER}
	for i, exp := range expectedTypes {
		if tokens[i].Token != exp {
			t.Errorf("token %d: expected %v, got %v", i, exp, tokens[i].Token)
		}
	}
}

func TestTokenizeAssignmentWithComplexExpression(t *testing.T) {
	tokens, err := internal.Tokenize("result = (10 + 5) * 2")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(tokens) != 9 {
		t.Fatalf("expected 9 tokens, got %d", len(tokens))
	}
	expectedTypes := []ast.TokenType{
		ast.IDENTIFIER, ast.EQUAL, ast.Open_Parentheses, ast.NUMBER, ast.Plus, ast.NUMBER,
		ast.Close_Parentheses, ast.Multiplication, ast.NUMBER,
	}
	for i, exp := range expectedTypes {
		if tokens[i].Token != exp {
			t.Errorf("token %d: expected %v, got %v", i, exp, tokens[i].Token)
		}
	}
}

func TestTokenizeMultipleStatements(t *testing.T) {
	// Test case should handle assignments with various identifiers
	tokens, err := internal.Tokenize("count = 10 + count")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(tokens) != 5 {
		t.Fatalf("expected 5 tokens, got %d", len(tokens))
	}
}

func TestTokenizeVariableDeclarationWithSpaces(t *testing.T) {
	tokens, err := internal.Tokenize("var  x")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(tokens) != 2 {
		t.Fatalf("expected 2 tokens, got %d", len(tokens))
	}
	if tokens[0].Token != ast.VAR || tokens[1].Token != ast.IDENTIFIER {
		t.Errorf("expected VAR followed by IDENTIFIER")
	}
}

func TestTokenizeAssignmentWithSpaces(t *testing.T) {
	tokens, err := internal.Tokenize("x  =  5")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(tokens) != 3 {
		t.Fatalf("expected 3 tokens, got %d", len(tokens))
	}
}

func TestTokenizeAssignmentWithIdentifierExpression(t *testing.T) {
	tokens, err := internal.Tokenize("x = y")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(tokens) != 3 {
		t.Fatalf("expected 3 tokens, got %d", len(tokens))
	}
	expectedTypes := []ast.TokenType{ast.IDENTIFIER, ast.EQUAL, ast.IDENTIFIER}
	for i, exp := range expectedTypes {
		if tokens[i].Token != exp {
			t.Errorf("token %d: expected %v, got %v", i, exp, tokens[i].Token)
		}
	}
}
