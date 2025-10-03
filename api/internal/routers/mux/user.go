package mux

import (
	"api/internal/handlers"
	"log"
	"net/http"
)

type UserMux struct {
	Handler *handlers.UserHandler
}

func NewUserMux(Handler *handlers.UserHandler) *UserMux {
	return &UserMux{Handler: Handler}
}

func (mux *UserMux) Init() {
	server := http.NewServeMux()
	server.HandleFunc("/create", mux.Handler.Create)
	server.HandleFunc("/get/{id}", mux.Handler.Get)
	server.HandleFunc("/delete/{id}", mux.Handler.DeleteByID)
	log.Println(http.ListenAndServe(":8080", server))
}
