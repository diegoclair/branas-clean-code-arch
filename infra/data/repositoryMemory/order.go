package repositorymemory

import (
	"github.com/diegoclair/branas-clean-code-arch/domain/entity"
	"github.com/diegoclair/branas-clean-code-arch/domain/repository"
)

type orderRepositoryMemory struct {
	orders []entity.Order
}

func NewOrderRepositoryMemory() repository.OrderRepository {
	return &orderRepositoryMemory{}
}

func (r *orderRepositoryMemory) Save(order entity.Order) (err error) {
	r.orders = append(r.orders, order)
	return nil
}
