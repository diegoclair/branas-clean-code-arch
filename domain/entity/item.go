package entity

type Item struct {
	ItemID      int64
	Category    string
	Description string
	Price       float64
}

func NewItem(id int64, category, description string, price float64) (item Item) {
	return Item{
		ItemID:      id,
		Category:    category,
		Description: description,
		Price:       price,
	}
}
