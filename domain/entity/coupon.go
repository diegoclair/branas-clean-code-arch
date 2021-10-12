package entity

import "time"

type Coupon struct {
	Code           string
	Percentage     int64
	ExpirationDate time.Time
}

func NewCoupon(code string, percentage int64, expirationDate time.Time) Coupon {
	return Coupon{
		Code:           code,
		Percentage:     percentage,
		ExpirationDate: expirationDate,
	}
}

func (c Coupon) isExpired() bool {
	if c.ExpirationDate.IsZero() {
		return false
	}
	return time.Now().After(c.ExpirationDate)
}
