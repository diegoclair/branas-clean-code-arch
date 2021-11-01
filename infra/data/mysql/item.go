package mysql

import (
	"github.com/diegoclair/branas-clean-code-arch/domain/contract"
	"github.com/diegoclair/branas-clean-code-arch/domain/entity"
)

type itemDatabase struct {
	conn connenction
}

func NewItemDatabase(conn connenction) contract.ItemRepository {
	return &itemDatabase{
		conn: conn,
	}
}

func (r *itemDatabase) FindByID(id int64) (item entity.Item, err error) {
	query := `
		SELECT
			ti.id,
			ti.category,
			ti.description,
			ti.price,
			ti.height,
			ti.width,
			ti.length,
			ti.weight
		
		FROM tab_item 	ti
		WHERE ti.id 	= ?
	`
	row := r.conn.QueryRow(query, id)
	err = row.Scan(
		&item.ItemID,
		&item.Category,
		&item.Description,
		&item.Price,
		&item.Height,
		&item.Width,
		&item.Length,
		&item.Weight,
	)
	if err != nil {
		return item, err
	}

	return item, nil
}
