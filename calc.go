package main

import (
	"GoHomework/parsing"
	"flag"
	"fmt"
	"strings"
)

func main() {
	flag.Parse()
	var otherArgs []string = flag.Args()
	if len(otherArgs) < 1 {
		fmt.Println("Нет выражения")
		return
	} else if len(otherArgs) > 1 {
		fmt.Println("Лишние аргуметы")
		return
	}

	var infixExpression string = otherArgs[0]

	infixExpression = strings.ReplaceAll(infixExpression, " ", "")

	postfixExpression, err := parsing.InfixToPostfix(infixExpression)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(strings.Join(postfixExpression.ToSlice(), ""))

	result, err := parsing.SolvePostfix(postfixExpression.ToSlice())
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(result)
}
