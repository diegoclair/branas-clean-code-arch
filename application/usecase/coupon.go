package usecase

import "github.com/diegoclair/branas-clean-code-arch/domain/contract"

type couponUsecase struct {
	couponRepo contract.CouponRepository
}

func newCouponUsecase(us *Usecase) *couponUsecase {
	return &couponUsecase{
		couponRepo: us.db.Coupon(),
	}
}

func (u *couponUsecase) CouponValidation(code string) (isValid bool, err error) {

	coupon, err := u.couponRepo.FindByCode(code)
	if err != nil {
		return isValid, err
	}
	return coupon.IsValid(), nil
}
