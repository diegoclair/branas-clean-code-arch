package entity

import (
	"reflect"
	"testing"
	"time"
)

var (
	itemGuitarra    = NewItem(1, "Instrumentos Musicais", "Guitarra", 1119, 100, 30, 10, 3)
	itemAplificador = NewItem(1, "Instrumentos Musicais", "Amplificador", 4259.99, 100, 50, 50, 20)
	itemCabo        = NewItem(1, "Instrumentos Musicais", "Cabo", 30, 10, 10, 10, 0.9)
)

func TestNewOrderDocumentValidation(t *testing.T) {
	type args struct {
		document string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Should not create an order with an invalid document",
			args: args{
				document: "1111111111-11",
			},
			wantErr: true,
		},
		{
			name: "Should create an order with a valid document",
			args: args{
				document: "012.345.678-90",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewOrder(tt.args.document)
			if err != nil && tt.wantErr && err.Error() != "invalid document" {
				t.Errorf("MakeOrder() -> got = %v, want = %v", err, "invalid document")
			}
			if err != nil && !tt.wantErr {
				t.Errorf("MakeOrder() -> got = %v, want = nil", err)
			}
			if err == nil && tt.wantErr {
				t.Errorf("MakeOrder() got = %v, want some error", err)
			}
		})
	}
}

func TestNewOrderAddItems(t *testing.T) {
	type args struct {
		document string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Should create an order with 3 items",
			args: args{
				document: "012.345.678-90",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			order, _ := NewOrder(tt.args.document)
			order.AddItem(itemGuitarra, 1)
			order.AddItem(itemAplificador, 1)
			order.AddItem(itemCabo, 3)
			const totalShouldBe = 5468.99
			total := order.GetTotal()
			if total != totalShouldBe {
				t.Errorf("GetTotal() got %v, want %v", total, totalShouldBe)
			}
		})
	}
}

func TestNewOrderAddCoupon(t *testing.T) {
	type args struct {
		document string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Should create an order with 3 items with a coupon",
			args: args{
				document: "012.345.678-90",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			order, _ := NewOrder(tt.args.document)
			order.AddItem(itemGuitarra, 1)
			order.AddItem(itemAplificador, 1)
			order.AddItem(itemCabo, 3)
			order.AddCoupon(NewCoupon("VALE20", 20, time.Time{}))
			const totalShouldBe = 4375.19
			total := order.GetTotal()
			if total != totalShouldBe {
				t.Errorf("GetTotal() got %v, want %v", total, totalShouldBe)
			}
		})
	}
}

func TestOrderAddExpiredCoupon(t *testing.T) {
	type args struct {
		document string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Should get error when try to add an expired coupon",
			args: args{
				document: "012.345.678-90",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			order, _ := NewOrder(tt.args.document)
			order.AddItem(itemGuitarra, 1)
			err := order.AddCoupon(expiredCoupon)
			if err == nil {
				t.Error("Expected error with an expired coupon and get error = nil")
			}
			expiredErrorMessage := "coupon is expired"
			if err != nil && err.Error() != expiredErrorMessage {
				t.Errorf("got %v - want %v", err.Error(), expiredErrorMessage)
			}

		})
	}
}

func TestNewOrderItem(t *testing.T) {
	var (
		id       int64   = 1
		quantity int64   = 2
		price    float64 = 60
	)
	type args struct {
		id       int64
		quantity int64
		price    float64
	}
	tests := []struct {
		name string
		args args
		want OrderItem
	}{
		{
			name: "Should create an order item and validate total",
			args: args{
				id:       id,
				quantity: quantity,
				price:    price,
			},
			want: OrderItem{
				ItemID:   id,
				Quantity: quantity,
				Price:    price,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewOrderItem(tt.args.id, tt.args.quantity, tt.args.price)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewOrderItem() = %v, want %v", got, tt.want)
			}
			const totalShouldBe = 120
			total := got.GetTotal()
			if total != totalShouldBe {
				t.Errorf("NewOrderItem().GetTotal() = %v, want %v", total, totalShouldBe)
			}

		})
	}
}

func TestCalculateOrderFreight(t *testing.T) {
	type args struct {
		document string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Should create an order with 3 items and calculate it freight",
			args: args{
				document: "012.345.678-90",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			order, _ := NewOrder(tt.args.document)
			order.AddItem(itemGuitarra, 1)
			order.AddItem(itemAplificador, 1)
			order.AddItem(itemCabo, 3)
			const freightShouldBe = 260
			freight := order.GetFreight()
			if freight != freightShouldBe {
				t.Errorf("GetFreight() got %v, want %v", freight, freightShouldBe)
			}
		})
	}
}
