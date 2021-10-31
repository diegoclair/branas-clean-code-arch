package mysql

import (
	"github.com/diegoclair/branas-clean-code-arch/domain/entity"
	"github.com/diegoclair/branas-clean-code-arch/domain/repository"
)

type orderDatabase struct {
	conn connenction
}

func NewOrderDatabase(conn connenction) repository.OrderRepository {
	return &orderDatabase{
		conn: conn,
	}
}

func (r *orderDatabase) Save(order entity.Order) (err error) {
	query := `
		INSERT INTO tab_order 
			(code, cpf, coupon_id, issue_date, freight, sequence) VALUES
			(?, ?, ?, ?, ?, ?);
	`
	result, err := r.conn.Exec(query,
		order.Code,
		order.GetDocument(),
		order.Coupon.CouponID,
		order.IssueDate,
		order.GetFreight(),
		order.Sequence,
	)
	if err != nil {
		return err
	}

	orderID, err := result.LastInsertId()
	if err != nil {
		return err
	}

	for _, item := range order.OrderItems {
		query := `
			INSERT INTO tab_order_item
				(order_id, item_id, price, quantity) VALUES
				(?, ?, ?, ?);
		`
		_, err = r.conn.Exec(query,
			orderID,
			item.ItemID,
			item.Price,
			item.Quantity,
		)
		if err != nil {
			return err
		}
	}

	return nil
}
