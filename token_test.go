package main

import (
	"testing"
)

func checkValidToken(i, j Token) bool {
	if i.Type == j.Type && i.Value == j.Value {
		return true
	}
	return false
}

func checkValidTokenWithPosition(i, j Token) bool {
	if i.Type == j.Type && i.Value == j.Value {
		return true
	}
	return false
}

func TestTokenIsOperator(t *testing.T) {
	if !PLUS.isOperator() {
		t.Error("PLUS should be recognized as an operator")
	}
	if !MINUS.isOperator() {
		t.Error("MINUS should be recognized as an operator")
	}
	if !MULTIPLY.isOperator() {
		t.Error("MULTIPLY should be recognized as an operator")
	}
	if !DIVIDE.isOperator() {
		t.Error("DIVIDE should be recognized as an operator")
	}
	if !PLUS_UNARY.isOperator() {
		t.Error("PLUS_UNARY should be recognized as an operator")
	}
	if !MINUS_UNARY.isOperator() {
		t.Error("MINUS_UNARY should be recognized as an operator")
	}
	if NUMBER.isOperator() {
		t.Error("NUMBER should not be recognized as an operator")
	}
	if LPAREN.isOperator() {
		t.Error("LPAREN should not be recognized as an operator")
	}
	if RPAREN.isOperator() {
		t.Error("RPAREN should not be recognized as an operator")
	}
	if ILLEGAL.isOperator() {
		t.Error("ILLEGAL should not be recognized as an operator")
	}
}

func TestTokenizer(t *testing.T) {

	type exprToToken struct {
		expression string
		tokens     []Token
		fun        func(Token, Token) bool
	}

	exprToTokens := map[string]exprToToken{
		"simpleNumber": {
			expression: "1",
			tokens: []Token{
				{
					Type:  NUMBER,
					Value: "1",
				},
			},
			fun: checkValidToken,
		},
		"addition": {
			expression: "2 + 3",
			tokens: []Token{
				{
					Type:  NUMBER,
					Value: "2",
				},
				{
					Type:  PLUS,
					Value: "+",
				},
				{
					Type:  NUMBER,
					Value: "3",
				},
			},
			fun: checkValidToken,
		},
		"subtraction": {
			expression: "5 - 2",
			tokens: []Token{
				{
					Type:  NUMBER,
					Value: "5",
				},
				{
					Type:  MINUS,
					Value: "-",
				},
				{
					Type:  NUMBER,
					Value: "2",
				},
			},
			fun: checkValidToken,
		},
		"multiplication": {
			expression: "4 * 6",
			tokens: []Token{
				{
					Type:  NUMBER,
					Value: "4",
				},
				{
					Type:  MULTIPLY,
					Value: "*",
				},
				{
					Type:  NUMBER,
					Value: "6",
				},
			},
			fun: checkValidToken,
		},
		"division": {
			expression: "10 / 2",
			tokens: []Token{
				{
					Type:  NUMBER,
					Value: "10",
				},
				{
					Type:  DIVIDE,
					Value: "/",
				},
				{
					Type:  NUMBER,
					Value: "2",
				},
			},
			fun: checkValidToken,
		},
		"expression1": {
			expression: "(5 + 3) * (10 - 2) / 4",
			tokens: []Token{
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
			},
			fun: checkValidToken,
		},
		"expression2": {
			expression: "2.5 * (3 + 1.5) / (7 - 3.2)",
			tokens: []Token{
				{Type: NUMBER, Value: "2.5"},
				{Type: MULTIPLY, Value: "*"},
				{Type: LPAREN, Value: "("},
				{Type: NUMBER, Value: "3"},
				{Type: PLUS, Value: "+"},
				{Type: NUMBER, Value: "1.5"},
				{Type: RPAREN, Value: ")"},
				{Type: DIVIDE, Value: "/"},
				{Type: LPAREN, Value: "("},
				{Type: NUMBER, Value: "7"},
				{Type: MINUS, Value: "-"},
				{Type: NUMBER, Value: "3.2"},
				{Type: RPAREN, Value: ")"},
			},
			fun: checkValidToken,
		},
	}

	for name, exptt := range exprToTokens {
		outputTokens := tokenize(exptt.expression)
		if len(outputTokens) != len(exptt.tokens) {
			t.Errorf("Test for %v failed:\n different length of tokens.", name)
		}
		for i := 0; i < len(outputTokens); i++ {
			if !exptt.fun(outputTokens[i], exptt.tokens[i]) {
				t.Errorf("Test for %v failed:\n Expected %v\n Got %v", name, exptt.tokens[i], outputTokens[i])
			}
		}

	}

}
