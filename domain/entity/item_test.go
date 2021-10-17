package entity

import (
	"reflect"
	"testing"
)

var (
	id          int64   = 5
	quantity    int64   = 1
	category    string  = "Guitarra"
	description string  = "Instrumentos Musicais"
	price       float64 = 59.99
	width       float64 = 100
	height      float64 = 30
	length      float64 = 10
	weight      float64 = 3
)

type itemArgs struct {
	id          int64
	quantity    int64
	category    string
	description string
	price       float64
	width       float64
	height      float64
	length      float64
	weight      float64
}

func TestNewItem(t *testing.T) {

	tests := []struct {
		name     string
		args     itemArgs
		wantItem Item
	}{
		{
			name: "Should create an item",
			args: itemArgs{
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
			if gotItem := NewItem(tt.args.id, tt.args.category, tt.args.description, tt.args.price, 0, 0, 0, 0); !reflect.DeepEqual(gotItem, tt.wantItem) {
				t.Errorf("NewItem() = %v, want %v", gotItem, tt.wantItem)
			}
		})
	}
}

func TestCalculateVolume(t *testing.T) {

	tests := []struct {
		name       string
		args       itemArgs
		wantVolume float64
	}{
		{
			name: "Should create an item and get it volume",
			args: itemArgs{
				id:          id,
				quantity:    quantity,
				category:    category,
				description: description,
				price:       price,
				width:       width,
				height:      height,
				length:      length,
				weight:      weight,
			},
			wantVolume: 0.03,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotItem := NewItem(tt.args.id, tt.args.category, tt.args.description, tt.args.price, tt.args.width, tt.args.height, tt.args.length, tt.args.weight)
			gotVolume := gotItem.getVolume()
			if gotVolume != tt.wantVolume {
				t.Errorf("getVolume() = %v, want %v", gotVolume, tt.wantVolume)
			}
		})
	}
}

func TestCalculateDensity(t *testing.T) {

	tests := []struct {
		name        string
		args        itemArgs
		wantDensity float64
	}{
		{
			name: "Should create an item and get it density",
			args: itemArgs{
				id:          id,
				quantity:    quantity,
				category:    category,
				description: description,
				price:       price,
				width:       width,
				height:      height,
				length:      length,
				weight:      weight,
			},
			wantDensity: 100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotItem := NewItem(tt.args.id, tt.args.category, tt.args.description, tt.args.price, tt.args.width, tt.args.height, tt.args.length, tt.args.weight)
			gotDensity := gotItem.getDensity()
			if gotDensity != tt.wantDensity {
				t.Errorf("getDensity() = %v, want %v", gotDensity, tt.wantDensity)
			}
		})
	}
}

func TestCalculateFreight(t *testing.T) {

	tests := []struct {
		name        string
		args        itemArgs
		wantFreight float64
	}{
		{
			name: "Should create an item and get it freight",
			args: itemArgs{
				id:          id,
				quantity:    quantity,
				category:    category,
				description: description,
				price:       price,
				width:       width,
				height:      height,
				length:      length,
				weight:      weight,
			},
			wantFreight: 30,
		},
		{
			name: "Should create an item and get the minimum freight",
			args: itemArgs{
				id:          id,
				quantity:    quantity,
				category:    category,
				description: description,
				price:       30,
				width:       10,
				height:      10,
				length:      10,
				weight:      0.9,
			},
			wantFreight: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotItem := NewItem(tt.args.id, tt.args.category, tt.args.description, tt.args.price, tt.args.width, tt.args.height, tt.args.length, tt.args.weight)
			gotFreight := gotItem.getFreight()
			if gotFreight != tt.wantFreight {
				t.Errorf("getFreight() = %v, want %v", gotFreight, tt.wantFreight)
			}
		})
	}
}
