package usecase

import (
	"testing"

	repositorymemory "github.com/diegoclair/branas-clean-code-arch/infra/data/repositoryMemory"
)

func TestPlaceOrder(t *testing.T) {

	tests := []struct {
		name      string
		args      modelOrder
		wantTotal float64
		wantErr   bool
	}{
		{
			name: "Shoud place an order with 3 items",
			args: modelOrder{
				Cpf: "847.903.332-05",
				OrderItems: []modelOrderItem{
					{
						ID:       1,
						Quantity: 1,
					},
					{
						ID:       2,
						Quantity: 1,
					},
					{
						ID:       3,
						Quantity: 3,
					},
				},
			},
			wantTotal: 5468.99,
			wantErr:   false,
		},
		{
			name: "Shoud get an error with non exists item id",
			args: modelOrder{
				Cpf: "847.903.332-05",
				OrderItems: []modelOrderItem{
					{
						ID:       55,
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
			itemRepo := repositorymemory.NewItemRepositoryMemory()
			orderRepo := repositorymemory.NewOrderRepositoryMemory()
			newPlaceOrder := NewPlaceOrder(itemRepo, orderRepo)
			got, err := newPlaceOrder.Execute(tt.args)
			if err == nil && tt.wantErr || err != nil && !tt.wantErr {
				t.Errorf("got err = %v, want err? %v", err, tt.wantErr)
			}
			if got.Total != tt.wantTotal {
				t.Errorf("got total = %v, want %v", got.Total, tt.wantTotal)
			}

		})
	}
}
