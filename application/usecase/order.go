package usecase

import (
	"github.com/diegoclair/branas-clean-code-arch/application/dto"
	"github.com/diegoclair/branas-clean-code-arch/domain/contract"
	"github.com/diegoclair/branas-clean-code-arch/domain/entity"
)

type orderUsecase struct {
	itemRepo   contract.ItemRepository
	orderRepo  contract.OrderRepository
	couponRepo contract.CouponRepository
}

func NewOrderUsecase(itemRepo contract.ItemRepository, orderRepo contract.OrderRepository, couponRepo contract.CouponRepository) *orderUsecase {
	return &orderUsecase{
		itemRepo:   itemRepo,
		orderRepo:  orderRepo,
		couponRepo: couponRepo,
	}
}

func (u *orderUsecase) PlaceOrder(input dto.OrderInput) (response dto.OrderOutput, err error) {

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

func (u *orderUsecase) GetOrder() (order entity.Order, err error) {

	return order, nil
}
