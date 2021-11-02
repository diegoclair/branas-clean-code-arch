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

func NewOrderUsecase(db contract.RepoManager) *orderUsecase {
	return &orderUsecase{
		itemRepo:   db.Item(),
		orderRepo:  db.Order(),
		couponRepo: db.Coupon(),
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

	err = u.orderRepo.Save(&order)
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

	order.OrderItems, err = u.orderRepo.GetOrderItemsByOrderID(order.OrderID)
	if err != nil {
		return orderOutput, err
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

func (u *orderUsecase) GetOrders() (ordersOutput []dto.OrderOutput, err error) {

	orders, err := u.orderRepo.GetOrders()
	if err != nil {
		return ordersOutput, err
	}

	for i := range orders {
		orders[i].OrderItems, err = u.orderRepo.GetOrderItemsByOrderID(orders[i].OrderID)
		if err != nil {
			return ordersOutput, err
		}

		if orders[i].Coupon.CouponID > 0 {
			coupon, err := u.couponRepo.FindByCode(orders[i].Coupon.Code)
			if err != nil {
				return ordersOutput, err
			}
			orders[i].AddCoupon(coupon)
		}
		orderOutput := dto.OrderOutput{}
		ordersOutput = append(ordersOutput, orderOutput.Assembly(orders[i]))
	}

	return ordersOutput, nil
}
