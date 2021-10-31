package mysql

import (
	"github.com/diegoclair/branas-clean-code-arch/domain/entity"
	"github.com/diegoclair/branas-clean-code-arch/domain/repository"
)

type couponDatabase struct {
	conn connenction
}

func NewCouponDatabase(conn connenction) repository.CouponRepository {
	return &couponDatabase{
		conn: conn,
	}
}

func (r *couponDatabase) FindByCode(code string) (coupon entity.Coupon, err error) {
	query := `
		SELECT
			tc.id,
			tc.code,
			tc.percentage,
			tc.expiration_date
		
		FROM tab_coupon 	tc
		WHERE tc.code	  	= ?
	`
	row := r.conn.QueryRow(query, code)
	err = row.Scan(
		&coupon.CouponID,
		&coupon.Code,
		&coupon.Percentage,
		&coupon.ExpirationDate,
	)
	if err != nil {
		return coupon, err
	}

	return coupon, nil
}
