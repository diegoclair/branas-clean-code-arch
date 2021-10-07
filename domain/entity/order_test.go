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
				t.Errorf("MakeOrder() -> got = %v, expect = %v", err, "invalid document")
			}
			if err != nil && !tt.wantErr {
				t.Errorf("MakeOrder() -> got = %v, expect = nil", err)
			}
			if err == nil && tt.wantErr {
				t.Errorf("MakeOrder() got = %v, expect some error", err)
			}
		})
	}
}
