package repositorymemory

import (
	"github.com/diegoclair/branas-clean-code-arch/domain/contract"
	"github.com/diegoclair/branas-clean-code-arch/domain/entity"
)

type orderRepositoryMemory struct {
	orders []entity.Order
}

func NewOrderRepositoryMemory() contract.OrderRepository {
	return &orderRepositoryMemory{}
}

func (r *orderRepositoryMemory) Save(order entity.Order) (err error) {
	r.orders = append(r.orders, order)
	return nil
}
