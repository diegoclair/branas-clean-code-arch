package main

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

func Test_calculateCPFDigit(t *testing.T) {
	type args struct {
		document string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "should return 8 for first digit",
			args: args{
				document: "935411347",
			},
			want: 8,
		},
		{
			name: "should return 0 for second digit",
			args: args{
				document: "9354113478",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateCPFDigit(tt.args.document); got != tt.want {
				t.Errorf("calculateCPFDigit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_cleanNumber(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Should return 0123456789",
			args: args{
				value: "er01.2-34y5.67-89",
			},
			want: "0123456789",
		},
		{
			name: "Should return 11111111111",
			args: args{
				value: "111.111.111-11",
			},
			want: "11111111111",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cleanNumber(tt.args.value); got != tt.want {
				t.Errorf("cleanNumber() = %v, want %v", got, tt.want)
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
