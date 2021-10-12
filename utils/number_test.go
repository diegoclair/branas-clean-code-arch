package utils

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
			if got := CleanNumber(tt.args.value); got != tt.want {
				t.Errorf("cleanNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Round(t *testing.T) {
	type args struct {
		f      float64
		places int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Rounds zero to zero when using 2 places",
			args: args{
				f:      0,
				places: 2,
			},
			want: 0,
		},
		{
			name: "Does not round an integer when using 2 places",
			args: args{
				f:      1,
				places: 2,
			},
			want: 1,
		},
		{
			name: "Rounds 1 decimal place to an integer when using 0 place",
			args: args{
				f:      1.1,
				places: 0,
			},
			want: 1,
		},
		{
			name: "Does not round 1 decimal place when using 1 place",
			args: args{
				f:      1.1,
				places: 1,
			},
			want: 1.1,
		},
		{
			name: "Does not round 1 decimal place when using 2 places",
			args: args{
				f:      1.1,
				places: 2,
			},
			want: 1.1,
		},
		{
			name: "Rounds up when the the last decimal digit is 5 or greater",
			args: args{
				f:      1.155,
				places: 2,
			},
			want: 1.16,
		},
		{
			name: "Rounds down when the the last decimal digit is less then 5",
			args: args{
				f:      1.153,
				places: 2,
			},
			want: 1.15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Round(tt.args.f, tt.args.places); got != tt.want {
				t.Errorf("Round() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_RoundUp(t *testing.T) {
	type args struct {
		f      float64
		places int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Rounds zero to zero when using 2 places",
			args: args{
				f:      0,
				places: 2,
			},
			want: 0,
		},
		{
			name: "Does not round integer when using 2 places",
			args: args{
				f:      1,
				places: 2,
			},
			want: 1,
		},
		{
			name: "Does not round 1 decimal place when using 1 place",
			args: args{
				f:      1.1,
				places: 1,
			},
			want: 1.1,
		},
		{
			name: "Rounds up 1 decimal place to an integer when using 0 place",
			args: args{
				f:      1.1,
				places: 0,
			},
			want: 2,
		},
		{
			name: "Rounds 1 decimal place to 1 decimal place when using 2 places",
			args: args{
				f:      1.1,
				places: 2,
			},
			want: 1.1,
		},
		{
			name: "Rounds up when the the last decimal digit is 5 or greater",
			args: args{
				f:      1.155,
				places: 2,
			},
			want: 1.16,
		},
		{
			name: "Rounds up when the the last decimal digit is less then 5",
			args: args{
				f:      1.153,
				places: 2,
			},
			want: 1.16,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RoundUp(tt.args.f, tt.args.places); got != tt.want {
				t.Errorf("RoundUp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_RoundDown(t *testing.T) {
	type args struct {
		f      float64
		places int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Rounds zero to zero when using 2 places",
			args: args{
				f:      0,
				places: 2,
			},
			want: 0,
		},
		{
			name: "Does not round integer when using 2 places",
			args: args{
				f:      1,
				places: 2,
			},
			want: 1,
		},
		{
			name: "Does not round 1 decimal place when using 1 place",
			args: args{
				f:      1.1,
				places: 1,
			},
			want: 1.1,
		},
		{
			name: "Rounds up 1 decimal place to an integer when using 0 place",
			args: args{
				f:      1.1,
				places: 0,
			},
			want: 1,
		},
		{
			name: "Rounds 1 decimal place to 1 decimal place when using 2 places",
			args: args{
				f:      1.1,
				places: 2,
			},
			want: 1.1,
		},
		{
			name: "Rounds up when the the last decimal digit is 5 or greater",
			args: args{
				f:      1.155,
				places: 2,
			},
			want: 1.15,
		},
		{
			name: "Rounds up when the the last decimal digit is less then 5",
			args: args{
				f:      1.153,
				places: 2,
			},
			want: 1.15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RoundDown(tt.args.f, tt.args.places); got != tt.want {
				t.Errorf("RoundDown() = %v, want %v", got, tt.want)
			}
		})
	}
}
