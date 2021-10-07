package entity

type Item struct {
	ItemID      int64
	Category    string
	Description string
	Price       float64
	Quantity    int64
}

func NewItem(id, quantity int64, category, description string, price float64) (item Item) {
	return Item{
		ItemID:      id,
		Category:    category,
		Description: description,
		Price:       price,
		Quantity:    quantity,
	}
}
