package usecase

import (
	"github.com/diegoclair/branas-clean-code-arch/application/dto"
	"github.com/diegoclair/branas-clean-code-arch/domain/contract"
	"github.com/diegoclair/branas-clean-code-arch/domain/service"
)

type freightUsecase struct {
	itemRepo contract.ItemRepository
}

func NewFreightUsecase(db contract.RepoManager) *freightUsecase {
	return &freightUsecase{
		itemRepo: db.Item(),
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
