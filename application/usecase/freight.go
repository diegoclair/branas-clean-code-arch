package usecase

import (
	"github.com/diegoclair/branas-clean-code-arch/application/dto"
	"github.com/diegoclair/branas-clean-code-arch/domain/repository"
)

type newFreight struct {
	itemRepo repository.ItemRepository
}

func NewFreight(itemRepo repository.ItemRepository) *newFreight {
	return &newFreight{
		itemRepo: itemRepo,
	}
}

func (u *newFreight) Execute(input dto.FreightSimulationInput) (freight float64, err error) {

	for _, inputItem := range input.Items {

		item, err := u.itemRepo.FindByID(inputItem.ItemID)
		if err != nil {
			return freight, err
		}
		freight += item.GetFreight() * float64(inputItem.Quantity)
	}
	return freight, nil
}
