package usecase

import "github.com/diegoclair/branas-clean-code-arch/domain/contract"

type newCoupon struct {
	couponRepo contract.CouponRepository
}

func NewCoupon(couponRepo contract.CouponRepository) *newCoupon {
	return &newCoupon{
		couponRepo: couponRepo,
	}
}

func (u *newCoupon) CouponValidation(code string) (isValid bool, err error) {

	coupon, err := u.couponRepo.FindByCode(code)
	if err != nil {
		return isValid, err
	}
	return coupon.IsValid(), nil
}
