package loops

import (
	"strings"
)

func areDifferent(line1 string, line2 string, numFields int, numChars int, ignoreRegister bool) bool {
	return formatted(line1, numFields, numChars, ignoreRegister) != formatted(line2, numFields, numChars, ignoreRegister)
}

func formatted(line string, numFields int, numChars int, ignoreRegister bool) string {
	if ignoreRegister {
		line = strings.ToLower(line)
	}
	if numFields != 0 {
		var splitLine = strings.Split(line, " ")
		if len(splitLine) <= numFields {
			return ""
		}
		line = strings.Join(splitLine[numFields:], " ")
	}
	if len(line) <= numChars {
		return ""
	}
	return line[numChars:]
}
