package repository

import "github.com/diegoclair/branas-clean-code-arch/domain/entity"

type ItemRepository interface {
	FindByID(id int64) (entity.Item, error)
}
