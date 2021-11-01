package usecase

import (
	"github.com/diegoclair/branas-clean-code-arch/application/dto"
	"github.com/diegoclair/branas-clean-code-arch/domain/contract"
	"github.com/diegoclair/branas-clean-code-arch/domain/service"
)

type newFreight struct {
	itemRepo contract.ItemRepository
}

func NewFreight(itemRepo contract.ItemRepository) *newFreight {
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
		freight += service.FreightCalculator(item) * float64(inputItem.Quantity)
	}
	return freight, nil
}
