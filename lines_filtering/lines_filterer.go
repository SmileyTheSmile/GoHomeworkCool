package lines_filtering

import (
	"GoHomework/cmd_args"
	"bufio"
	"io"
)

func ChosenLinesGeneratorEpic(in io.Reader, args cmd_args.CommandLineArgs) chan string {
	inChan := make(chan string)
	outChan := make(chan string)
	go readLines(in, inChan)
	go generatorLoopEpic(args, inChan, outChan)
	return outChan
}

func readLines(in io.Reader, inChan chan string) {
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		inChan <- scanner.Text()
	}
	close(inChan)
}

func generatorLoopEpic(args cmd_args.CommandLineArgs, inChan chan string, outChan chan string) {
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
