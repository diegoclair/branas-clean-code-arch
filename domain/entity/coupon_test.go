package entity

import (
	"reflect"
	"testing"
)

func TestNewCoupon(t *testing.T) {
	var (
		code       string = "VALE20"
		percentage int64  = 20
	)
	type args struct {
		code       string
		percentage int64
	}
	tests := []struct {
		name string
		args args
		want Coupon
	}{
		{
			name: "Should create a coupon",
			args: args{
				code:       code,
				percentage: percentage,
			},
			want: Coupon{
				Code:       code,
				Percentage: percentage,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCoupon(tt.args.code, tt.args.percentage); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCoupon() = %v, want %v", got, tt.want)
			}
		})
	}
}
