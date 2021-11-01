package usecase

import (
	"strconv"
	"testing"
	"time"

	"github.com/diegoclair/branas-clean-code-arch/application/dto"
	"github.com/diegoclair/branas-clean-code-arch/domain/contract"
	"github.com/diegoclair/branas-clean-code-arch/domain/entity"
	repositorymemory "github.com/diegoclair/branas-clean-code-arch/infra/data/repositoryMemory"
)

func TestPlaceOrder(t *testing.T) {

	itemRepo := repositorymemory.NewItemRepositoryMemory()
	orderRepo := repositorymemory.NewOrderRepositoryMemory()
	couponRepo := repositorymemory.NewCouponRepositoryMemory()

	tests := []struct {
		name      string
		args      dto.CreateOrderInput
		wantTotal float64
		wantErr   bool
	}{
		{
			name: "Shoud place an order with 3 items",
			args: dto.CreateOrderInput{
				Cpf: "847.903.332-05",
				OrderItems: []dto.OrderItems{
					{
						Item:     dto.Item{ItemID: 1},
						Quantity: 1,
					},
					{
						Item:     dto.Item{ItemID: 2},
						Quantity: 1,
					},
					{
						Item:     dto.Item{ItemID: 3},
						Quantity: 3,
					},
				},
			},
			wantTotal: 5468.99,
			wantErr:   false,
		},
		{
			name: "Shoud place an order with 3 items and a coupon",
			args: dto.CreateOrderInput{
				Cpf:    "847.903.332-05",
				Coupon: "VALE20",
				OrderItems: []dto.OrderItems{
					{
						Item:     dto.Item{ItemID: 1},
						Quantity: 1,
					},
					{
						Item:     dto.Item{ItemID: 2},
						Quantity: 1,
					},
					{
						Item:     dto.Item{ItemID: 3},
						Quantity: 3,
					},
				},
			},
			wantTotal: 4375.19,
			wantErr:   false,
		},
		{
			name: "Shoud get an error with non exists item id",
			args: dto.CreateOrderInput{
				Cpf: "847.903.332-05",
				OrderItems: []dto.OrderItems{
					{
						Item:     dto.Item{ItemID: 55},
						Quantity: 1,
					},
				},
			},
			wantTotal: 0,
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &orderUsecase{
				itemRepo:   itemRepo,
				orderRepo:  orderRepo,
				couponRepo: couponRepo,
			}
			got, err := u.PlaceOrder(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("got err = %v, want err? %v", err, tt.wantErr)
			}
			if got.Total != tt.wantTotal {
				t.Errorf("got total = %v, want %v", got.Total, tt.wantTotal)
			}

		})
	}
}

func TestGetOrderByCode(t *testing.T) {

	orderRepo := repositorymemory.NewOrderRepositoryMemory()
	itemRepo := repositorymemory.NewItemRepositoryMemory()
	couponRepo := repositorymemory.NewCouponRepositoryMemory()

	code := prepareOrderInput(orderRepo, itemRepo, couponRepo)
	type args struct {
		orderRepo  contract.OrderRepository
		itemRepo   contract.ItemRepository
		couponRepo contract.CouponRepository
		code       string
	}
	tests := []struct {
		name      string
		args      args
		wantTotal float64
		wantErr   bool
	}{
		{
			name: "Should get an order by it code",
			args: args{
				orderRepo:  orderRepo,
				itemRepo:   itemRepo,
				couponRepo: couponRepo,
				code:       code,
			},
			wantTotal: 4375.19,
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &orderUsecase{
				itemRepo:   itemRepo,
				orderRepo:  orderRepo,
				couponRepo: couponRepo,
			}
			got, err := u.GetOrderByCode(tt.args.code)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetOrderByCode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got.Total != tt.wantTotal {
				t.Errorf("GetOrderByCode().Total = %v, wantCode %v", got.Total, tt.wantTotal)
			}
		})
	}
}

func prepareOrderInput(orderRepo contract.OrderRepository, itemRepo contract.ItemRepository, couponRepo contract.CouponRepository) (code string) {
	orderInput := dto.CreateOrderInput{
		Cpf:    "847.903.332-05",
		Coupon: "VALE20",
		OrderItems: []dto.OrderItems{
			{
				Item:     dto.Item{ItemID: 1},
				Quantity: 1,
			},
			{
				Item:     dto.Item{ItemID: 2},
				Quantity: 1,
			},
			{
				Item:     dto.Item{ItemID: 3},
				Quantity: 3,
			},
		},
	}
	const sequence = 1
	order, _ := entity.NewOrder(orderInput.Cpf, sequence)
	for _, orderItem := range orderInput.OrderItems {
		item, _ := itemRepo.FindByID(orderItem.ItemID)
		order.AddItem(item, orderItem.Quantity)
	}
	coupon, _ := couponRepo.FindByCode(orderInput.Coupon)
	order.AddCoupon(coupon)
	orderRepo.Save(&order)

	y, _, _ := time.Now().Date()
	year := strconv.Itoa(y)
	code = year + "0000000" + strconv.Itoa(sequence)
	return code
}
