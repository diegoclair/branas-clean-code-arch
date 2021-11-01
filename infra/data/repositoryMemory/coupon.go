package repositorymemory

import (
	"errors"
	"time"

	"github.com/diegoclair/branas-clean-code-arch/domain/contract"
	"github.com/diegoclair/branas-clean-code-arch/domain/entity"
)

type couponRepositoryMemory struct {
	coupons []entity.Coupon
}

func NewCouponRepositoryMemory() contract.CouponRepository {
	return &couponRepositoryMemory{
		coupons: []entity.Coupon{
			{
				Code:           "VALE20",
				Percentage:     20,
				ExpirationDate: time.Now().Add(time.Hour * 999),
			},
		},
	}
}

func (r *couponRepositoryMemory) FindByCode(code string) (coupon entity.Coupon, err error) {
	for _, coupon := range r.coupons {
		if coupon.Code == code {
			return coupon, nil
		}
	}
	return coupon, errors.New("coupon not found")
}
