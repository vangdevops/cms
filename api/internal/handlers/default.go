package handlers

import (
	"api/internal/services"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type UserHandler struct {
	UserService *services.UserService
}

func NewUserHandler(UserService *services.UserService) *UserHandler {
	return &UserHandler{UserService: UserService}
}

func (handler *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	handler.UserService.Create("hello", 1, "+11111111111")
	my := json.NewEncoder(w)
	my.Encode("Created")
}

func (handler *UserHandler) Get(w http.ResponseWriter, r *http.Request) {
	num, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		log.Println(err)
	} else {
		user := handler.UserService.GetByID(num)
		my := json.NewEncoder(w)
		my.Encode(user)
	}
}
