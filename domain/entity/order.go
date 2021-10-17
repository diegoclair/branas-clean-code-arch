package entity

import (
	"errors"

	"github.com/diegoclair/branas-clean-code-arch/utils"
)

type Order struct {
	OrderID    int64
	Document   string
	OrderItems []OrderItem
	Coupon     Coupon
	freight    float64
}

func NewOrder(document string) (order Order, err error) {
	if !utils.IsValidDocumentNumber(document) {
		return order, errors.New("invalid document")
	}
	order.Document = document
	return order, nil
}

func (o *Order) addItem(item Item, quantity int64) {
	o.freight += item.getFreight() * float64(quantity)
	o.OrderItems = append(o.OrderItems, OrderItem{ItemID: item.ItemID, Price: item.Price, Quantity: quantity})
}

func (o *Order) addCoupon(coupon Coupon) error {
	if coupon.isExpired() {
		return errors.New("coupon is expired")
	}
	o.Coupon = coupon
	return nil
}

func (o *Order) getTotal() (total float64) {
	for _, orderItem := range o.OrderItems {
		total += orderItem.getTotal()
	}
	if o.Coupon.Percentage > 0 {
		total -= total * (float64(o.Coupon.Percentage) / 100)
	}
	return utils.Round(total, 2)
}

func (o *Order) getFreight() float64 {
	return o.freight
}

type OrderItem struct {
	ItemID   int64
	Quantity int64
	Price    float64
}

func NewOrderItem(id, quantity int64, price float64) OrderItem {
	return OrderItem{
		ItemID:   id,
		Quantity: quantity,
		Price:    price,
	}
}

func (oi *OrderItem) getTotal() float64 {
	return oi.Price * float64(oi.Quantity)
}
