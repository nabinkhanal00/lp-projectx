package main

import (
	"fmt"
	"strconv"
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

func rpn(tokens []Token) ([]*Token, error) {
	// final result
	postfix := []*Token{}

	// would be better if we implement a real stack
	// making it efficient
	// stack for storing operators
	stack := []*Token{}

	current := 0
	for current < len(tokens) {
		// get the pointer to the current token
		token := &tokens[current]
		if token.Type == NUMBER {
			// push it in the result if it is a operand
			postfix = append(postfix, token)
		} else if token.Type == LPAREN {
			// push it in the staci of operators
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

			// append to stack if stack is empty or
			// remove elements from the stack until
			// the precedence of current token is less
			// or equal to top of stack
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

	// pop all the elements from t
	for len(stack) > 0 {
		postfix = append(postfix, stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}
	return postfix, nil
}

func isTokenizationValid(tokens []Token) error {
	parenCount := 0
	var prevToken *Token
	for _, token := range tokens {
		if token.Type == LPAREN {
			parenCount++
		} else if token.Type == RPAREN {
			parenCount--
			if parenCount < 0 {
				return fmt.Errorf("Parentheses donot match: extra ) at position %d", token.Position)

			}
		}
		if prevToken != nil && (prevToken.Type.isOperator() && token.Type.isOperator()) {

			return fmt.Errorf("Invalid token %v at position %d.", token.Value, token.Position)
		}
		if token.Type == ILLEGAL {
			return fmt.Errorf("Illegal token %v at position %d.", token.Value, token.Position)
		}
		prevToken = &token
	}
	return nil
}

func evaluate(postfix []*Token) (string, []string, error) {
	// push operands on the stack
	// if an operator is found
	// pop one or two operands based on the type of operator
	// perform the operation
	// push result on the stack
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
	result := formatFloat(stack[len(stack)-1], 15)
	return fmt.Sprintf("%v", result), steps, nil
}

func Calculate(expression string) (string, []string, error) {
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
