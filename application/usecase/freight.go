package usecase

import (
	"github.com/diegoclair/branas-clean-code-arch/application/dto"
	"github.com/diegoclair/branas-clean-code-arch/domain/contract"
	"github.com/diegoclair/branas-clean-code-arch/domain/service"
)

type freightUsecase struct {
	itemRepo contract.ItemRepository
}

func NewFreightUsecase(itemRepo contract.ItemRepository) *freightUsecase {
	return &freightUsecase{
		itemRepo: itemRepo,
	}
}

func (u *freightUsecase) FreightSimulation(input dto.FreightSimulationInput) (freight float64, err error) {

	for _, inputItem := range input.Items {

		item, err := u.itemRepo.FindByID(inputItem.ItemID)
		if err != nil {
			return freight, err
		}
		freight += service.FreightCalculator(item) * float64(inputItem.Quantity)
	}
	return freight, nil
}
