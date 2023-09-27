package main

import (
	"GoHomework/cmd_args"
	"GoHomework/lines_filtering"
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
			return
		}
		defer file.Close()
		in = bufio.NewScanner(file)
	}

	var out *bufio.Writer
	if args.Output == "" {
		out = bufio.NewWriter(os.Stdout)
	} else {
		file, err := os.Create(args.Output)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		defer file.Close()
		out = bufio.NewWriter(file)
	}

	for line := range lines_filtering.ChosenLinesGenerator(in, *args) {
		out.WriteString(line)
	}
	out.Flush()
}
