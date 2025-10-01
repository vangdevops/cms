package handlers

import (
	"api/internal/services"
	"encoding/json"
	"log"
	"net/http"
)

type UserHandler struct {
	UserService *services.UserService
}

func NewUserHandler(UserService *services.UserService) *UserHandler {
	return &UserHandler{UserService: UserService}
}

func (handler *UserHandler) Handler(w http.ResponseWriter, r *http.Request) {
	err := handler.UserService.Insert("hello")
	if err != nil {
		log.Println("test fail")
	}
	my := json.NewEncoder(w)
	my.Encode("")
}
