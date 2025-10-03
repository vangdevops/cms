package main

import (
	"api/internal/handlers"
	"api/internal/repository/memory"
	"api/internal/routers/mux"
	"api/internal/services"
)

func main() {
	repo := memory.NewUserRepository()
	service := services.NewUserService(repo)
	handler := handlers.NewUserHandler(service)
	server := mux.NewUserMux(handler)
	server.Init()
}
