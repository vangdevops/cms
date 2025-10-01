package usecase

import "api/internal/entity"

type UserRepository interface {
	Create(user *entity.User) error
	GetByID(id int64) *entity.User
	//DeleteByID(id int64) error
}

type ShopRepository interface {
	Insert(user *entity.User) error
}

type ProductRepository interface {
	Insert(user *entity.User) error
}

type OrderRepository interface {
	Insert(user *entity.User) error
}
