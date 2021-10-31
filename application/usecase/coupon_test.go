package usecase

import (
	"testing"

	"github.com/diegoclair/branas-clean-code-arch/domain/repository"
	repositorymemory "github.com/diegoclair/branas-clean-code-arch/infra/data/repositoryMemory"
)

func TestCouponValidation(t *testing.T) {
	type fields struct {
		couponRepo repository.CouponRepository
	}
	type args struct {
		code string
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		wantIsValid bool
		wantErr     bool
	}{
		{
			name: "Should validate a valid coupon",
			fields: fields{
				couponRepo: repositorymemory.NewCouponRepositoryMemory(),
			},
			args: args{
				code: "VALE20",
			},
			wantIsValid: true,
			wantErr:     false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &newCoupon{
				couponRepo: tt.fields.couponRepo,
			}
			gotIsValid, err := u.CouponValidation(tt.args.code)
			if (err != nil) != tt.wantErr {
				t.Errorf("newCoupon.Execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotIsValid != tt.wantIsValid {
				t.Errorf("newCoupon.Execute() = %v, want %v", gotIsValid, tt.wantIsValid)
			}
		})
	}
}
