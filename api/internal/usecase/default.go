package usecase

import "api/internal/entity"

type UserRepository interface {
	Insert(user *entity.User) error
}
