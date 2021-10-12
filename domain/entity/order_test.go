package entity

import (
	"reflect"
	"testing"
	"time"
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
			order.addItem(NewItem(1, "Instrumentos Musicais", "Guitarra", 1119), 1)
			order.addItem(NewItem(1, "Instrumentos Musicais", "Amplificador", 4259.99), 1)
			order.addItem(NewItem(1, "Instrumentos Musicais", "Cabo", 30), 3)
			const totalShouldBe = 5468.99
			total := order.getTotal()
			if total != totalShouldBe {
				t.Errorf("getTotal() got %v, want %v", total, totalShouldBe)
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
			name: "Should create an order with 3 items",
			args: args{
				document: "012.345.678-90",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			order, _ := NewOrder(tt.args.document)
			order.addItem(NewItem(1, "Instrumentos Musicais", "Guitarra", 1119), 1)
			order.addItem(NewItem(1, "Instrumentos Musicais", "Amplificador", 4259.90), 1)
			order.addItem(NewItem(1, "Instrumentos Musicais", "Cabo", 30), 3)
			order.addCoupon(NewCoupon("VALE20", 20, time.Time{}))
			const totalShouldBe = 4375.12
			total := order.getTotal()
			if total != totalShouldBe {
				t.Errorf("getTotal() got %v, want %v", total, totalShouldBe)
			}
		})
	}
}

func TestNewOrderAddExpiredCoupon(t *testing.T) {
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
			order.addItem(NewItem(1, "Instrumentos Musicais", "Guitarra", 1119), 1)
			err := order.addCoupon(NewCoupon("VALE20", 20, time.Date(2021, time.April, 2, 0, 0, 0, 0, time.Local)))
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
			total := got.getTotal()
			if total != totalShouldBe {
				t.Errorf("NewOrderItem().getTotal() = %v, want %v", total, totalShouldBe)
			}

		})
	}
}
