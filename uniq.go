package main

import (
	"GoHomework/cmd_args"
	"GoHomework/loops"
	"fmt"

	"bufio"
	"os"
)

func main() {
	args, err := cmd_args.GetCommandLineArgs()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var in *bufio.Scanner
	if args.Input == "" {
		in = bufio.NewScanner(os.Stdin)
	} else {
		file, err := os.Open(args.Input)
		if err != nil {
			fmt.Println(err.Error())
		}
		defer file.Close()
		in = bufio.NewScanner(file)
	}

	var out *bufio.Writer
	if args.Output == "" {
		out = bufio.NewWriter(os.Stdout)
	} else {
		file, err := os.OpenFile(args.Output, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
		if err != nil {
			fmt.Println(err.Error())
		}
		defer file.Close()
		out = bufio.NewWriter(file)
	}

	switch {
	case *args.CountOccurrences:
		loops.CountOccurrencesLoop(in, out, *args.NumOfFieldsToIgnore, *args.NumOfCharsToIgnore, *args.IgnoreRegister)
	case *args.PrintOnlyRepeated:
		loops.PrintOnlyRepeatedLoop(in, out, *args.NumOfFieldsToIgnore, *args.NumOfCharsToIgnore, *args.IgnoreRegister)
	case *args.PrintOnlyNotRepeated:
		loops.PrintOnlyNotRepeatedLoop(in, out, *args.NumOfFieldsToIgnore, *args.NumOfCharsToIgnore, *args.IgnoreRegister)
	default:
		loops.NormalLoop(in, out, *args.NumOfFieldsToIgnore, *args.NumOfCharsToIgnore, *args.IgnoreRegister)
	}
	out.Flush()
}
