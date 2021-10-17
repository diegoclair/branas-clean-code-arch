package repository

import "github.com/diegoclair/branas-clean-code-arch/domain/entity"

type OrderRepository interface {
	Save(entity.Order) error
}
