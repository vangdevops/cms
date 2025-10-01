package services

import (
	"api/internal/entity"
	"api/internal/usecase"
)

type UserService struct {
	UserRepo usecase.UserRepository
}

func NewUserService(UserRepo usecase.UserRepository) *UserService {
	return &UserService{UserRepo: UserRepo}
}

func (Service *UserService) Create(user string, id int64, number string) error {
	Service.UserRepo.Create(&entity.User{Name: user, ID: id, Number: number})
	return nil
}

func (Service *UserService) GetByID(id int64) *entity.User {
	user := Service.UserRepo.GetByID(id)
	return user
}
