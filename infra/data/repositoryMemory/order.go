package repositorymemory

import (
	"errors"

	"github.com/diegoclair/branas-clean-code-arch/domain/contract"
	"github.com/diegoclair/branas-clean-code-arch/domain/entity"
)

type orderRepositoryMemory struct {
	orders []entity.Order
}

func NewOrderRepositoryMemory() contract.OrderRepository {
	return &orderRepositoryMemory{}
}

func (r *orderRepositoryMemory) Save(order *entity.Order) (err error) {
	order.Total = order.GetTotal()
	r.orders = append(r.orders, *order)
	return nil
}

func (r *orderRepositoryMemory) GetByCode(code string) (retVal entity.Order, err error) {
	for _, order := range r.orders {
		if order.Code == code {
			order.OrderItems = []entity.OrderItem{}
			return order, nil
		}
	}

	return retVal, errors.New("order not found")
}

func (r *orderRepositoryMemory) GetOrders() (retVal []entity.Order, err error) {
	return r.orders, nil
}

func (r *orderRepositoryMemory) GetOrderItemsByOrderID(orderID int64) (orderItems []entity.OrderItem, err error) {

	for _, order := range r.orders {
		if order.OrderID == orderID {
			orderItems = append(orderItems, order.OrderItems...)
			break
		}
	}
	return orderItems, nil
}

func (r *orderRepositoryMemory) Count() (count int64, err error) {
	return int64(len(r.orders)), nil
}
