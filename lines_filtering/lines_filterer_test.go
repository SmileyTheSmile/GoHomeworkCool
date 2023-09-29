package lines_filtering

import (
	"GoHomework/cmd_args"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_FilterLines(t *testing.T) {
	type args struct {
		lines []string
		args  cmd_args.CommandLineArgs
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Без параметров",
			args: args{
				lines: []string{
					"I love music.",
					"I love music.",
					"I love music.",
					"",
					"I love music of Kartik.",
					"I love music of Kartik.",
					"Thanks.",
					"I love music of Kartik.",
					"I love music of Kartik.",
				},
				args: cmd_args.CommandLineArgs{},
			},
			want: []string{
				"I love music.\n",
				"\n",
				"I love music of Kartik.\n",
				"Thanks.\n",
				"I love music of Kartik.\n",
			},
		},
		{
			name: "С параметром -c",
			args: args{
				lines: []string{
					"I love music.",
					"I love music.",
					"I love music.",
					"",
					"I love music of Kartik.",
					"I love music of Kartik.",
					"Thanks.",
					"I love music of Kartik.",
					"I love music of Kartik.",
				},
				args: cmd_args.CommandLineArgs{
					CountOccurrences: true,
				},
			},
			want: []string{
				"3 I love music.\n",
				"1 \n",
				"2 I love music of Kartik.\n",
				"1 Thanks.\n",
				"2 I love music of Kartik.\n",
			},
		},
		{
			name: "С параметром -d",
			args: args{
				lines: []string{
					"I love music.",
					"I love music.",
					"I love music.",
					"",
					"I love music of Kartik.",
					"I love music of Kartik.",
					"Thanks.",
					"I love music of Kartik.",
					"I love music of Kartik.",
				},
				args: cmd_args.CommandLineArgs{
					PrintOnlyRepeated: true,
				},
			},
			want: []string{
				"I love music.\n",
				"I love music of Kartik.\n",
				"I love music of Kartik.\n",
			},
		},
		{
			name: "С параметром -u",
			args: args{
				lines: []string{
					"I love music.",
					"I love music.",
					"I love music.",
					"",
					"I love music of Kartik.",
					"I love music of Kartik.",
					"Thanks.",
					"I love music of Kartik.",
					"I love music of Kartik.",
				},
				args: cmd_args.CommandLineArgs{
					PrintOnlyNotRepeated: true,
				},
			},
			want: []string{
				"\n",
				"Thanks.\n",
			},
		},
		{
			name: "С параметром -i",
			args: args{
				lines: []string{
					"I LOVE music.",
					"I love music.",
					"I love music.",
					"",
					"I love Music of Kartik.",
					"I love music of Kartik.",
					"Thanks.",
					"I love MuSIC of Kartik.",
					"I love music of Kartik.",
				},
				args: cmd_args.CommandLineArgs{
					IgnoreRegister: true,
				},
			},
			want: []string{
				"I LOVE music.\n",
				"\n",
				"I love Music of Kartik.\n",
				"Thanks.\n",
				"I love MuSIC of Kartik.\n",
			},
		},
		{
			name: "С параметром -f 1",
			args: args{
				lines: []string{
					"We love music.",
					"I love music.",
					"They love music.",
					"",
					"I love music of Kartik.",
					"We love music of Kartik.",
					"Thanks.",
				},
				args: cmd_args.CommandLineArgs{
					NumOfFieldsToIgnore: 1,
				},
			},
			want: []string{
				"We love music.\n",
				"\n",
				"I love music of Kartik.\n",
				"Thanks.\n",
			},
		},
		{
			name: "С параметром -s 1",
			args: args{
				lines: []string{
					"I love music.",
					"A love music.",
					"C love music.",
					"",
					"I love music of Kartik.",
					"We love music of Kartik.",
					"Thanks.",
				},
				args: cmd_args.CommandLineArgs{
					NumOfCharsToIgnore: 1,
				},
			},
			want: []string{
				"I love music.\n",
				"\n",
				"I love music of Kartik.\n",
				"We love music of Kartik.\n",
				"Thanks.\n",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []string = FilterLines(tt.args.lines, tt.args.args)
			require.Equal(t, got, tt.want, tt.name)
		})
	}
}

