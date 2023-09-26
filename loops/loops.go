package loops

import (
	"bufio"
	"fmt"
)

func NormalLoop(in *bufio.Scanner, out *bufio.Writer, numFields int, numChars int, ignoreRegister bool) {
	var lastRepeated string
	for in.Scan() {
		var newLine = in.Text()
		if areDifferent(newLine, lastRepeated, numFields, numChars, ignoreRegister) {
			lastRepeated = newLine
			out.WriteString(lastRepeated + "\n")
		}
	}
}

// -c
func CountOccurrencesLoop(in *bufio.Scanner, out *bufio.Writer, numFields int, numChars int, ignoreRegister bool) {
	repetitionsNum := 1

	in.Scan()
	var lastRepeated = in.Text()
	for in.Scan() {
		var newLine = in.Text()
		if areDifferent(newLine, lastRepeated, numFields, numChars, ignoreRegister) {
			out.WriteString(fmt.Sprint(repetitionsNum) + " " + lastRepeated + "\n")
			lastRepeated = newLine
			repetitionsNum = 0
		}
		repetitionsNum += 1
	}
	out.WriteString(fmt.Sprint(repetitionsNum) + " " + lastRepeated + "\n")
}

// -d
func PrintOnlyRepeatedLoop(in *bufio.Scanner, out *bufio.Writer, numFields int, numChars int, ignoreRegister bool) {
	in.Scan()
	var lastLine = in.Text()
	var lastRepeated = lastLine + "diff"
	for in.Scan() {
		var newLine = in.Text()
		if areDifferent(newLine, lastLine, numFields, numChars, ignoreRegister) {
			lastRepeated = lastLine
		} else if areDifferent(lastRepeated, lastLine, numFields, numChars, ignoreRegister) {
			out.WriteString(lastLine + "\n")
			lastRepeated = lastLine
		}
		lastLine = newLine
	}
}

// -u
func PrintOnlyNotRepeatedLoop(in *bufio.Scanner, out *bufio.Writer, numFields int, numChars int, ignoreRegister bool) {
	in.Scan()
	var lastLine = in.Text()
	var lastRepeated = lastLine
	for in.Scan() {
		var newLine = in.Text()
		if areDifferent(lastRepeated, lastLine, numFields, numChars, ignoreRegister) {
			if areDifferent(newLine, lastLine, numFields, numChars, ignoreRegister) {
				out.WriteString(lastLine + "\n")
			}
			lastRepeated = lastLine
		}
		lastLine = newLine
	}
	if areDifferent(lastRepeated, lastLine, numFields, numChars, ignoreRegister) {
		out.WriteString(lastLine + "\n")
	}
}
