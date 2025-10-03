package entity

type Order struct {
	ProductName string
	ShopID      int64
	Quantity    uint8
	UserID      int64
	PayType     string
	Address     string
	Delivery    string
}
