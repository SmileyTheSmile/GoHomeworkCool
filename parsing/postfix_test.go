package parsing

import "testing"

func TestSolvePostfix(t *testing.T) {
	type args struct {
		postfixExpression []string
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{
			name: "Простой пример",
			args: args{
				postfixExpression: []string{"1", "2", "+"},
			},
			want:    3,
			wantErr: false,
		},
		{
			name: "Скобки",
			args: args{
				postfixExpression: []string{"1", "2", "3", "4", "/", "*", "+"},
			},
			want:    2.5,
			wantErr: false,
		},
		/*
			{
				name: "Унарная функция",
				args: args{
					postfixExpression: []string{"1", "2", "3", "4", "/", "*", "+", "4", "sqrt", "+"},
				},
				want:    4.5,
				wantErr: false,
			},
		*/
		{
			name: "Числа из нескольких цифр",
			args: args{
				postfixExpression: []string{"1", "2", "3", "4", "/", "*", "+", "56", "+"},
			},
			want:    58.5,
			wantErr: false,
		},
		{
			name: "Дроби",
			args: args{
				postfixExpression: []string{"1", "2", "3", "4", "/", "*", "+", "+", "56", "+", "7.8", "+"},
			},
			want:    66.3,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SolvePostfix(tt.args.postfixExpression)
			if (err != nil) != tt.wantErr {
				t.Errorf("SolvePostfix() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SolvePostfix() got = %v, want %v", got, tt.want)
			}
		})
	}
}
