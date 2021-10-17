package repositorymemory

import (
	"errors"

	"github.com/diegoclair/branas-clean-code-arch/domain/entity"
	"github.com/diegoclair/branas-clean-code-arch/domain/repository"
)

type itemRepositoryMemory struct {
	items []entity.Item
}

func NewItemRepositoryMemory() repository.ItemRepository {
	return &itemRepositoryMemory{
		items: []entity.Item{
			entity.NewItem(1, "Instrumentos Musicais", "Guitarra", 1119, 100, 30, 10, 3),
			entity.NewItem(2, "Instrumentos Musicais", "Amplificador", 4259.99, 100, 50, 50, 20),
			entity.NewItem(3, "Instrumentos Musicais", "Cabo", 30, 10, 10, 10, 0.9),
		},
	}
}

func (r *itemRepositoryMemory) FindByID(id int64) (item entity.Item, err error) {
	for _, item := range r.items {
		if item.ItemID == id {
			return item, nil
		}
	}
	return item, errors.New("item not found")
}
