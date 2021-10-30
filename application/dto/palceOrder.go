package dto

import "github.com/diegoclair/branas-clean-code-arch/domain/entity"

type OrderItemInput struct {
	Cpf        string
	OrderItems []OrderItems
}

type OrderItems struct {
	ItemID   int64
	Quantity int64
}

type OrderItemOutput struct {
	Total float64
}

func (o OrderItemOutput) Assembly(order entity.Order) OrderItemOutput {
	o.Total = order.GetTotal()
	return o
}