func Test_applyArgs(t *testing.T) {
	type args struct {
		line string
		args cmd_args.CommandLineArgs
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Регистр",
			args: args{
				line: "loL",
				args: cmd_args.CommandLineArgs{IgnoreRegister: true},
			},
			want: "lol",
		},
		{
			name: "Игнорирование слов",
			args: args{
				line: "lol lol",
				args: cmd_args.CommandLineArgs{NumOfFieldsToIgnore: 1},
			},
			want: "lol",
		},
		{
			name: "Игнорирование букв",
			args: args{
				line: "lol lolol",
				args: cmd_args.CommandLineArgs{NumOfCharsToIgnore: 4},
			},
			want: "lolol",
		},
		{
			name: "Игнорирование букв и слов",
			args: args{
				line: "lol lololol",
				args: cmd_args.CommandLineArgs{
					NumOfFieldsToIgnore: 1,
					NumOfCharsToIgnore:  4,
				},
			},
			want: "lol",
		},
		{
			name: "Игнорирование букв, слов и регистра",
			args: args{
				line: "lOl lololOL",
				args: cmd_args.CommandLineArgs{
					IgnoreRegister:      true,
					NumOfFieldsToIgnore: 1,
					NumOfCharsToIgnore:  4,
				},
			},
			want: "lol",
		},
		{
			name: "Игнорирование букв, короткая строка",
			args: args{
				line: "lOl",
				args: cmd_args.CommandLineArgs{
					NumOfCharsToIgnore: 4,
				},
			},
			want: "",
		},
		{
			name: "Игнорирование слов, короткая строка",
			args: args{
				line: "lOl LOL",
				args: cmd_args.CommandLineArgs{
					NumOfFieldsToIgnore: 3,
				},
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := applyArgs(tt.args.line, tt.args.args)
			require.Equal(t, got, tt.want, tt.name)
		})
	}
}

func Test_processLine(t *testing.T) {
	type args struct {
		line           string
		repetitionsNum int
		args           cmd_args.CommandLineArgs
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 bool
	}{
		{
			name: "Посчитать количество входов",
			args: args{
				line:           "lol",
				repetitionsNum: 3,
				args: cmd_args.CommandLineArgs{
					CountOccurrences: true,
				},
			},
			want:  "3 lol\n",
			want1: true,
		},
		{
			name: "Вывод повторящихся строк с флагом -d",
			args: args{
				line:           "lol",
				repetitionsNum: 3,
				args: cmd_args.CommandLineArgs{
					PrintOnlyRepeated: true,
				},
			},
			want:  "lol\n",
			want1: true,
		},
		{
			name: "Вывод не повторящейся строки с флагом -d",
			args: args{
				line:           "lol",
				repetitionsNum: 1,
				args: cmd_args.CommandLineArgs{
					PrintOnlyRepeated: true,
				},
			},
			want:  "",
			want1: false,
		},
		{
			name: "Вывод повторящихся строк с флагом -u",
			args: args{
				line:           "lol",
				repetitionsNum: 3,
				args: cmd_args.CommandLineArgs{
					PrintOnlyNotRepeated: true,
				},
			},
			want:  "",
			want1: false,
		},
		{
			name: "Вывод не повторящейся строки с флагом -u",
			args: args{
				line:           "lol",
				repetitionsNum: 1,
				args: cmd_args.CommandLineArgs{
					PrintOnlyNotRepeated: true,
				},
			},
			want:  "lol\n",
			want1: true,
		},
		{
			name: "Вывод строки без флагов",
			args: args{
				line:           "lol",
				repetitionsNum: 1,
				args:           cmd_args.CommandLineArgs{},
			},
			want:  "lol\n",
			want1: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := processLine(tt.args.line, tt.args.repetitionsNum, tt.args.args)
			require.Equal(t, got, tt.want, tt.name)
			require.Equal(t, got1, tt.want1, tt.name)
		})
	}
}
