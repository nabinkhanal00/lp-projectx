package calculate

import (
	"fmt"
	"unicode"
)

type TokenType int

const (
	NUMBER TokenType = iota
	PLUS
	MINUS
	PLUS_UNARY
	MINUS_UNARY
	MULTIPLY
	DIVIDE
	LPAREN
	RPAREN
	ILLEGAL
)

func (t TokenType) isOperator() bool {
	switch t {
	case PLUS:
		return true
	case MINUS:
		return true
	case MULTIPLY:
		return true
	case DIVIDE:
		return true
	case PLUS_UNARY:
		return true
	case MINUS_UNARY:
		return true
	default:
		return false
	}
}

func (t TokenType) toString() string {
	switch int(t) {
	case int(NUMBER):
		return "NUMBER"
	case int(PLUS):
		return "PLUS"
	case int(MINUS):
		return "MINUS"
	case int(PLUS_UNARY):
		return "PLUS_UNARY"
	case int(MINUS_UNARY):
		return "MINUS_UNARY"
	case int(MULTIPLY):
		return "MULTIPLY"
	case int(DIVIDE):
		return "DIVIDE"
	case int(LPAREN):
		return "LPAREN"
	case int(RPAREN):
		return "RPAREN"
	case int(ILLEGAL):
		return "ILLEGAL"
	}
	return "UNKNOWN"
}

type Token struct {
	Type     TokenType
	Value    string
	Position int
}

func (t Token) String() string {
	return fmt.Sprintf("TOKEN(TYPE: %s VALUE: %s POSIITON: %d)", t.Type.toString(), t.Value, t.Position)
}

func tokenize(expr string) []Token {
	var tokens []Token
	index := 0
	runeArray := []rune(expr)
	previous_paren := false
	for index < len(runeArray) {
		char := runeArray[index]
		switch {
		case unicode.IsDigit(char) || char == '.':
			var currentNumber string
			foundIllegal := false
			foundDot := false
			foundOther := false
			for index < len(runeArray) {
				char = runeArray[index]
				if unicode.IsDigit(char) || char == '.' {
					currentNumber += string(char)
					if char == '.' {
						if foundDot {
							foundIllegal = true
						}
						foundDot = true
					}
					index++
				} else {
					foundOther = true
					break
				}
			}
			if foundOther {
				index--
			}
			if foundIllegal {
				tokens = append(tokens, Token{Type: ILLEGAL, Value: currentNumber, Position: index})
			} else {
				tokens = append(tokens, Token{Type: NUMBER, Value: currentNumber, Position: index})
			}
		case char == ' ':
			break
		case char == '+':
			if previous_paren || index == 0 {
				tokens = append(tokens, Token{Type: PLUS_UNARY, Value: string(char), Position: index})
			} else {
				tokens = append(tokens, Token{Type: PLUS, Value: string(char), Position: index})
			}
		case char == '-':
			if previous_paren || index == 0 {
				tokens = append(tokens, Token{Type: MINUS_UNARY, Value: string(char), Position: index})
			} else {

				tokens = append(tokens, Token{Type: MINUS, Value: string(char), Position: index})
			}
		case char == '*':
			tokens = append(tokens, Token{Type: MULTIPLY, Value: string(char), Position: index})
		case char == '/':
			tokens = append(tokens, Token{Type: DIVIDE, Value: string(char), Position: index})
		case char == '(':
			tokens = append(tokens, Token{Type: LPAREN, Value: string(char), Position: index})
		case char == ')':
			tokens = append(tokens, Token{Type: RPAREN, Value: string(char), Position: index})
		default:
			tokens = append(tokens, Token{Type: ILLEGAL, Value: string(char), Position: index})
		}
		if char == '(' {
			previous_paren = true
		} else {
			previous_paren = false
		}
		index++
	}
	correctedTokens := []Token{}
	var prevToken *Token
	for _, token := range tokens {
		if prevToken != nil && token.Type == LPAREN && prevToken.Type == NUMBER {
			correctedTokens = append(correctedTokens, Token{
				Type:     MULTIPLY,
				Value:    "*",
				Position: -1,
			})
		}
		correctedTokens = append(correctedTokens, token)
		prevToken = &token
	}

	return correctedTokens
}
