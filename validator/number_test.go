package validator

import "testing"

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
