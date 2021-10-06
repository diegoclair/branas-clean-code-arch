package main

import "testing"

func TestMakeOrder(t *testing.T) {
	type args struct {
		document string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Should create an order with a valid document",
			args: args{
				document: "01234567890",
			},
			wantErr: false,
		},
		{
			name: "Should not create an order with an invalid document",
			args: args{
				document: "41234",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := MakeOrder(tt.args.document)
			if err != nil && !tt.wantErr || err == nil && tt.wantErr {
				t.Errorf("MakeOrder() with errors = %v, we expect error = %v", err, tt.wantErr)
			}
		})
	}
}
