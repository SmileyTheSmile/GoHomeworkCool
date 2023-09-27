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

	var inFile *os.File
	if args.Input == "" {
		inFile = os.Stdin
	} else {
		inFile, err = os.Open(args.Input)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		defer inFile.Close()
	}

	var outFile *os.File
	if args.Output == "" {
		outFile = os.Stdout
	} else {
		outFile, err = os.Create(args.Output)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		defer outFile.Close()
	}

	out := bufio.NewWriter(outFile)

	for line := range lines_filtering.ChosenLinesGenerator(inFile, *args) {
		out.WriteString(line + "\n")
	}
	out.Flush()
}
