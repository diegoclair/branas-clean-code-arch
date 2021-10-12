package utils

import (
	"testing"
)

func Test_invalidEqualNumbers(t *testing.T) {
	type args struct {
		document string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Equal numbers should return invalid=true",
			args: args{
				document: "11111111111",
			},
			want: true,
		},
		{
			name: "Different numbers should return invalid=false",
			args: args{
				document: "11111311111",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := invalidEqualNumbers(tt.args.document); got != tt.want {
				t.Errorf("invalidEqualNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculateCPFDigits(t *testing.T) {
	type args struct {
		document string
	}
	tests := []struct {
		name            string
		args            args
		wantFirstDigit  int
		wantSecondDigit int
	}{
		{
			name: "should return 8 for first digit and 0 for second digit",
			args: args{
				document: "93541134780",
			},
			wantFirstDigit:  8,
			wantSecondDigit: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFirstDigit, gotSecondDigit := calculateCPFDigits(tt.args.document)
			if gotFirstDigit != tt.wantFirstDigit {
				t.Errorf("calculateCPFDigits() gotFirstDigit = %v, want %v", gotFirstDigit, tt.wantFirstDigit)
			}
			if gotSecondDigit != tt.wantSecondDigit {
				t.Errorf("calculateCPFDigits() gotSecondDigit = %v, want %v", gotSecondDigit, tt.wantSecondDigit)
			}
		})
	}
}

func Test_validateCPF(t *testing.T) {
	type args struct {
		cpf string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Valid Document 935.411.347-80",
			args: args{
				cpf: "93541134780",
			},
			want: true,
		},
		{
			name: "Invalid Document 123.456.789-99",
			args: args{
				cpf: "12345678999",
			},
			want: false,
		},
		{
			name: "Invalid Document 111.111.111-11",
			args: args{
				cpf: "11111111111",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validateCPF(tt.args.cpf); got != tt.want {
				t.Errorf("validateCPF() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsValidDocumentNumber(t *testing.T) {
	type args struct {
		document string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Valid Document 935.411.347-80",
			args: args{
				document: "935.411.347-80",
			},
			want: true,
		},
		{
			name: "Invalid Document 123.456.789-99",
			args: args{
				document: "123.456.789-99",
			},
			want: false,
		},
		{
			name: "Invalid Document 111.111.111-11",
			args: args{
				document: "111.111.111-11",
			},
			want: false,
		},
		{
			name: "Invalid data as document 935.411.4347-81020 - Should return false",
			args: args{
				document: "935.411.4347-81020",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidDocumentNumber(tt.args.document); got != tt.want {
				t.Errorf("IsValidDocumentNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
