package loops

import (
	"GoHomework/cmd_args"
	"bufio"
	"fmt"
)

func processArgs(lastRepeatedLine string, repetitionsNum int, args cmd_args.CommandLineArgs) string {
	switch {
	case args.CountOccurrences: // -c
		return fmt.Sprint(repetitionsNum) + " " + lastRepeatedLine + "\n"
	case args.PrintOnlyRepeated: // -d
		if repetitionsNum > 1 {
			return lastRepeatedLine + "\n"
		}
	case args.PrintOnlyNotRepeated: // -u
		if repetitionsNum == 1 {
			return lastRepeatedLine + "\n"
		}
	default:
		return lastRepeatedLine + "\n"
	}
	return ""
}

func ChosenLinesGenerator(in *bufio.Scanner, args cmd_args.CommandLineArgs) chan string {
	ch := make(chan string)

	go func(ch chan string) {
		repetitionsNum := 1

		in.Scan()
		var lastRepeatedLine = in.Text()
		for in.Scan() {
			var newLine = in.Text()
			if !isEqual(newLine, lastRepeatedLine, args) {
				ch <- processArgs(lastRepeatedLine, repetitionsNum, args)
				repetitionsNum = 0
				lastRepeatedLine = newLine
			}
			repetitionsNum += 1
		}
		ch <- processArgs(lastRepeatedLine, repetitionsNum, args)

		close(ch)
	}(ch)

	return ch
}
