package entity

type Item struct {
	ItemID      int64
	Category    string
	Description string
	Price       float64
	Width       float64
	Height      float64
	Length      float64
	Weight      float64
}

func NewItem(id int64, category, description string, price, width, height, length, weight float64) (item Item) {
	return Item{
		ItemID:      id,
		Category:    category,
		Description: description,
		Price:       price,
		Width:       width,
		Height:      height,
		Length:      length,
		Weight:      weight,
	}
}

func (i Item) getVolume() float64 {
	return i.Width / 100 * i.Height / 100 * i.Length / 100
}

func (i Item) getDensity() float64 {
	return i.Weight / i.getVolume()
}
