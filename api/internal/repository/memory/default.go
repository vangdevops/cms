package memory

import (
	"api/internal/entity"
	"log"
)

type UserRepository struct {
	users []*entity.User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		users: make([]*entity.User, 0),
	}
}

func (repo *UserRepository) Insert(user *entity.User) error {
	repo.users = append(repo.users, user)
	log.Println(user)
	return nil
}
