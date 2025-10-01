package main

import (
	"api/internal/handlers"
	"api/internal/repository/memory"
	"api/internal/services"
	"log"
	"net/http"
)

func main() {
	repo := memory.NewUserRepository()
	service := services.NewUserService(repo)
	handler := handlers.NewUserHandler(service)
	http.HandleFunc("/", handler.Handler)
	log.Println(http.ListenAndServe(":8080", nil))
}
