package contract

import "github.com/diegoclair/branas-clean-code-arch/domain/entity"

type CouponRepository interface {
	FindByCode(code string) (entity.Coupon, error)
}

type ItemRepository interface {
	FindByID(id int64) (entity.Item, error)
}

type OrderRepository interface {
	Save(*entity.Order) error
	GetByCode(code string) (entity.Order, error)
	GetOrderItemsByOrderID(orderID int64) ([]entity.OrderItem, error)
	Count() (int64, error)
}
