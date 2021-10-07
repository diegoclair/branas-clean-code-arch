package entity

import (
	"errors"

	"github.com/diegoclair/branas-clean-code-arch/validator"
)

type Order struct {
	OrderID    int64
	Document   string
	TotalPrice float64
	Items      []Item
}

func NewOrder(document string) (order Order, err error) {
	if !validator.IsValidDocumentNumber(document) {
		return order, errors.New("invalid document")
	}
	order.Document = document
	return order, nil
}

func (o *Order) addItem(item Item) {
	o.TotalPrice += item.Price * float64(item.Quantity)
	o.Items = append(o.Items, item)
}

func (o *Order) getTotal() float64 {
	return o.TotalPrice
}
