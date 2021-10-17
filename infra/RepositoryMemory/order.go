package repositorymemory

import (
	"github.com/diegoclair/branas-clean-code-arch/domain/entity"
	"github.com/diegoclair/branas-clean-code-arch/domain/repository"
)

type orderRepositoryMemory struct {
}

func NewOrderRepositoryMemory() repository.OrderRepository {
	return &orderRepositoryMemory{}
}

func (r *orderRepositoryMemory) Save(order entity.Order) (err error) {
	return nil
}
