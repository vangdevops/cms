package usecase

import "cms/internal/entity"

type OrderRepository interface {
	Add(ID int64, Order *entity.Order) error
	SetProduct(ID int64, Product *entity.Product) error
	SetPayType(ID int64, Paytype string) error
	SetAddress(ID int64, Address string) error
	SetDelivery(ID int64, Delivery string) error
	DeleteProduct(ID int64, Name string, Weight int) error
	Get(ID int64) (*entity.Order, error)
	List() (map[int64]*entity.Order, error)
}
