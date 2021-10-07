package entity

import "testing"

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
			order.addCoupon(NewCoupon("VALE20", 20))
			const totalShouldBe = 4375.12
			total := order.getTotal()
			if total != totalShouldBe {
				t.Errorf("getTotal() got %v, want %v", total, totalShouldBe)
			}
		})
	}
}
