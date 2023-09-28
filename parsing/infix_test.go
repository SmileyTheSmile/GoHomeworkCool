package parsing

import (
	"reflect"
	"strings"
	"testing"
)

func TestInfixToPostfix(t *testing.T) {
	type args struct {
		infix string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Простой пример",
			args: args{
				infix: "1+2", // 3
			},
			want:    "12+",
			wantErr: false,
		},
		{
			name: "Скобки",
			args: args{
				infix: "1+2*(3/4)", // 2.5
			},
			want:    "1234/*+",
			wantErr: false,
		},
		{
			name: "Унарная функция",
			args: args{
				infix: "1+2*(3/4)+sqrt(4)", // 4.5
			},
			want:    "1234/*+4sqrt+",
			wantErr: false,
		},
		{
			name: "Неизвестный оператор",
			args: args{
				infix: "1+2*(3/4)+sqrt(4)+haha(5)",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Числа из нескольких цифр",
			args: args{
				infix: "1+2*(3/4)+sqrt(4)+56", // 60.5
			},
			want:    "1234/*+4sqrt+56+",
			wantErr: false,
		},
		{
			name: "Дроби",
			args: args{
				infix: "1+2*(3/4)+sqrt(4)+56+7.8", // 68.3
			},
			want:    "1234/*+4sqrt+56+7.8+",
			wantErr: false,
		},
		{
			name: "Неизвестный оператор в конце",
			args: args{
				infix: "1+2*(3/4)+sqrt(4)+56+7.8hhaha",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got_raw, err := InfixToPostfix(tt.args.infix)
			got := strings.Join(got_raw.ToSlice(), "")
			if (err != nil) != tt.wantErr {
				t.Errorf("InfixToPostfix() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InfixToPostfix() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isNotMoreImportantOperator(t *testing.T) {
	type args struct {
		char        string
		top         string
		symbolStack Stack[string]
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isNotMoreImportantOperator(tt.args.char, tt.args.top, tt.args.symbolStack); got != tt.want {
				t.Errorf("isNotMoreImportantOperator() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isOperator(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isOperator(tt.args.word); got != tt.want {
				t.Errorf("isOperator() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isPartOfNumber(t *testing.T) {
	type args struct {
		char rune
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPartOfNumber(tt.args.char); got != tt.want {
				t.Errorf("isPartOfNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isSpecialChar(t *testing.T) {
	type args struct {
		char string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSpecialChar(tt.args.char); got != tt.want {
				t.Errorf("isSpecialChar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isUnaryMinus(t *testing.T) {
	type args struct {
		char  string
		i     int
		infix string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isUnaryMinus(tt.args.char, tt.args.i, tt.args.infix); got != tt.want {
				t.Errorf("isUnaryMinus() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_priority(t *testing.T) {
	type args struct {
		s string
	}
	type test struct {
		name string
		args args
		want int
	}
	tests := []test{}
	var priorities = map[string]int{
		"+":           1,
		"-":           2,
		"*":           3,
		"/":           3,
		"^":           4,
		"unary_minus": 5,
		"sqrt":        6,
		"ceil":        6,
	}
	keys := make([]string, 0, len(priorities))
	for _, key := range keys {
		tests = append(tests, test{
			name: key,
			args: args{key},
			want: priorities[key],
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := priority(tt.args.s); got != tt.want {
				t.Errorf("priority() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_processChar(t *testing.T) {
	type args struct {
		i              int
		char           rune
		numberBuffer   *strings.Builder
		operatorBuffer *strings.Builder
		postfixStack   *Stack[string]
		symbolStack    *Stack[string]
		infix          string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := processChar(tt.args.i, tt.args.char, tt.args.numberBuffer, tt.args.operatorBuffer, tt.args.postfixStack, tt.args.symbolStack, tt.args.infix); (err != nil) != tt.wantErr {
				t.Errorf("processChar() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_processCharIncomplete(t *testing.T) {
	type args struct {
		i              int
		char           rune
		numberBuffer   *strings.Builder
		operatorBuffer *strings.Builder
		postfixStack   *Stack[string]
		symbolStack    *Stack[string]
		infix          string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := processCharIncomplete(tt.args.i, tt.args.char, tt.args.numberBuffer, tt.args.operatorBuffer, tt.args.postfixStack, tt.args.symbolStack, tt.args.infix); (err != nil) != tt.wantErr {
				t.Errorf("processCharIncomplete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
