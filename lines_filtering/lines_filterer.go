package lines_filtering

import (
	"GoHomework/cmd_args"
	"bufio"
	"fmt"
	"io"
	"strings"
)

func ChosenLinesGenerator(in io.Reader, args cmd_args.CommandLineArgs) chan string {
	inChan := make(chan string)
	outChan := make(chan string)
	go readLines(in, inChan)
	go generatorLoop(args, inChan, outChan)
	return outChan
}

func readLines(in io.Reader, inChan chan string) {
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		inChan <- scanner.Text()
	}
	close(inChan)
}

func generatorLoop(args cmd_args.CommandLineArgs, inChan chan string, outChan chan string) {
	lastRepeatedLine, ok := <-inChan
	if !ok {
		close(outChan)
		return
	}

	lineRepetitions := 1

	for newLine, ok := <-inChan; ok; newLine, ok = <-inChan {
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
		return fmt.Sprint(repetitionsNum, " ", line)
	case args.PrintOnlyRepeated: // -d
		if repetitionsNum > 1 {
			return line
		}
	case args.PrintOnlyNotRepeated: // -u
		if repetitionsNum == 1 {
			return line
		}
	default:
		return line
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
