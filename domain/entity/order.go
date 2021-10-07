package entity

import (
	"errors"

	"github.com/diegoclair/branas-clean-code-arch/validator"
)

type Order struct {
	OrderID  int64
	Document string
	Items    []Item
}

func NewOrder(document string) (order Order, err error) {
	if !validator.IsValidDocumentNumber(document) {
		return order, errors.New("invalid document")
	}
	order.Document = document
	return order, nil
}

func (o *Order) addItem(item Item) {
	o.Items = append(o.Items, item)
}

func (o *Order) getTotal() (total float64) {
	for _, orderItem := range o.Items {
		total += orderItem.getTotal()
	}
	return total
}
