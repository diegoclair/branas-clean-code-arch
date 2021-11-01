package mysql

import (
	"errors"

	"github.com/diegoclair/branas-clean-code-arch/domain/contract"
	"github.com/diegoclair/branas-clean-code-arch/domain/entity"
)

type orderDatabase struct {
	conn connenction
}

func NewOrderDatabase(conn connenction) contract.OrderRepository {
	return &orderDatabase{
		conn: conn,
	}
}

func (r *orderDatabase) Save(order *entity.Order) (err error) {
	query := `
		INSERT INTO tab_order 
			(code, cpf, coupon_id, issue_date, freight, total, sequence) VALUES
			(?, ?, ?, ?, ?, ?);
	`
	result, err := r.conn.Exec(query,
		order.Code,
		order.Document,
		order.Coupon.CouponID,
		order.IssueDate,
		order.Freight,
		order.GetTotal(),
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

func (r *orderDatabase) GetByCode(code string) (order entity.Order, err error) {
	query := `
		SELECT 
				to.id,
				to.code,
				to.cpf,
				to.coupon_id,
				to.issue_date,
				to.freight,
				to.total,
				to.sequence

		FROM 	tab_order  to
		WHERE 	to.code		= ?
	`
	row := r.conn.QueryRow(query, code)
	err = row.Scan(
		&order.OrderID,
		&order.Code,
		&order.Document,
		&order.IssueDate,
		&order.Freight,
		&order.Total,
		&order.Sequence,
	)
	if err != nil {
		return order, err
	}

	return order, errors.New("order not found")
}

func (r *orderDatabase) GetOrders() (orders []entity.Order, err error) {
	query := `
		SELECT 
			to.id,
			to.code,
			to.cpf,
			to.coupon_id,
			to.issue_date,
			to.freight,
			to.total,
			to.sequence

		FROM 	tab_order  to
	`
	rows, err := r.conn.Query(query)
	if err != nil {
		return orders, err
	}
	for rows.Next() {
		order := entity.Order{}
		err = rows.Scan(
			&order.OrderID,
			&order.Code,
			&order.Document,
			&order.Coupon.CouponID,
			&order.IssueDate,
			&order.Freight,
			&order.Total,
			&order.Sequence,
		)
		if err != nil {
			return orders, err
		}
		orders = append(orders, order)
	}

	return orders, errors.New("order not found")
}

func (r *orderDatabase) GetOrderItemsByOrderID(orderID int64) (orderItems []entity.OrderItem, err error) {

	queryItem := `
		SELECT 
				toi.item_id,
				toi.price,
				toi.quantity,
				ti.description

		FROM 	tab_order_item  toi
		
		INNER JOIN 	tab_item 	ti
			ON 	ti.id 	= toi.item_id
		
		WHERE 	toi.order_id	= ?
	`

	rows, err := r.conn.Query(queryItem, orderID)
	if err != nil {
		return orderItems, err
	}
	for rows.Next() {
		orderItem := entity.OrderItem{}
		err = rows.Scan(
			&orderItem.ItemID,
			&orderItem.Price,
			&orderItem.Quantity,
			&orderItem.Description,
		)
		if err != nil {
			return orderItems, err
		}
		orderItems = append(orderItems, orderItem)
	}
	return orderItems, nil
}

func (r *orderDatabase) Count() (count int64, err error) {

	query := `
		SELECT COUNT(*) FROM tab_order
	`
	row := r.conn.QueryRow(query)
	err = row.Scan(&count)
	if err != nil {
		return count, err
	}

	return count, nil
}
