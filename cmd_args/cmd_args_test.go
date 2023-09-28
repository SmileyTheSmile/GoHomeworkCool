package cmd_args

import (
	"reflect"
	"testing"
)

func TestCommandLineArgs_ConflictingArgs(t *testing.T) {
	type fields struct {
		CountOccurrences     bool
		PrintOnlyRepeated    bool
		PrintOnlyNotRepeated bool
		IgnoreRegister       bool
		NumOfFieldsToIgnore  int
		NumOfCharsToIgnore   int
		Input                string
		Output               string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "-с",
			fields: fields{
				CountOccurrences: true,
			},
			want: false,
		},
		{
			name: "-d",
			fields: fields{
				PrintOnlyRepeated: true,
			},
			want: false,
		},
		{
			name: "-u",
			fields: fields{
				PrintOnlyNotRepeated: true,
			},
			want: false,
		},
		{
			name: "-с -u",
			fields: fields{
				CountOccurrences:     true,
				PrintOnlyNotRepeated: true,
			},
			want: true,
		},
		{
			name: "Только -с -d",
			fields: fields{
				CountOccurrences:  true,
				PrintOnlyRepeated: true,
			},
			want: true,
		},
		{
			name: "-d -u",
			fields: fields{
				PrintOnlyRepeated:    true,
				PrintOnlyNotRepeated: true,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			args := &CommandLineArgs{
				CountOccurrences:     tt.fields.CountOccurrences,
				PrintOnlyRepeated:    tt.fields.PrintOnlyRepeated,
				PrintOnlyNotRepeated: tt.fields.PrintOnlyNotRepeated,
				IgnoreRegister:       tt.fields.IgnoreRegister,
				NumOfFieldsToIgnore:  tt.fields.NumOfFieldsToIgnore,
				NumOfCharsToIgnore:   tt.fields.NumOfCharsToIgnore,
				Input:                tt.fields.Input,
				Output:               tt.fields.Output,
			}
			if got := args.ConflictingArgs(); got != tt.want {
				t.Errorf("ConflictingArgs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetCommandLineArgs(t *testing.T) {
	tests := []struct {
		name    string
		want    *CommandLineArgs
		wantErr bool
	}{
		{
			name:    "Получить аргументы",
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetCommandLineArgs()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCommandLineArgs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCommandLineArgs() got = %v, want %v", got, tt.want)
			}
		})
	}
}
