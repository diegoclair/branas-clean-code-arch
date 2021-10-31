package repository

import "github.com/diegoclair/branas-clean-code-arch/domain/entity"

type CouponRepository interface {
	FindByCode(code string) (entity.Coupon, error)
}
