package entity

type User struct {
	ID     int64
	Name   string
	Number string
}

type Shop struct {
	Name        string
	Products    []Product
	Description string
	ID          int64
	OwnerID     int64
}

type Product struct {
	Shop        Shop
	Name        string
	Description string
	Price       float64
	Weight      float64
	Quantity    int64
}

type Order struct {
	User     User
	Products []Product
	PayType  string
	Address  string
	Delivery string
}
