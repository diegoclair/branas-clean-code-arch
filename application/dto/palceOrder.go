package dto

import "github.com/diegoclair/branas-clean-code-arch/domain/entity"

type OrderInput struct {
	Cpf        string
	Coupon     string
	OrderItems []OrderItems
}

type OrderItems struct {
	ItemID   int64
	Quantity int64
}

type OrderOutput struct {
	Total     float64
	OrderCode string
}

func (o OrderOutput) Assembly(order entity.Order) OrderOutput {
	o.Total = order.GetTotal()
	o.OrderCode = order.Code
	return o
}
