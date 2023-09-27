package lines_filtering

import (
	"GoHomework/cmd_args"
	"bufio"
	"fmt"
	"strings"
)

func isEqual(line1 string, line2 string, args cmd_args.CommandLineArgs) bool {
	return formatted(line1, args) == formatted(line2, args)
}

func formatted(line string, args cmd_args.CommandLineArgs) string {
	if args.IgnoreRegister {
		line = strings.ToLower(line)
	}
	if args.NumOfFieldsToIgnore != 0 {
		var splitLine = strings.Fields(line)
		if len(splitLine) <= args.NumOfFieldsToIgnore {
			return ""
		}
		line = strings.Join(splitLine[args.NumOfFieldsToIgnore:], " ")
	}
	if len(line) <= args.NumOfCharsToIgnore {
		return ""
	}
	return line[args.NumOfCharsToIgnore:]
}

func processLine(line string, repetitionsNum int, args cmd_args.CommandLineArgs) string {
	switch {
	case args.CountOccurrences: // -c
		return fmt.Sprint(repetitionsNum) + " " + line + "\n"
	case args.PrintOnlyRepeated: // -d
		if repetitionsNum > 1 {
			return line + "\n"
		}
	case args.PrintOnlyNotRepeated: // -u
		if repetitionsNum == 1 {
			return line + "\n"
		}
	default:
		return line + "\n"
	}
	return ""
}

func ChosenLinesGenerator(in *bufio.Scanner, args cmd_args.CommandLineArgs) chan string {
	ch := make(chan string)

	go func(ch chan string) {
		lineRepetitionsNum := 1

		in.Scan()
		var lastRepeatedLine = in.Text()

		for in.Scan() {
			var newLine = in.Text()
			if !isEqual(newLine, lastRepeatedLine, args) {
				ch <- processLine(lastRepeatedLine, lineRepetitionsNum, args)
				lineRepetitionsNum = 0
				lastRepeatedLine = newLine
			}
			lineRepetitionsNum += 1
		}

		ch <- processLine(lastRepeatedLine, lineRepetitionsNum, args)

		close(ch)
	}(ch)

	return ch
}
