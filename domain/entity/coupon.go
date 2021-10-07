package entity

type Coupon struct {
	Code       string
	Percentage int64
}

func NewCoupon(code string, percentage int64) Coupon {
	return Coupon{
		Code:       code,
		Percentage: percentage,
	}
}
