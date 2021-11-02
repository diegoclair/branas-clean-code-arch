package entity

import (
	"errors"
	"strconv"
	"time"

	"github.com/diegoclair/branas-clean-code-arch/utils"
)

type Order struct {
	OrderID    int64
	Code       string
	Document   string
	OrderItems []OrderItem
	Coupon     Coupon
	Freight    float64
	IssueDate  time.Time
	Total      float64
	Sequence   int64
}

func NewOrder(document string, sequence int64) (order Order, err error) {
	if !utils.IsValidDocumentNumber(document) {
		return order, errors.New("invalid document")
	}
	order.Document = document
	order.IssueDate = time.Now()
	order.Sequence = sequence
	order.GenerateCode(order.IssueDate, order.Sequence)
	return order, nil
}

func (o *Order) GenerateCode(issueDate time.Time, sequence int64) {
	y, _, _ := issueDate.Date()
	year := strconv.Itoa(y)
	seq := strconv.Itoa(int(sequence))
	o.Code = year + utils.LeftPad(seq, 8, "0")
}

func (o *Order) AddItem(item Item, quantity int64) {
	o.OrderItems = append(o.OrderItems, OrderItem{Quantity: quantity, Item: item})
}

func (o *Order) AddCoupon(coupon Coupon) error {
	if coupon.isExpired() {
		return errors.New("coupon is expired")
	}
	o.Coupon = coupon
	return nil
}

func (o *Order) GetTotal() (total float64) {
	for _, orderItem := range o.OrderItems {
		total += orderItem.GetTotal()
	}
	if o.Coupon.Percentage > 0 {
		total -= total * (float64(o.Coupon.Percentage) / 100)
	}
	o.Total = utils.Round(total, 2)
	return o.Total
}

type OrderItem struct {
	Quantity int64
	Item
}

func NewOrderItem(id, quantity int64, price float64) OrderItem {
	return OrderItem{
		Quantity: quantity,
		Item: Item{
			ItemID: id,
			Price:  price,
		},
	}
}

func (oi *OrderItem) GetTotal() float64 {
	return oi.Price * float64(oi.Quantity)
}
