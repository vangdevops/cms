package main

import (
	"api/internal/adapter/web"
	"api/internal/entity"
	"api/internal/repository/memory"
	"log"
	"net/http"
)

func main() {
	repo := memory.NewUserRepository()
	adapter := web.NewUserAdapter(&repo)
	repo.Insert(&entity.User{Name: "hello"})
	http.HandleFunc("/", web.Handler)
	log.Println(http.ListenAndServe(":8080", nil))
}
