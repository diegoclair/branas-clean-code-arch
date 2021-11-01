package dto

import "github.com/diegoclair/branas-clean-code-arch/domain/entity"

type CreateOrderInput struct {
	Cpf        string
	Coupon     string
	OrderItems []OrderItems
}

type OrderItems struct {
	Quantity int64
	Item
}

type Item struct {
	ItemID      int64
	Description string
	Price       float64
}

type CreateOrderOutput struct {
	Total     float64
	OrderCode string
}

func (o CreateOrderOutput) Assembly(order entity.Order) CreateOrderOutput {
	o.Total = order.Total
	o.OrderCode = order.Code
	return o
}

type OrderOutput struct {
	OrderCode  string
	Total      float64
	OrderItems []OrderItems
}

func (o OrderOutput) Assembly(order entity.Order) OrderOutput {
	o.OrderCode = order.Code
	for _, orderItem := range order.OrderItems {
		orderItemOutput := OrderItems{
			Quantity: orderItem.Quantity,
			Item: Item{
				ItemID:      orderItem.ItemID,
				Description: orderItem.Description,
				Price:       orderItem.Price,
			},
		}
		o.OrderItems = append(o.OrderItems, orderItemOutput)
	}
	o.Total = order.Total
	return o
}
