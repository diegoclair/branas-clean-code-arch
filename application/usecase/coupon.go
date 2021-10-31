package usecase

import "github.com/diegoclair/branas-clean-code-arch/domain/repository"

type newCoupon struct {
	couponRepo repository.CouponRepository
}

func NewCoupon(couponRepo repository.CouponRepository) *newCoupon {
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
