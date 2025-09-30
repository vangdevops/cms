package main

import (
	"api/internal/entity"
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
	repo.Insert(&entity.User{Name: "hello"})
	http.HandleFunc("/", handler.Handler)
	log.Println(http.ListenAndServe(":8080", nil))
}
