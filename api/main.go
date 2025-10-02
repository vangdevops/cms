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
	mux := http.NewServeMux()
	mux.HandleFunc("/create", handler.Create)
	mux.HandleFunc("/get/{id}", handler.Get)
	log.Println(http.ListenAndServe(":8080", mux))
}
