package main

import (
	"reflect"
	"testing"
)

func TestRPN(t *testing.T) {
	tokens := []Token{
		{Type: LPAREN, Value: "("},
		{Type: NUMBER, Value: "5"},
		{Type: PLUS, Value: "+"},
		{Type: NUMBER, Value: "3"},
		{Type: RPAREN, Value: ")"},
		{Type: MULTIPLY, Value: "*"},
		{Type: LPAREN, Value: "("},
		{Type: NUMBER, Value: "10"},
		{Type: MINUS, Value: "-"},
		{Type: NUMBER, Value: "2"},
		{Type: RPAREN, Value: ")"},
		{Type: DIVIDE, Value: "/"},
		{Type: NUMBER, Value: "4"},
	}
	expectedPostfix := []*Token{
		{Type: NUMBER, Value: "5"},
		{Type: NUMBER, Value: "3"},
		{Type: PLUS, Value: "+"},
		{Type: NUMBER, Value: "10"},
		{Type: NUMBER, Value: "2"},
		{Type: MINUS, Value: "-"},
		{Type: MULTIPLY, Value: "*"},
		{Type: NUMBER, Value: "4"},
		{Type: DIVIDE, Value: "/"},
	}
	result, err := rpn(tokens)
	if err != nil {
		t.Errorf("RPN failed with error: %v", err)
	}
	if !reflect.DeepEqual(result, expectedPostfix) {
		t.Errorf("RPN failed. Expected: %v, Got: %v", expectedPostfix, result)
	}
}

func TestIsTokenizationValid(t *testing.T) {
	validTokens := []Token{
		{Type: LPAREN, Value: "("},
		{Type: NUMBER, Value: "5"},
		{Type: PLUS, Value: "+"},
		{Type: NUMBER, Value: "3"},
		{Type: RPAREN, Value: ")"},
	}
	if err := isTokenizationValid(validTokens); err != nil {
		t.Errorf("Tokenization validation failed. Expected valid tokens, got error: %v", err)
	}

	invalidTokens := []Token{
		{Type: LPAREN, Value: "("},
		{Type: NUMBER, Value: "5"},
		{Type: PLUS, Value: "+"},
		{Type: NUMBER, Value: "3"},
		{Type: RPAREN, Value: ")"},
		{Type: RPAREN, Value: ")"},
	}
	if err := isTokenizationValid(invalidTokens); err == nil {
		t.Errorf("Tokenization validation failed. Expected invalid tokens, got no error")
	}
}
func TestEvaluate(t *testing.T) {
	postfix := []*Token{
		{Type: NUMBER, Value: "5"},
		{Type: NUMBER, Value: "3"},
		{Type: PLUS, Value: "+"},
		{Type: NUMBER, Value: "10"},
		{Type: NUMBER, Value: "2"},
		{Type: MINUS, Value: "-"},
		{Type: MULTIPLY, Value: "*"},
		{Type: NUMBER, Value: "4"},
		{Type: DIVIDE, Value: "/"},
	}
	expectedResult := "16"
	result, _, err := evaluate(postfix)
	if err != nil {
		t.Errorf("Evaluate failed with error: %v", err)
	}
	if result != expectedResult {
		t.Errorf("Evaluate failed. Expected result: %s, Got: %s", expectedResult, result)
	}
}
