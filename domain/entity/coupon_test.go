package entity

import (
	"reflect"
	"testing"
	"time"
)

var (
	expiredCoupon = NewCoupon("VALE20", 20, time.Date(2021, time.April, 2, 0, 0, 0, 0, time.Local))
)

func TestNewCoupon(t *testing.T) {
	var (
		code           string = "VALE20"
		percentage     int64  = 20
		expirationDate time.Time
	)
	type args struct {
		code           string
		percentage     int64
		expirationDate time.Time
	}
	tests := []struct {
		name string
		args args
		want Coupon
	}{
		{
			name: "Should create a coupon",
			args: args{
				code:           code,
				percentage:     percentage,
				expirationDate: expirationDate,
			},
			want: Coupon{
				Code:           code,
				Percentage:     percentage,
				ExpirationDate: expirationDate,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCoupon(tt.args.code, tt.args.percentage, tt.args.expirationDate); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCoupon() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckCouponExpiration(t *testing.T) {
	type args struct {
		code           string
		percentage     int64
		expirationDate time.Time
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Should get error with an expired coupon",
			args: args{
				code:           expiredCoupon.Code,
				percentage:     expiredCoupon.Percentage,
				expirationDate: expiredCoupon.ExpirationDate,
			},
			want: true,
		},
		{
			name: "Validate coupon that is not expired",
			args: args{
				code:           expiredCoupon.Code,
				percentage:     expiredCoupon.Percentage,
				expirationDate: time.Now().Add(24 * time.Hour),
			},
			want: false,
		},
		{
			name: "Validate coupon that never expires",
			args: args{
				code:           expiredCoupon.Code,
				percentage:     expiredCoupon.Percentage,
				expirationDate: time.Time{},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewCoupon(tt.args.code, tt.args.percentage, tt.args.expirationDate)
			if c.isExpired() != tt.want {
				t.Errorf("got %v - want %v", c.isExpired(), tt.want)
			}
		})
	}
}
