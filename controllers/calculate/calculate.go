package calculate

import (
	"fmt"
	"strconv"
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

var priorityMap map[TokenType]int = map[TokenType]int{
	LPAREN:      0,
	PLUS:        1,
	MINUS:       1,
	MULTIPLY:    2,
	DIVIDE:      2,
	PLUS_UNARY:  3,
	MINUS_UNARY: 3,
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
func rpn(tokens []Token) ([]*Token, error) {
	postfix := []*Token{}

	// would be better if we implement a real stack
	// making it efficient
	stack := []*Token{}

	current := 0
	for current < len(tokens) {
		token := &tokens[current]
		if token.Type == NUMBER {
			postfix = append(postfix, token)
		} else if token.Type == LPAREN {
			stack = append(stack, token)
		} else if token.Type == RPAREN {
			// stack can never be empty here
			// we have checked the validity of parentheses
			// so it is safe to get the top of stack
			top := stack[len(stack)-1]
			prevToken := tokens[current-1]

			// we are checking for expression like '()'
			// continuous parentheses are invalid
			if prevToken.Type == LPAREN {
				return nil, fmt.Errorf("Invalid ) at position %d", top.Position)
			}

			// we donot need to check the emptyness of the stack
			// since matching parenthesis must be present
			for stack[len(stack)-1].Type != LPAREN {
				postfix = append(postfix, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}

			// removing LPAREN from the stack
			stack = stack[:len(stack)-1]
		} else if token.Type == ILLEGAL {
			return nil, fmt.Errorf("Illegal token %v at position %d.", token.Value, token.Position)
		} else {
			// for all the other operators

			if len(stack) == 0 || priorityMap[token.Type] > priorityMap[stack[len(stack)-1].Type] {
				stack = append(stack, token)
			} else {
				for len(stack) > 0 && priorityMap[token.Type] <= priorityMap[stack[len(stack)-1].Type] {
					top := stack[len(stack)-1]
					postfix = append(postfix, top)
					stack = stack[:len(stack)-1]
				}

				stack = append(stack, token)
			}
		}
		current++
	}

	for len(stack) > 0 {
		postfix = append(postfix, stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}
	return postfix, nil
}

func isTokenizationValid(tokens []Token) error {
	parenCount := 0

	for _, token := range tokens {
		if token.Type == LPAREN {
			parenCount++
		} else if token.Type == RPAREN {
			parenCount--
			if parenCount < 0 {
				return fmt.Errorf("Parentheses donot match: extra ) at position %d", token.Position)

			}
		}
		if token.Type == ILLEGAL {
			return fmt.Errorf("Illegal token %v at position %d.", token.Value, token.Position)
		}
	}
	return nil
}

func evaluate(postfix []*Token) (string, []string, error) {
	steps := []string{}
	stack := []float64{}
	for _, token := range postfix {
		if token.Type == NUMBER {

			value, err := strconv.ParseFloat(token.Value, 64)
			if err != nil {
				return "", nil, err
			}
			stack = append(stack, value)
		} else {
			if token.Type == PLUS_UNARY || token.Type == MINUS_UNARY {
				if len(stack) < 1 {
					return "", steps, fmt.Errorf("Invalid %s at position %d", token.Value, token.Position)
				}
				top := stack[len(stack)-1]
				var result float64 = top
				if token.Type == MINUS_UNARY {
					result = -1 * top
				}
				stack[len(stack)-1] = result
				step := fmt.Sprintf("%s%v = %v", token.Value, formatFloat(top, 15), formatFloat(result, 15))
				steps = append(steps, step)
			} else {
				if len(stack) < 2 {
					return "", steps, fmt.Errorf("Invalid %s at position %d", token.Value, token.Position)
				}
				valueFirst := stack[len(stack)-2]
				valueSecond := stack[len(stack)-1]
				stack = stack[:len(stack)-2]

				var result float64
				var step string
				switch token.Type {
				case PLUS:
					result = valueFirst + valueSecond
				case MINUS:
					result = valueFirst - valueSecond
				case MULTIPLY:
					result = valueFirst * valueSecond
				case DIVIDE:
					result = valueFirst / valueSecond
				default:
					return "", steps, fmt.Errorf("Unknown operator: %s at position %d", token.Value, token.Position)
				}

				step = fmt.Sprintf("%v %v %v = %v", formatFloat(valueFirst, 15), token.Value, formatFloat(valueSecond, 15), formatFloat(result, 15))
				steps = append(steps, step)
				stack = append(stack, result)
			}
		}
	}
	if len(stack) != 1 {
		return "", steps, fmt.Errorf("Unknown error")
	}
	result := strconv.FormatFloat(stack[len(stack)-1], 'g', 15, 64)
	return fmt.Sprintf("%v", result), steps, nil
}

func formatFloat(value float64, precision int) string {
	return strconv.FormatFloat(value, 'g', precision, 64)
}

func calculate(expression string) (string, []string, error) {
	tokens := tokenize(expression)
	err := isTokenizationValid(tokens)
	if err != nil {
		return "", nil, err
	}
	postfix, err := rpn(tokens)
	if err != nil {
		return "", nil, err
	}
	result, steps, err := evaluate(postfix)

	if err != nil {
		return "", nil, err
	}

	return result, steps, nil
}
