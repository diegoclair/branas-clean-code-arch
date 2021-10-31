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
	document   string
	OrderItems []OrderItem
	Coupon     Coupon
	freight    float64
	IssueDate  time.Time
	Sequence   int64
}

func NewOrder(document string, sequence int64) (order Order, err error) {
	if !utils.IsValidDocumentNumber(document) {
		return order, errors.New("invalid document")
	}
	order.document = document
	order.IssueDate = time.Now()
	order.Sequence = sequence
	order.generateCode(order.IssueDate, order.Sequence)
	return order, nil
}
func (o *Order) generateCode(issueDate time.Time, sequence int64) {
	y, _, _ := issueDate.Date()
	year := strconv.Itoa(y)
	seq := strconv.Itoa(int(sequence))
	o.Code = year + utils.LeftPad(seq, 8, "0")
}

func (o *Order) AddItem(item Item, quantity int64) {
	o.freight += item.getFreight() * float64(quantity)
	o.OrderItems = append(o.OrderItems, OrderItem{ItemID: item.ItemID, Price: item.Price, Quantity: quantity})
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
	return utils.Round(total, 2)
}

func (o *Order) GetFreight() float64 {
	return o.freight
}

func (o *Order) GetDocument() string {
	return o.document
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

func (oi *OrderItem) GetTotal() float64 {
	return oi.Price * float64(oi.Quantity)
}
