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

func (u *orderUsecase) PlaceOrder(input dto.CreateOrderInput) (response dto.CreateOrderOutput, err error) {

	sequence, err := u.orderRepo.Count()
	if err != nil {
		return response, err
	}
	sequence++

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

func (u *orderUsecase) GetOrderByCode(code string) (orderOutput dto.OrderOutput, err error) {

	order, err := u.orderRepo.GetByCode(code)
	if err != nil {
		return orderOutput, err
	}

	ordermItems, err := u.orderRepo.GetOrderItemsByOrderID(order.OrderID)
	if err != nil {
		return orderOutput, err
	}

	if len(ordermItems) > 0 {
		for _, ordermItem := range ordermItems {
			item, err := u.itemRepo.FindByID(ordermItem.ItemID)
			if err != nil {
				return orderOutput, err
			}
			order.AddItem(item, ordermItem.Quantity)
		}
	}

	if order.Coupon.Code != "" {
		coupon, err := u.couponRepo.FindByCode(order.Coupon.Code)
		if err != nil {
			return orderOutput, err
		}
		order.AddCoupon(coupon)
	}

	return orderOutput.Assembly(order), nil
}
