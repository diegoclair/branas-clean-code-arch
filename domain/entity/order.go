package entity

import (
	"errors"

	"github.com/diegoclair/branas-clean-code-arch/validator"
)

type Order struct {
	OrderID    int64
	Document   string
	OrderItems []OrderItem
	Coupon     Coupon
}

func NewOrder(document string) (order Order, err error) {
	if !validator.IsValidDocumentNumber(document) {
		return order, errors.New("invalid document")
	}
	order.Document = document
	return order, nil
}

func (o *Order) addItem(item Item, quantity int64) {
	o.OrderItems = append(o.OrderItems, OrderItem{ItemID: item.ItemID, Price: item.Price, Quantity: quantity})
}

func (o *Order) addCoupon(coupon Coupon) {
	o.Coupon = coupon
}

func (o *Order) getTotal() (total float64) {
	for _, orderItem := range o.OrderItems {
		total += orderItem.getTotal()
	}
	if o.Coupon.Percentage > 0 {
		total -= total * (float64(o.Coupon.Percentage) / 100)
	}
	return total
}

type OrderItem struct {
	ItemID   int64
	Quantity int64
	Price    float64
}

func (oi *OrderItem) getTotal() float64 {
	return oi.Price * float64(oi.Quantity)
}
