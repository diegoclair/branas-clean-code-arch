package usecase

import (
	"testing"

	"github.com/diegoclair/branas-clean-code-arch/application/dto"
	repositorymemory "github.com/diegoclair/branas-clean-code-arch/infra/data/repositoryMemory"
)

func TestSimulateFreight(t *testing.T) {

	itemRepo := repositorymemory.New().Item()

	tests := []struct {
		name        string
		args        dto.FreightSimulationInput
		wantFreight float64
		wantErr     bool
	}{
		{
			name: "Shoud simulate products freight",
			args: dto.FreightSimulationInput{
				Items: []dto.OrderItems{
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
			wantFreight: 260,
			wantErr:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &freightUsecase{
				itemRepo: itemRepo,
			}
			got, err := u.FreightSimulation(tt.args)
			if err == nil && tt.wantErr || err != nil && !tt.wantErr {
				t.Errorf("got err = %v, want err? %v", err, tt.wantErr)
			}
			if got != tt.wantFreight {
				t.Errorf("got freight = %v, want %v", got, tt.wantFreight)
			}

		})
	}
}
