package usecase

import (
	"github.com/diegoclair/branas-clean-code-arch/application/dto"
	"github.com/diegoclair/branas-clean-code-arch/domain/entity"
	"github.com/diegoclair/branas-clean-code-arch/domain/repository"
)

type newPlaceOrder struct {
	itemRepo   repository.ItemRepository
	orderRepo  repository.OrderRepository
	couponRepo repository.CouponRepository
}

func NewPlaceOrder(itemRepo repository.ItemRepository, orderRepo repository.OrderRepository, couponRepo repository.CouponRepository) *newPlaceOrder {
	return &newPlaceOrder{
		itemRepo:   itemRepo,
		orderRepo:  orderRepo,
		couponRepo: couponRepo,
	}
}

func (u *newPlaceOrder) Execute(input dto.OrderInput) (response dto.OrderOutput, err error) {

	var sequence int64 = 1
	order, err := entity.NewOrder(input.Cpf, sequence)
	if err != nil {
		return response, err
	}
	if input.Coupon != "" {
		coupon, err := u.couponRepo.FindByCode(input.Coupon)
		if err != nil {
			return response, err
		}
		err = order.AddCoupon(coupon)
		if err != nil {
			return response, err
		}
	}

	for _, orderItem := range input.OrderItems {

		item, err := u.itemRepo.FindByID(orderItem.ItemID)
		if err != nil {
			return response, err
		}
		order.AddItem(item, orderItem.Quantity)
	}

	err = u.orderRepo.Save(order)
	if err != nil {
		return response, err
	}

	return response.Assembly(order), nil
}
