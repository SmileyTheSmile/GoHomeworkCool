package lines_filtering

import (
	"GoHomework/cmd_args"
	"fmt"
	"strings"
)

func ChosenLinesGenerator(lines []string, args cmd_args.CommandLineArgs) chan string {
	outChan := make(chan string)
	go generatorLoop(lines, args, outChan)
	return outChan
}

func generatorLoop(lines []string, args cmd_args.CommandLineArgs, outChan chan string) {
	lastRepeatedLine, lineRepetitions := lines[0], 1

	for _, newLine := range lines[1:] {
		if applyArgs(newLine, args) != applyArgs(lastRepeatedLine, args) {
			outChan <- processLine(lastRepeatedLine, lineRepetitions, args)
			lineRepetitions = 0
			lastRepeatedLine = newLine
		}
		lineRepetitions += 1
	}

	outChan <- processLine(lastRepeatedLine, lineRepetitions, args)

	close(outChan)
}

func processLine(line string, repetitionsNum int, args cmd_args.CommandLineArgs) string {
	switch {
	case args.CountOccurrences: // -c
		return fmt.Sprint(repetitionsNum, " ", line, "\n")
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

func applyArgs(line string, args cmd_args.CommandLineArgs) string {
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
