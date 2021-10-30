package usecase

import (
	"github.com/diegoclair/branas-clean-code-arch/application/dto"
	"github.com/diegoclair/branas-clean-code-arch/domain/entity"
	"github.com/diegoclair/branas-clean-code-arch/domain/repository"
)

type newPlaceOrder struct {
	itemRepo  repository.ItemRepository
	orderRepo repository.OrderRepository
}

func NewPlaceOrder(itemRepo repository.ItemRepository, orderRepo repository.OrderRepository) *newPlaceOrder {
	return &newPlaceOrder{
		itemRepo:  itemRepo,
		orderRepo: orderRepo,
	}
}

func (u *newPlaceOrder) Execute(input dto.OrderItemInput) (response dto.OrderItemOutput, err error) {

	order, err := entity.NewOrder(input.Cpf)
	if err != nil {
		return response, err
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
