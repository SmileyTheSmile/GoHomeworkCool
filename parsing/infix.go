package parsing

import (
	datastructs "GoHomework/datastructs"
	"errors"
	"strings"
	"unicode"
)

const unaryMinus string = "unary_minus"
const charOperators string = "+-*/^"
const specialChars string = charOperators + "("

func priority(s string) int {
	var priorities = map[string]int{
		"+":           1,
		"-":           2,
		"*":           3,
		"/":           3,
		"^":           4,
		"unary_minus": 5,
	}
	return priorities[s]
}

func isUnaryMinus(char string, i int, infix string) bool {
	return char == "-" &&
		(i == 0 ||
			strings.ContainsAny(string(infix[i-1]), specialChars))
}

func isOperator(char string) bool {
	return strings.ContainsAny(charOperators, char)
}

func processChar(
	i int,
	char rune,
	numberBuffer *strings.Builder,
	postfix []string,
	symbolStack *datastructs.Stack[string],
	infix string,
) ([]string, error) {
	if unicode.IsDigit(char) || char == '.' {
		numberBuffer.WriteRune(char)
		return postfix, nil
	}

	if numberBuffer.Len() != 0 {
		postfix = append(postfix, numberBuffer.String())
		numberBuffer.Reset()
	}

	switch {
	case char == '(':
		symbolStack.Push(string(char))
	case char == ')':
		for top, ok := symbolStack.Top(); top != "("; top, ok = symbolStack.Top() {
			popVal, _ := symbolStack.Pop()
			postfix = append(postfix, popVal)
			if !ok {
				return []string{}, errors.New("неправильные скобки")
			}
		}
		symbolStack.Pop()
	case isUnaryMinus(string(char), i, infix):
		symbolStack.Push(unaryMinus)
	case isOperator(string(char)):
		for top, ok := symbolStack.Top(); ok && priority(string(char)) <= priority(top); top, ok = symbolStack.Top() {
			popVal, _ := symbolStack.Pop()
			postfix = append(postfix, popVal)
		}
		symbolStack.Push(string(char))
	default:
		return []string{}, errors.New("неизвестный символ")
	}
	return postfix, nil
}

func InfixToPostfix(infix string) (datastructs.Stack[string], error) {
	var (
		numberBuffer strings.Builder
		postfix      []string
		symbolStack  datastructs.Stack[string]
		err          error
	)

	for i, char := range infix {
		postfix, err = processChar(i, char, &numberBuffer, postfix, &symbolStack, infix)
		if err != nil {
			return []string{}, err
		}
	}

	if numberBuffer.Len() != 0 {
		postfix = append(postfix, numberBuffer.String())
	}

	for !symbolStack.Empty() {
		popVal, _ := symbolStack.Pop()
		if popVal == "(" {
			return datastructs.Stack[string]{}, errors.New("неправильные скобки")
		}
		postfix = append(postfix, popVal)
	}

	return postfix, nil
}
