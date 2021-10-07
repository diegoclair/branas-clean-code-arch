package entity

import (
	"errors"

	"github.com/diegoclair/branas-clean-code-arch/validator"
)

type Order struct {
	OrderID  int64
	Document string
}

func NewOrder(document string) (order Order, err error) {
	if !validator.IsValidDocumentNumber(document) {
		return order, errors.New("invalid document")
	}
	order.Document = document
	return order, nil
}
