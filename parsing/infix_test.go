package parsing

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
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
		/*
			{
				name: "Унарная функция",
				args: args{
					infix: "1+2*(3/4)+sqrt(4)", // 4.5
				},
				want:    "1234/*+4sqrt+",
				wantErr: false,
			},
		*/
		{
			name: "Неизвестный оператор",
			args: args{
				infix: "1+2*(3/4)+haha(5)",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Числа из нескольких цифр",
			args: args{
				infix: "1+2*(3/4)+56", // 58.5
			},
			want:    "1234/*+56+",
			wantErr: false,
		},
		{
			name: "Дроби",
			args: args{
				infix: "1+2*(3/4)+56+7.8", // 66.3
			},
			want:    "1234/*+56+7.8+",
			wantErr: false,
		},
		{
			name: "Неизвестный оператор в конце",
			args: args{
				infix: "1+2*(3/4)+56+7.8hhaha",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Неправильные скобки",
			args: args{
				infix: "1+2*(3/4))+56+7.8",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Неправильные скобки",
			args: args{
				infix: "1+2*((3/4)+56+7.8",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got_raw, err := InfixToPostfix(tt.args.infix)
			got := strings.Join(got_raw.ToSlice(), "")
			require.Equal(t, got, tt.want, tt.name)
			require.Equal(t, err != nil, tt.wantErr, tt.name)
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
			got := priority(tt.args.s)
			require.Equal(t, got, tt.want, tt.name)
		})
	}
}
