package cmd_args

import (
	"errors"
	"flag"
)

type CommandLineArgs struct {
	CountOccurrences     bool
	PrintOnlyRepeated    bool
	PrintOnlyNotRepeated bool
	IgnoreRegister       bool
	NumOfFieldsToIgnore  int
	NumOfCharsToIgnore   int
	Input                string
	Output               string
}

func (args *CommandLineArgs) ConflictingArgs() bool {
	if args.CountOccurrences && args.PrintOnlyRepeated ||
		args.CountOccurrences && args.PrintOnlyNotRepeated ||
		args.PrintOnlyNotRepeated && args.PrintOnlyRepeated {
		return true
	}
	return false
}

func GetCommandLineArgs() (*CommandLineArgs, error) {
	var args = CommandLineArgs{}

	flag.BoolVar(&args.CountOccurrences, "c", false, "Подсчитать количество встречаний строки во входных данных.")
	flag.BoolVar(&args.PrintOnlyRepeated, "d", false, "Вывести только те строки, которые повторились во входных данных.")
	flag.BoolVar(&args.PrintOnlyNotRepeated, "u", false, "Вывести только те строки, которые не повторились во входных данных.")
	flag.BoolVar(&args.IgnoreRegister, "i", false, "Не учитывать регистр букв.")
	flag.IntVar(&args.NumOfFieldsToIgnore, "f", 0, "Не учитывать первые num_fields полей в строке. Полем в строке является непустой набор символов отделённый пробелом.")
	flag.IntVar(&args.NumOfCharsToIgnore, "s", 0, "Не учитывать первые num_chars символов в строке. При использовании вместе с параметром -f учитываются первые символы после num_fields полей (не учитывая пробел-разделитель после последнего поля).")
	flag.Parse()

	if args.ConflictingArgs() {
		return nil, errors.New("Конфликтующие флаги.")
	}

	var otherArgs []string = flag.Args()
	if len(otherArgs) < 1 {
		return nil, errors.New("Нет ввода")
	} else if len(otherArgs) == 1 {
		args.Input = otherArgs[0]
	} else if len(otherArgs) == 2 {
		args.Input, args.Output = otherArgs[0], otherArgs[1]
	}

	return &args, nil
}
