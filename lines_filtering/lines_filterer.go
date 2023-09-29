package lines_filtering

import (
	"GoHomework/cmd_args"
	"fmt"
	"strings"
)

func FilterLines(lines []string, args cmd_args.CommandLineArgs) []string {
	lastRepeatedLine, lineRepetitions := lines[0], 1

	var result []string
	for _, newLine := range lines[1:] {
		if applyArgs(newLine, args) != applyArgs(lastRepeatedLine, args) {
			line, ok := processLine(lastRepeatedLine, lineRepetitions, args)
			if ok {
				result = append(result, line)
			}
			lineRepetitions = 0
			lastRepeatedLine = newLine
		}
		lineRepetitions += 1
	}

	line, ok := processLine(lastRepeatedLine, lineRepetitions, args)
	if ok {
		result = append(result, line)
	}

	return result
}

func processLine(line string, repetitionsNum int, args cmd_args.CommandLineArgs) (string, bool) {
	switch {
	case args.CountOccurrences: // -c
		return fmt.Sprint(repetitionsNum, " ", line, "\n"), true
	case args.PrintOnlyRepeated: // -d
		if repetitionsNum > 1 {
			return line + "\n", true
		}
	case args.PrintOnlyNotRepeated: // -u
		if repetitionsNum == 1 {
			return line + "\n", true
		}
	default:
		return line + "\n", true
	}
	return "", false
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
