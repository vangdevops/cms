package handlers

import (
	"api/internal/services"
	"encoding/json"
	"net/http"
)

type UserHandler struct {
	UserService *services.UserService
}

func NewUserHandler(UserService *services.UserService) *UserHandler {
	return &UserHandler{UserService: UserService}
}

func (handler *UserHandler) Handler(w http.ResponseWriter, r *http.Request) {
	my := json.NewEncoder(w)
	my.Encode("")
}
