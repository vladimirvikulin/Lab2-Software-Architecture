package lab2

import (
	"math"
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
				return 0, InvalidExpressionError{}
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
			case "^":
				result = math.Pow(operand1, operand2)
			case "*":
				result = operand1 * operand2
			case "/":
				if operand2 == 0 {
					return 0, DivisionByZeroError{}
				}
				result = operand1 / operand2
			}
			stack = append(stack, result)
		default:
			operand, err := strconv.ParseFloat(token, 64)
			if err != nil {
				return 0, InvalidExpressionError{}
			}
			stack = append(stack, operand)
		}
	}
	if len(stack) != 1 {
		return 0, InvalidExpressionError{}
	}
	return stack[0], nil
}

type InvalidExpressionError struct{}

func (e InvalidExpressionError) Error() string {
	return "invalid expression"
}

type DivisionByZeroError struct{}

func (e DivisionByZeroError) Error() string {
	return "division by zero"
}
