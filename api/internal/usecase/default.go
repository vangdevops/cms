package usecase

import "api/internal/entity"

type UserRepository interface {
	Create(user *entity.User) error
	GetByID(id int64) (*entity.User, error)
	DeleteByID(id int64) error
}

type ShopRepository interface {
	Create()
	Delete()
	Change()
}

type ProductRepository interface {
	Create()
	Delete()
	Change()
}

type OrderRepository interface {
	Create()
	Delete()
	Change()
}
