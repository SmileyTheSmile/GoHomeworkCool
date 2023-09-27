package loops

import (
	"GoHomework/cmd_args"
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
