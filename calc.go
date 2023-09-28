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
	} else if len(otherArgs) > 1 {
		fmt.Println("Лишние аргуметы")
	}

	var infixExpression string = otherArgs[0] // "1+2*(3^4-5.6)*(7+8.9*10)-sqrt(4)-11" == 14462.8

	infixExpression = strings.ReplaceAll(infixExpression, " ", "")

	//fmt.Println(infixExpression)
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
