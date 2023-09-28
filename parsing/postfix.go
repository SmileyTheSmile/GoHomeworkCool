package parsing

import (
	"errors"
	"math"
	"strconv"
)

var BINARY_OPERATIONS = map[string]func(float64, float64) (float64, error){
	"+": func(left float64, right float64) (float64, error) {
		return left + right, nil
	},
	"-": func(left float64, right float64) (float64, error) {
		return left - right, nil
	},
	"*": func(left float64, right float64) (float64, error) {
		return left * right, nil
	},
	"^": func(left float64, right float64) (float64, error) {
		return math.Pow(left, right), nil
	},
	"/": func(left float64, right float64) (float64, error) {
		if right == 0 {
			return 0, errors.New("Деление на 0")
		}
		return left / right, nil
	},
}

var UNARY_OPERATIONS = map[string]func(float64) float64{
	"ceil": func(expression float64) float64 {
		return math.Ceil(expression)
	},
	"unary_minus": func(expression float64) float64 {
		return -expression
	},
	"sqrt": func(expression float64) float64 {
		return math.Sqrt(expression)
	},
}

func SolvePostfix(postfixExpression []string) (float64, error) {
	var operandsStack Stack[float64]
	for _, token := range postfixExpression {
		switch {
		case BINARY_OPERATIONS[token] != nil:
			right, _ := operandsStack.Pop()
			left, _ := operandsStack.Pop()
			result, err := BINARY_OPERATIONS[token](left, right)
			//fmt.Println(left, token, right, "=", result)
			if err != nil {
				return 0, err
			}
			operandsStack.Push(result)
		case UNARY_OPERATIONS[token] != nil:
			expression, _ := operandsStack.Pop()
			result := UNARY_OPERATIONS[token](expression)
			//fmt.Println(token, expression, "=", result)
			operandsStack.Push(result)
		default:
			operand, _ := strconv.ParseFloat(token, 64)
			operandsStack.Push(operand)
		}
	}
	result, _ := operandsStack.Pop()
	return result, nil
}
