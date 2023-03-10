package calculator

import (
	"fmt"
	"strconv"
	"strings"
)

func Evaluate(expression string) (float64, error) {
	tokens := strings.Fields(expression)

	stack := make([]float64, 0, len(tokens))
	for _, token := range tokens {
		switch token {
		case "+", "-", "*", "/", "^":
			if len(stack) < 2 {
				return 0, fmt.Errorf("invalid expression")
			}
			var result float64
			operand2 := stack[len(stack)-1]
			operand1 := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			switch token {
			case "+":
				result = operand1 + operand2
			case "-":
				result = operand1 - operand2
			case "*":
				result = operand1 * operand2
			case "/":
				if operand2 == 0 {
					return 0, fmt.Errorf("division by zero")
				}
				result = operand1 / operand2
			}
			stack = append(stack, result)
		default:
			operand, err := strconv.ParseFloat(token, 64)
			if err != nil {
				return 0, fmt.Errorf("invalid expression")
			}
			stack = append(stack, operand)
		}
	}
	return stack[0], nil
}
