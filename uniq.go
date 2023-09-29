package main

import (
	"GoHomework/cmd_args"
	"GoHomework/lines_filtering"
	"fmt"

	"bufio"
	"os"
)

func readLines(args cmd_args.CommandLineArgs) ([]string, error) {
	var inFile *os.File
	if args.Input == "" {
		inFile = os.Stdin
	} else {
		inFile, err := os.Open(args.Input)
		if err != nil {
			return []string{}, err
		}
		defer inFile.Close()
	}

	var lines []string
	in := bufio.NewScanner(inFile)
	for in.Scan() {
		lines = append(lines, in.Text())
	}
	return lines, nil
}

func writeLines(lines []string, args cmd_args.CommandLineArgs) error {
	var outFile *os.File
	if args.Output == "" {
		outFile = os.Stdout
	} else {
		outFile, err := os.Create(args.Output)
		if err != nil {
			return err
		}
		defer outFile.Close()
	}

	out := bufio.NewWriter(outFile)
	for _, line := range lines_filtering.FilterLines(lines, args) {
		out.WriteString(line)
	}
	out.Flush()
	return nil
}

func main() {
	args, err := cmd_args.GetCommandLineArgs()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	lines, err := readLines(*args)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = writeLines(lines, *args)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
