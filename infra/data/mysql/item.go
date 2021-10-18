package mysql

import (
	"github.com/diegoclair/branas-clean-code-arch/domain/entity"
	"github.com/diegoclair/branas-clean-code-arch/domain/repository"
)

type itemDatabase struct {
	conn connenction
}

func NewItemDatabase(conn connenction) repository.ItemRepository {
	return &itemDatabase{
		conn: conn,
	}
}

func (r *itemDatabase) FindByID(id int64) (item entity.Item, err error) {
	query := `
		select * from tab_item where id = ?
	`
	r.conn.QueryRow(query, id)
	return item, nil
}
