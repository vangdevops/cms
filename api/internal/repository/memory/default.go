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

func (repo *UserRepository) Create(user *entity.User) error {
	repo.users = append(repo.users, user)
	return nil
}

func (repo *UserRepository) GetByID(id int64) *entity.User {
	for _, val := range repo.users {
		if val.ID == id {
			log.Println(id)
			return val
		}
	}
	return nil
}
