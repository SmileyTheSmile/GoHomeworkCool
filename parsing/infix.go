package parsing

import (
	datastructs "GoHomework/datastructs"
	"errors"
	"strings"
	"unicode"
)

func priority(s string) int {
	var priorities = map[string]int{
		"+": 1,
		"-": 2,
		"*": 3,
		"/": 3,
		"^": 4,
		//"unary_minus": 5,
		//"sqrt":        6,
		//"ceil":        6,
	}
	return priorities[s]
}

func isUnaryMinus(char string, i int, infix string) bool {
	return char == "-" &&
		(i == 0 ||
			strings.ContainsAny(string(infix[i-1]), "(+-*/^"))
}

func isSpecialChar(char string) bool {
	return strings.ContainsAny("+-*/^", char)
}

func isOperator(word string) bool {
	return word == "ceil" || word == "sqrt" || isSpecialChar(word)
}

func isNotMoreImportantOperator(char string, top string, symbolStack datastructs.Stack[string]) bool {
	return !symbolStack.Empty() && priority(char) <= priority(top)
}

func processChar(
	i int,
	char rune,
	numberBuffer *strings.Builder,
	operatorBuffer *strings.Builder,
	postfixStack *datastructs.Stack[string],
	symbolStack *datastructs.Stack[string],
	infix string,
) error {
	if unicode.IsDigit(char) || char == '.' {
		numberBuffer.WriteRune(char)
		return nil
	}
	if unicode.IsLetter(char) {
		operatorBuffer.WriteRune(char)
		return nil
	}

	if numberBuffer.Len() != 0 {
		postfixStack.Push(numberBuffer.String())
		numberBuffer.Reset()
	}
	if operatorBuffer.Len() != 0 {
		word := operatorBuffer.String()
		if !isOperator(word) {
			return errors.New("Неверный оператор")
		}
		symbolStack.Push(word)
		operatorBuffer.Reset()
	}

	switch {
	case char == '(':
		symbolStack.Push(string(char))
	case char == ')':
		for top, ok := symbolStack.Top(); top != "("; top, ok = symbolStack.Top() {
			popVal, _ := symbolStack.Pop()
			postfixStack.Push(popVal)
			if !ok {
				return errors.New("Неправильные скобки")
			}
		}
		symbolStack.Pop()
	case isUnaryMinus(string(char), i, infix):
		symbolStack.Push("unary_minus")
	case isSpecialChar(string(char)):
		for top, _ := symbolStack.Top(); isNotMoreImportantOperator(string(char), top, *symbolStack); top, _ = symbolStack.Top() {
			popVal, _ := symbolStack.Pop()
			postfixStack.Push(popVal)
		}
		symbolStack.Push(string(char))
	default:
		return errors.New("Неизвестный символ")
	}
	return nil
}

func InfixToPostfix(infix string) (datastructs.Stack[string], error) {
	var numberBuffer strings.Builder
	var operatorBuffer strings.Builder
	var postfixStack datastructs.Stack[string]
	var symbolStack datastructs.Stack[string]

	for i, char := range infix {
		err := processChar(i, char, &numberBuffer, &operatorBuffer, &postfixStack, &symbolStack, infix)
		if err != nil {
			return datastructs.Stack[string]{}, err
		}
	}

	if numberBuffer.Len() != 0 {
		postfixStack.Push(numberBuffer.String())
	}

	if operatorBuffer.Len() != 0 {
		return datastructs.Stack[string]{}, errors.New("Странный оператор в конце выражения")
	}

	for !symbolStack.Empty() {
		popVal, _ := symbolStack.Pop()
		if popVal == "(" {
			return datastructs.Stack[string]{}, errors.New("Неправильные скобки")
		}
		postfixStack.Push(popVal)
	}

	return postfixStack, nil
}

/*
func processCharIncomplete(
	i int,
	char rune,
	numberBuffer *strings.Builder,
	operatorBuffer *strings.Builder,
	postfixStack *datastructs.Stack[string],
	symbolStack *datastructs.Stack[string],
	infix string,
) error {
	if unicode.IsDigit(char) || char == '.' {
		numberBuffer.WriteRune(char)
		return nil
	} else if unicode.IsLetter(char) {
		operatorBuffer.WriteRune(char)
		return nil
	} else if isSpecialChar(string(char)) {
		fmt.Println(char)
		operatorBuffer.WriteRune(char)
	}

	if numberBuffer.Len() != 0 {
		postfixStack.Push(numberBuffer.String())
		numberBuffer.Reset()
		return nil
	}
	if operatorBuffer.Len() != 0 {
		word := operatorBuffer.String()
		if !isOperator(word) {
			return errors.New("Неверный оператор")
		}
		if isUnaryMinus(word, i, infix) {
			symbolStack.Push("unary_minus")
			return nil
		}
		for top, _ := symbolStack.Top(); isNotMoreImportantOperator(word, top, *symbolStack); top, _ = symbolStack.Top() {
			popVal, _ := symbolStack.Pop()
			postfixStack.Push(popVal)
		}
		symbolStack.Push(word)
		operatorBuffer.Reset()
		return nil
	}

	switch {
	case char == '(':
		symbolStack.Push(string(char))
	case char == ')':
		for top, _ := symbolStack.Top(); top != "("; top, _ = symbolStack.Top() {
			popVal, _ := symbolStack.Pop()
			postfixStack.Push(popVal)
		}
		symbolStack.Pop()
	default:
		return errors.New("Неизвестный символ")
	}
	return nil
}
*/
