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
)

func TestNewItem(t *testing.T) {

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
			if gotItem := NewItem(tt.args.id, tt.args.category, tt.args.description, tt.args.price, 0, 0, 0); !reflect.DeepEqual(gotItem, tt.wantItem) {
				t.Errorf("NewItem() = %v, want %v", gotItem, tt.wantItem)
			}
		})
	}
}

func TestCalculateVolume(t *testing.T) {

	type args struct {
		id          int64
		quantity    int64
		category    string
		description string
		price       float64
		width       float64
		height      float64
		length      float64
	}
	tests := []struct {
		name       string
		args       args
		wantVolume float64
	}{
		{
			name: "Should create an item and get it volume",
			args: args{
				id:          id,
				quantity:    quantity,
				category:    category,
				description: description,
				price:       price,
				width:       width,
				height:      height,
				length:      length,
			},
			wantVolume: 0.03,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotItem := NewItem(tt.args.id, tt.args.category, tt.args.description, tt.args.price, tt.args.width, tt.args.height, tt.args.length)
			gotVolume := gotItem.getVolume()
			if gotVolume != tt.wantVolume {
				t.Errorf("getVolume() = %v, want %v", gotVolume, tt.wantVolume)
			}
		})
	}
}
