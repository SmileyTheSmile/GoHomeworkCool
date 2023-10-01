package parsing

import (
	datastructs "GoHomework/datastructs"
	"errors"
	"math"
	"strconv"
)

var BinaryOperations = map[string]func(float64, float64) (float64, error){
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

var UnaryOperations = map[string]func(float64) float64{
	"unary_minus": func(operand float64) float64 {
		return -operand
	},
}

func SolvePostfix(postfixExpression []string) (float64, error) {
	var operandsStack datastructs.Stack[float64]
	for _, token := range postfixExpression {
		switch {
		case BinaryOperations[token] != nil:
			right, _ := operandsStack.Pop()
			left, _ := operandsStack.Pop()
			result, err := BinaryOperations[token](left, right)
			if err != nil {
				return 0, err
			}
			operandsStack.Push(result)
		case UnaryOperations[token] != nil:
			expression, _ := operandsStack.Pop()
			result := UnaryOperations[token](expression)
			operandsStack.Push(result)
		default:
			operand, _ := strconv.ParseFloat(token, 64)
			operandsStack.Push(operand)
		}
	}
	result, _ := operandsStack.Pop()
	return result, nil
}
