package usecase

import (
	"github.com/diegoclair/branas-clean-code-arch/application/dto"
	"github.com/diegoclair/branas-clean-code-arch/domain/contract"
)

type Usecase struct {
	db contract.RepoManager
}

func New(db contract.RepoManager) *Usecase {
	return &Usecase{
		db: db,
	}
}

type Manager interface {
	OrderUsecase(us *Usecase, freightUsecase FreightUsecase) OrderUsecase
	FreightUsecase(us *Usecase) FreightUsecase
	CouponUsecase(us *Usecase) CouponUsecase
}

type OrderUsecase interface {
	GetOrderByCode(code string) (orderOutput dto.OrderOutput, err error)
	GetOrders() (ordersOutput []dto.OrderOutput, err error)
	PlaceOrder(input dto.CreateOrderInput) (response dto.CreateOrderOutput, err error)
}

type FreightUsecase interface {
	FreightSimulation(input dto.FreightSimulationInput) (freight float64, err error)
}

type CouponUsecase interface {
	CouponValidation(code string) (isValid bool, err error)
}

type usecaseManager struct {
}

func NewUsecaseManager() Manager {
	return &usecaseManager{}
}

func (u *usecaseManager) OrderUsecase(us *Usecase, freightUsecase FreightUsecase) OrderUsecase {
	return newOrderUsecase(us, freightUsecase)
}

func (u *usecaseManager) FreightUsecase(us *Usecase) FreightUsecase {
	return newFreightUsecase(us)
}

func (u *usecaseManager) CouponUsecase(us *Usecase) CouponUsecase {
	return newCouponUsecase(us)
}
