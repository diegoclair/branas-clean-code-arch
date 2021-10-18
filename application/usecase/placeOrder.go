package usecase

import (
	"github.com/diegoclair/branas-clean-code-arch/domain/entity"
	"github.com/diegoclair/branas-clean-code-arch/domain/repository"
)

type modelOrderItem struct {
	ID       int64
	Quantity int64
}
type modelOrder struct {
	Cpf        string
	OrderItems []modelOrderItem
}
type outputOrder struct {
	Total float64
}

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

func (u *newPlaceOrder) Execute(input modelOrder) (output outputOrder, err error) {

	order, err := entity.NewOrder(input.Cpf)
	if err != nil {
		return output, err
	}
	for _, orderItem := range input.OrderItems {

		item, err := u.itemRepo.FindByID(orderItem.ID)
		if err != nil {
			return output, err
		}
		order.AddItem(item, orderItem.Quantity)
	}

	err = u.orderRepo.Save(order)
	if err != nil {
		return output, err
	}

	output.Total = order.GetTotal()
	return output, nil
}
