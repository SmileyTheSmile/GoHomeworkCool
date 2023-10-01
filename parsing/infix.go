package parsing

import (
	datastructs "GoHomework/datastructs"
	"errors"
	"strings"
	"unicode"
)

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
			strings.ContainsAny(string(infix[i-1]), "(+-*/^"))
}

func isOperator(char string) bool {
	return strings.ContainsAny("+-*/^", char)
}

func processChar(
	i int,
	char rune,
	numberBuffer *strings.Builder,
	postfix *[]string,
	symbolStack *datastructs.Stack[string],
	infix string,
) error {
	if unicode.IsDigit(char) || char == '.' {
		numberBuffer.WriteRune(char)
		return nil
	}

	if numberBuffer.Len() != 0 {
		*postfix = append(*postfix, numberBuffer.String())
		numberBuffer.Reset()
	}

	switch {
	case char == '(':
		symbolStack.Push(string(char))
	case char == ')':
		for top, ok := symbolStack.Top(); top != "("; top, ok = symbolStack.Top() {
			popVal, _ := symbolStack.Pop()
			*postfix = append(*postfix, popVal)
			if !ok {
				return errors.New("Неправильные скобки")
			}
		}
		symbolStack.Pop()
	case isUnaryMinus(string(char), i, infix):
		symbolStack.Push("unary_minus")
	case isOperator(string(char)):
		for top, ok := symbolStack.Top(); ok && priority(string(char)) <= priority(top); top, ok = symbolStack.Top() {
			popVal, _ := symbolStack.Pop()
			*postfix = append(*postfix, popVal)
		}
		symbolStack.Push(string(char))
	default:
		return errors.New("Неизвестный символ")
	}
	return nil
}

func InfixToPostfix(infix string) (datastructs.Stack[string], error) {
	var (
		numberBuffer strings.Builder
		postfixStack []string
		symbolStack  datastructs.Stack[string]
	)

	for i, char := range infix {
		err := processChar(i, char, &numberBuffer, &postfixStack, &symbolStack, infix)
		if err != nil {
			return datastructs.Stack[string]{}, err
		}
	}

	if numberBuffer.Len() != 0 {
		postfixStack = append(postfixStack, numberBuffer.String())
	}

	for !symbolStack.Empty() {
		popVal, _ := symbolStack.Pop()
		if popVal == "(" {
			return datastructs.Stack[string]{}, errors.New("Неправильные скобки")
		}
		postfixStack = append(postfixStack, popVal)
	}

	return postfixStack, nil
}
