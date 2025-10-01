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

func (Service *UserService) Insert(user string) error {
	Service.UserRepo.Insert(&entity.User{Name: user})
	return nil
}
