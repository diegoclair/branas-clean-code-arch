package entity

import (
	"reflect"
	"testing"
)

func TestNewItem(t *testing.T) {
	var (
		id          int64   = 5
		quantity    int64   = 1
		category    string  = "Guitarra"
		description string  = "Instrumentos Musicais"
		price       float64 = 59.99
	)

	type args struct {
		id          int64
		quantity    int64
		category    string
		description string
		price       float64
	}
	tests := []struct {
		name     string
		args     args
		wantItem Item
	}{
		{
			name: "Should create an item",
			args: args{
				id:          id,
				quantity:    quantity,
				category:    category,
				description: description,
				price:       price,
			},
			wantItem: Item{
				ItemID:      id,
				Category:    category,
				Description: description,
				Price:       price,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotItem := NewItem(tt.args.id, tt.args.category, tt.args.description, tt.args.price); !reflect.DeepEqual(gotItem, tt.wantItem) {
				t.Errorf("NewItem() = %v, want %v", gotItem, tt.wantItem)
			}
		})
	}
}
