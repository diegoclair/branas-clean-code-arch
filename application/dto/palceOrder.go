package dto

type OrderItemInput struct {
	Cpf        string
	OrderItems []OrderItems
}

type OrderItems struct {
	ItemID   int64
	Quantity int64
}

type OrderItemOutput struct {
	Total float64
}
