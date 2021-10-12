package entity

type Item struct {
	ItemID      int64
	Category    string
	Description string
	Price       float64
	Width       float64
	Height      float64
	Length      float64
}

func NewItem(id int64, category, description string, price, width, height, length float64) (item Item) {
	return Item{
		ItemID:      id,
		Category:    category,
		Description: description,
		Price:       price,
		Width:       width,
		Height:      height,
		Length:      length,
	}
}

func (i Item) getVolume() float64 {
	return i.Width / 100 * i.Height / 100 * i.Length / 100
}
